
# NUMA Lock-Free Lab (Go)

steps:
- Step 1: Mutex queue — simple but poor p99 under contention.
- Step 2: Lock-free MSQueue (linked list) — fast, but lots of node churn (GC pressure).
- Step 3: EBR-inspired pooling — reuse nodes to reduce GC pauses (conceptually similar to SMR/EBR batching).
- Step 4: NUMA awareness — run with `numactl` to bind CPU/memory and compare tails.
- Step 5: Adaptive backoff — stabilize p99 under bursts by damping CAS storms.

## Quick start
```bash
cd numa_lockfree_lab
go run . -step=1 -rate=200000 -producers=4 -consumers=4 -duration=10s
go run . -step=2 -rate=200000 -producers=4 -consumers=4 -duration=10s
go run . -step=3 -rate=200000 -producers=4 -consumers=4 -duration=10s
go run . -step=5 -rate=200000 -producers=4 -consumers=4 -duration=10s -burst=6
```

### NUMA experiments (Linux)
> Requires a multi-socket machine.

**Good (local)**:
```bash
numactl --cpunodebind=0 --membind=0 go run . -step=3 -rate=200000 -producers=8 -consumers=8 -duration=15s
```

**Bad (cross-node)**:
```bash
numactl --cpunodebind=1 --membind=0 go run . -step=3 -rate=200000 -producers=8 -consumers=8 -duration=15s
```

Compare `p99/p99.9`. Expect fatter tails in the cross-node case.

### Flags
- `-step` = 1..5 (see above)
- `-rate` = open-loop constant arrival rate (ops/sec) — **CO-safe**
- `-duration` = test duration
- `-producers`, `-consumers` = goroutines
- `-burst` > 1.0 injects bursts (on/off pattern)
- `-nogc` disables Go GC during the run to make allocator effects visible

### Notes
- Go has GC, so I **simulate** reclamation effects with node pooling in Step 3 (reduces churn like EBR batching would in C/C++). 
- Adaptive backoff (Step 5) uses CAS-failure-driven spin/yield to tighten tails under bursts.
