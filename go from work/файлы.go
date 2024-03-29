package main

import (
        "bytes"
        "fmt"
        "log"
        "os"
)
func main() {
	
        dataForFile := []byte("Тестовая строка, предназначенная для записи в файл")
        file_name := "test_file.txt"

        // Создаем новый файл и записываем в него данные dataForFile
        if err := os.WriteFile(file_name, dataForFile, 0600); err != nil {
                log.Fatal(err)
        }

        // Читаем данные из того же файла
        dataFromFile, err := os.ReadFile(file_name)
        if err != nil {
                log.Fatal(err)
        }

        // Сравниваем исходные данные с записанными в файл и прочитанными из него
        fmt.Printf("dataForFile == dataFromFile: %v\n", bytes.Equal(dataFromFile, dataForFile))

        // Получаем текущую директорию
        currentDir, err := os.Getwd()
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println((currentDir))

        // Изучаем содержимое директории
        filesFromDir, err := os.ReadDir(currentDir)
        if err != nil {
                if err != nil {
                        log.Fatal(err)
                }
        }
        for _, file := range filesFromDir {
                // Проходим по всем найденным файлам и печатаем их имя и размер
                info, _ := file.Info()
                fmt.Printf("|_name: %s, size: %d\n", file.Name(), info.Size())
        }

        // Output:
        // dataForFile == dataFromFile: true
        // /home/<user>/<pwd>
        // |_name: main.go, size: 1491
        // |_name: test.txt, size: 93
        // |_...
}

----------------------------------------------------------------------------
	
	
// создаем файл
os.Create("text.txt")
// переименовываем файл
os.Rename("text.txt", "new_text.txt")
// удаляем файл
os.Remove("new_text.txt")
// кстати, os позволяет работать не только с файлами
os.Open(path) // открыть файл
// выходим из программы:
os.Exit(0)
		
file1, _ := os.Create("text.txt")
file2, _ := os.Create("text.txt")
info1, _ := file1.Stat() // функция Stat возвращает информацию о файле и ошибку
info2, _ := file2.Stat()
fmt.Println(os.SameFile(info1, info2)) // true

// вот что мы можем получить из FileInfo:
// A FileInfo describes a file and is returned by Stat and Lstat.
type FileInfo interface {
	Name() string       // base name of the file
	Size() int64        // length in bytes for regular files; system-dependent for others
	Mode() FileMode     // file mode bits
	ModTime() time.Time // modification time
	IsDir() bool        // abbreviation for Mode().IsDir()
	Sys() interface{}   // underlying data source (can return nil)
}

// записываем в файл

file1, _ := os.Create("text.txt")
file1.WriteString("1 строка \n")
file1.WriteString("2 строка \n")
file1.Close()

// внутри файла будет:
// 1 строка 
// 2 строка 
	
----------------------------------------------------------------------------
	
	
Сам пробовал создать файл и записать туда
		dataInFile := []byte("запись в файл")

	filetest, err := os.Create("text.txt")
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(filetest.Name(), dataInFile, 0600); err != nil {
		log.Fatal(err)
	}

	datafromfile, err := os.ReadFile(filetest.Name())
	if err != nil {
		log.Fatal(err)
	}

	filetest.Close()
	fmt.Print(string(datafromfile))
		
	
	
----------------------------------------------------------------------------
		
	bufio
		
	
	bufio.Reader
		
Данный тип создается с помощью функций:

func NewReader(rd io.Reader) *Reader // создает Reader со стандартным буфером 4096 байт

func NewReaderSize(rd io.Reader, size int) *Reader // создает Reader с произвольным буфером
//Рассмотрим некоторые из методов bufio.Reader и примеры работы:

file, err := os.Open("test.txt")
if err != nil {
	...
}
defer file.Close()

rd := bufio.NewReader(file)

buf := make([]byte, 10)
n, err := rd.Read(buf) // читаем в buf 10 байт из ранее открытого файла
if err != nil && err != io.EOF {
	// io.EOF не совсем ошибка - это состояние, указывающее, что файл прочитан до конца
	...
}
fmt.Printf("прочитано %d байт: %s\n", n, buf) // прочитано 10 байт: bufio ...

