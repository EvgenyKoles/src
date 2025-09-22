package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"runtime/debug"
	"sort"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// -------------------- Payload & metrics --------------------

type payload struct{ t time.Time }

type quantiles struct{ P50, P95, P99, P999 time.Duration }

func computeQs(ds []time.Duration) quantiles {
	if len(ds) == 0 {
		return quantiles{}
	}
	cp := make([]int64, len(ds))
	for i, d := range ds {
		cp[i] = d.Nanoseconds()
	}
	sort.Slice(cp, func(i, j int) bool { return cp[i] < cp[j] })
	at := func(p float64) time.Duration {
		if len(cp) == 1 {
			return time.Duration(cp[0])
		}
		idx := int(math.Ceil(p*float64(len(cp))) - 1)
		if idx < 0 {
			idx = 0
		}
		if idx >= len(cp) {
			idx = len(cp) - 1
		}
		return time.Duration(cp[idx])
	}
	return quantiles{at(0.50), at(0.95), at(0.99), at(0.999)}
}

// -------------------- Queue interface --------------------

type Queue interface {
	Enq(v any) bool
	Deq() (any, bool)
}

// -------------------- Step 1: Mutex queue --------------------

type MutexQ struct {
	mu sync.Mutex
	L  []any
}

func NewMutexQ() *MutexQ { return &MutexQ{L: make([]any, 0, 1024)} }

func (q *MutexQ) Enq(v any) bool {
	q.mu.Lock()
	q.L = append(q.L, v)
	q.mu.Unlock()
	return true
}

func (q *MutexQ) Deq() (any, bool) {
	q.mu.Lock()
	if len(q.L) == 0 {
		q.mu.Unlock()
		return nil, false
	}
	v := q.L[0]
	q.L = q.L[1:]
	q.mu.Unlock()
	return v, true
}

// -------------------- Step 2+: Lock-free MSQueue --------------------

type node struct {
	val  any
	next atomic.Pointer[node]
	_    [56]byte // padding ~1 cache line
}

type MSQueue struct {
	head            atomic.Pointer[node]
	tail            atomic.Pointer[node]
	enableAdaptive  bool // step 5
	enablePooling   bool // step 3–5
}

func NewMSQueue(adaptive, pooling bool) *MSQueue {
	q := &MSQueue{enableAdaptive: adaptive, enablePooling: pooling}
	dummy := &node{}
	q.head.Store(dummy)
	q.tail.Store(dummy)
	return q
}

func (q *MSQueue) Enq(v any) bool {
	n := getNode(v, q.enablePooling)
	var casFails int
	for {
		t := q.tail.Load()
		next := t.next.Load()
		if t == q.tail.Load() {
			if next == nil {
				if t.next.CompareAndSwap(nil, n) {
					q.tail.CompareAndSwap(t, n) // help advance
					return true
				}
			} else {
				q.tail.CompareAndSwap(t, next)
				casFails++
				if q.enableAdaptive {
					adaptiveBackoff(casFails)
				}
			}
		}
	}
}

func (q *MSQueue) Deq() (any, bool) {
	var casFails int
	for {
		h := q.head.Load()
		t := q.tail.Load()
		next := h.next.Load()
		if h == q.head.Load() {
			if h == t {
				if next == nil {
					return nil, false
				}
				q.tail.CompareAndSwap(t, next)
				casFails++
				if q.enableAdaptive {
					adaptiveBackoff(casFails)
				}
			} else {
				v := next.val
				if q.head.CompareAndSwap(h, next) {
					// retire old head
					retire(h, q.enablePooling)
					return v, true
				}
				casFails++
				if q.enableAdaptive {
					adaptiveBackoff(casFails)
				}
			}
		}
	}
}

// -------------------- Step 3: EBR-inspired pooling --------------------
// Пул узлов уменьшает аллокации/GC-паузы: концептуально похоже на batch free в EBR.

const shards = 64
var pools [shards]chan *node

func initPools(capacity int) {
	for i := 0; i < shards; i++ {
		pools[i] = make(chan *node, capacity)
	}
}

func shard() int {
	// используем адрес локальной переменной как источник энтропии (дешево и без блокировок)
	var dummy int
	x := uintptr(unsafe.Pointer(&dummy))
	return int((x >> 8) & (shards - 1))
}

func getNode(v any, pooling bool) *node {
	if !pooling {
		return &node{val: v}
	}
	s := shard()
	select {
	case n := <-pools[s]:
		n.val = v
		n.next.Store(nil)
		return n
	default:
		return &node{val: v}
	}
}

func putNode(n *node, pooling bool) {
	if !pooling || n == nil {
		return
	}
	s := shard()
	select {
	case pools[s] <- n:
	default:
	}
}

