package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"io/ioutil"
)

type Todo struct {
	Name      string `json:"name"` // если не укажем модификаторы json, апи будет возвращат так как указано в этой структуре
	Done      bool   `json:"done"` // это специальное переформирование для конвертации в json, с помощью пакета на строке 4
}

func main() {
	todos := []Todo{
		{"Выучить Go", false},
	}
// рутпасс который отдает обычную html, не шаблон а обычную
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// здесь надо отдать статический файл, который будет общаться с API из браузера
		// открываем файл. взять прочитать файл и вывести его
		fileContents, err := ioutil.ReadFile("index.html")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// и выводим содержимое файла
		w.Write(fileContents)
	})
//специальный пасс тудус, который в зависимости от типа запроса делает разное
	http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request ", r.URL.Path)
		defer r.Body.Close()

		// разные методы обрабатываются по-разному
		switch r.Method {
		// GET для получения данныхт
		case http.MethodGet:
			// преобразуем структуру в json
			productsJson, _ := json.Marshal(todos) //берет данные из структуры и преабразует в json
			w.Header().Set("Content-Type", "application/json") //тут браузер понял что мы ему возвращаем json
			w.WriteHeader(http.StatusOK)
			w.Write(productsJson)
		// POST для добавления чего-то нового туду
		case http.MethodPost:
			decoder := json.NewDecoder(r.Body) // преабразуем новое в структуру
			todo := Todo{}
			// преобразуем json запрос в структуру
			err := decoder.Decode(&todo)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			todos = append(todos, todo)
		// PUT для обновления существующей информации
		case http.MethodPut:
			id := r.URL.Path[len("/todos/"):]
			index, _ := strconv.ParseInt(id, 10, 0)
			todos[index].Done = true
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}

	})

	http.ListenAndServe(":8080", nil)
}