s, err := rd.ReadString('\n') // читаем данные до разрыва абзаца ('\n')
fmt.Printf("%s\n", s)         // ... здесь будет строка
//bufio.Reader позволяет читать данные по байтам, рунам, строкам и пр., указывать символ, на котором необходимо прекратить чтение. Когда данные будут прочитаны до конца, метод вернет ошибку io.EOF.





	bufio.Writer
	
//bufio.Writer создан для записи в объекты, удовлетворяющие интерфейсу io.Writer, но предоставляет ряд более высокоуровневых методов, в частности метод WriteString(s string):

file, err := os.Create("test.txt")
if err != nil {
	...
}
defer file.Close()

w := bufio.NewWriter(file)
n, err := w.WriteString("Запишем строку")
if err != nil {
	...
}
fmt.Printf("Записано %d байт\n", n) // Записано 27 байт

// bufio.Writer имеет собственный буфер, чтобы быть уверенным, что данные точно записаны,
// вызываем метод Flush()
w.Flush()
//Как вы уже поняли, создается объект функцией NewWriter(w io.Writer).

	bufio.Scanner

//bufio.Scanner создан для построчного чтения данных. Создается он функцией NewScanner(r io.Reader), посмотрим, как работает этот тип:

file, err := os.Open("test.txt")
if err != nil {
	panic(err)
}
defer file.Close()

s := bufio.NewScanner(file)

// Я заранее записал в файл 5 цифр, каждую на новой строке
for s.Scan() { // возвращает true, пока файл не будет прочитан до конца
	fmt.Printf("%s\n", s.Text()) // s.Text() содержит данные, считанные на данной итерации
}

// 1
// 2
// 3
// 4
// 5
		
----------------------------------------------------------
package main

import (
	//"bytes"
	"bufio"
	//"fmt"
	"io"
	"os"
	"strconv"

)

func main() {

	var sum int
// Мы создаем новый объект сканнер, аргументом которому передаем потток ввода.
	scanner := bufio.NewScanner(os.Stdin)
// Теперь у нас появляется возможность "слушать" ввод:scanner.Scan()
// возвращает true, пока файл не будет прочитан до конца

	for scanner.Scan() {
		s, _ := strconv.Atoi(scanner.Text())
		sum = sum + s
	}
	
	
	io.WriteString(os.Stdout, strconv.Itoa(sum)) // вывод

}

---------------------------------------------------------------------

	buffio Reader


reader := bufio.NewReader(os.Stdin)

for {
    line, -, err := reader.ReadLine() // ReadLine возвращает line []byte, isPrefix bool, err error
    if err == io.EOF {
	    break
	}
    
    txt, _ := string(line) // Конвертируем байты в строку. Эскейпим ошибку, лень обрабатывать.
    // Для числа, мы бы делали так:
    // num, _ := strconv.Atoi(string(line)) 
}

// Или

for {
    s, err := rd.ReadString('\n')
    if err == io.EOF {
        break
    }
    ...
}

// Вывод

writer := bufio.NewWriter(os.Stdout)
w.WriteString(txt) // Записываем строку
w.Flush() // При выполнении методов WriteString(), WriteRune(), WriteByte() и вы не поверите Write();
          // данные вначале накапливаются в буфере, а чтобы сбросить их в источник данных,
          // необходимо вызвать метод Flush().
          

---------------------------------------------
	
	
	encoding/csv
		
		