func retire(n *node, pooling bool) { putNode(n, pooling) }

// -------------------- Step 5: Adaptive backoff --------------------

func adaptiveBackoff(fails int) {
	if fails < 4 {
		spin(64 << fails)
		return
	}
	if fails < 8 {
		runtime.Gosched()
		return
	}
	time.Sleep(time.Microsecond)
}

func spin(iters int) {
	for i := 0; i < iters; i++ {
		// busy-wait; в Go нет явного PAUSE
	}
}

// -------------------- Open-loop (CO-safe) driver --------------------

func runBenchmark(step int, rate int, duration time.Duration, producers, consumers int, burst float64, nogc bool) {
	if nogc {
		debug.SetGCPercent(-1)
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Широковещательный done
	done := make(chan struct{})
	timer := time.NewTimer(duration)
	go func() {
		<-timer.C
		close(done)
	}()

	// queue on step
	var q Queue
	switch step {
	case 1:
		q = NewMutexQ()
	case 2:
		q = NewMSQueue(false, false)
	case 3, 4:
		initPools(1 << 14)
		q = NewMSQueue(false, true)
	case 5:
		//initPools(1 << 14)
		q = NewMSQueue(true, true)
	default:
		panic("unknown step")
	}

	latCh := make(chan time.Duration, 1<<20)

	// Consumers
	var wg sync.WaitGroup
	for i := 0; i < consumers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				// быстрый выход
				select {
				case <-done:
					return
				default:
				}

				v, ok := q.Deq()
				if !ok {
					// don't burn cpu on empty queue
					time.Sleep(50 * time.Microsecond)
					continue
				}

				lat := time.Since(v.(payload).t)

				// send metrics, wait done and don't block
				select {
				case <-done:
					return
				case latCh <- lat:
				default:
					// drop, if buffer is full
				}
			}
		}()
	}

	// Producers (open-loop, CO-safe)
	perSec := rate
	perProd := perSec / max(1, producers)
	interval := time.Second / time.Duration(max(1, perProd))
	start := time.Now()

	for p := 0; p < producers; p++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			next := start
			for {
				// выход по done
				select {
				case <-done:
					return
				default:
				}

				// bursts: on 10% of time windows
				mult := 1.0
				if burst > 1.0 && (time.Now().UnixNano()/int64(time.Millisecond))%2000 < 200 {
					mult = burst
				}
				next = next.Add(time.Duration(float64(interval) / mult))

				// уважаем done во время ожидания
				if sleep := time.Until(next); sleep > 0 {
					t := time.NewTimer(sleep)
					select {
					case <-done:
						t.Stop()
						return
					case <-t.C:
					}
				}

				if ok := q.Enq(payload{t: time.Now()}); !ok {
					// в bounded структурах можно подождать
					time.Sleep(50 * time.Microsecond)
				}
			}
		}()
	}

	// Ждём завершения по done
	<-done
	wg.Wait()       // все горутины завершились
	close(latCh)    // теперь канал можно закрыть

	var lats []time.Duration
	for x := range latCh {
		lats = append(lats, x)
	}
	qs := computeQs(lats)

	fmt.Printf("Step=%d  Producers=%d  Consumers=%d  Rate=%d ops/s  Dur=%s  Samples=%d\n",
		step, producers, consumers, rate, duration.String(), len(lats))
	fmt.Printf("p50=%s  p95=%s  p99=%s  p99.9=%s\n", qs.P50, qs.P95, qs.P99, qs.P999)

	if step == 4 {
		fmt.Println("\n[NUMA Hint] Run with Linux numactl to see cross-node penalties:")
		fmt.Println("  Good (local): numactl --cpunodebind=0 --membind=0 go run . -step=4 -rate=200000 -producers=8 -consumers=8 -duration=15s")
		fmt.Println("  Bad  (cross): numactl --cpunodebind=1 --membind=0 go run . -step=4 -rate=200000 -producers=8 -consumers=8 -duration=15s")
	}
}

// -------------------- main --------------------

func main() {
	var (
		step      = flag.Int("step", 1, "1=mutex, 2=lock-free MSQ, 3=MSQ+pool (EBR-like), 4=NUMA demo, 5=adaptive backoff")
		rate      = flag.Int("rate", 100000, "open-loop rate (ops/sec)")
		duration  = flag.Duration("duration", 5*time.Second, "duration")
		producers = flag.Int("producers", 1, "producers")
		consumers = flag.Int("consumers", 1, "consumers")
		burst     = flag.Float64("burst", 1.0, "burst factor (>1.0 => on/off spikes)")
		nogc      = flag.Bool("nogc", false, "disable GC during run")
	)
	flag.Parse()


	runBenchmark(*step, *rate, *duration, *producers, *consumers, *burst, *nogc)
}

// helpers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}