func main() {

	buf := bytes.NewBuffer(nil)

	w := csv.NewWriter(buf)

	for i := 1; i <= 3; i++ {
		// Запись данных может производится поэтапно, например в цикле
		val1 := fmt.Sprintf("row %d col 1", i)
		val2 := fmt.Sprintf("row %d col 2", i)
		val3 := fmt.Sprintf("row %d col 3", i)

		if err := w.Write([]string{val1, val2, val3}); err != nil {
			// Аргументом Write является срез строк
			// ...
		}
	}
	w.Flush() // Этот метод приведет к фактической записи данных из буфера csv.Writer в buf

	// Либо данные можно записать за один раз
	// w.WriteAll([][]string{ // Аргументом WriteAll является срез срезов строк
	// 	{"row 4 col 1", "row 4 col 2", "row 4 col 3"},
	// 	{"row 5 col 1", "row 5 col 2", "row 5 col 3"},
	// })

	r := csv.NewReader(buf) // читаем

	for i := 1; i <= 2; i++ {
		// Читать данные мы тоже можем построчно, получая срез строк за каждую итерацию
		row, err := r.Read()
		if err != nil && err != io.EOF {
			// Здесь тоже нужно учитывать конец файла
			// ...
		}
		fmt.Println(row)
	}

	// Либо прочитать данные за один раз
	data, err := r.ReadAll()
	if err != nil {
		// Когда мы читаем данные до конца файла io.EOF не возвращается, 
		//а служит сигналом к завершению чтения
		// ...
	}

	for _, row := range data {
		fmt.Println(row)
	}

}

---------------------------------------------
	
	path и path/filepath
		
		
	func Walk(root string, walkFn WalkFunc) error
//root - директория, с которой начинается обход
//walkFn - функция вида 
	func(path string, info os.FileInfo, err error) error
	
	for example:
		
	
	
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err // Если по какой-то причине мы получили ошибку, проигнорируем эту итерацию
	}

	if info.IsDir() { 
		return nil // Проигнорируем директории
	}

	fmt.Printf("Name: %s\tSize: %d byte\tPath: %s\n", info.Name(), info.Size(), path)
	return nil
}

func main() {
	const root = "./test" // Файлы моей программы находятся в другой директории

	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}

	// Name: file1     Size: 6 byte    Path: test/dir1/file1
	// Name: file2     Size: 6 byte    Path: test/dir1/file2
	// Name: file3     Size: 6 byte    Path: test/dir2/file3
	// Name: file4     Size: 6 byte    Path: test/dir3/file4
	// Name: file5     Size: 6 byte    Path: test/dir3/file5
	// Name: file6     Size: 6 byte    Path: test/dir3/file6
}



дополнительно, если хотим проигнорировать дирректории некторые
	 // Проигнорируем директории
   if info.IsDir() {  // в этой строке проверяется дирректория ли это
      // Не спускаемся в директории .git и .idea
      if info.Name() == ".git" || info.Name() == ".idea" {
         return filepath.SkipDir
      }

---------------------------------------------

	задачка найти в куче файлов один с информацией
	
package main

import (
	"encoding/csv"
	"fmt"
	//"io"
	"log"
	"os"
	"path/filepath"
	//"strings"

)

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	myfile, err := os.Open(path) // открываем каждый файл
	if err != nil {
		log.Fatal(err)
	}

	defer myfile.Close() // закрываем вконце
	r := csv.NewReader(myfile) // читаем файл

	record, _ := r.ReadAll() // пробегаемся по файлу
		if len(record) == 10 { // в файле длина записи равно 10ти
   		fmt.Println(record[4][2]) // record - данные в файле
	}

	return nil

}

func main() {

	const root = "./task"

	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}

}



-----------------------------------------------------------

задача найти в огромном файле типа 
1843245343472903481;114969837749034510;3591136229644178518;5518506782192540109;9169919572500480803;2961265865280294526;5117390539497588519;8500878058245396291;9172550547724248583;2264506197246516816;1605981750134825800;7763391713410542823;

число ноль и вернуть его позицию в файле.
	
	
package main

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	//"io"
	"log"
	"os"
	"path/filepath"
	"strings"

)

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	myfile, err := os.Open(path) // открываем каждый файл
	if err != nil {
		log.Fatal(err)
	}
	defer myfile.Close() // закрываем вконце
	
	r := csv.NewReader(myfile) // читаем файл
	
	record, _ := r.Read() // пробегаемся по файлу, именно Реад, не реадалл
		
	splitedData := strings.Split(record[0], ";")


	for i, value := range splitedData {

		if value == "0"{
		fmt.Print(i)
		}
	}

	return nil

}

func main() {

	const root = "./dich"

	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}
}


-------------------------------------









