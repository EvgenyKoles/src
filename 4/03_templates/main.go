package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Todo struct {
	Name string
	Done bool
}

func IsNotDone(todo Todo) bool {
	return !todo.Done
}

type Handler struct{
}


func(h Handler) ServeHttp()


// создаем шаблон, передаем туда функции. далее пасрсим файлы в этой дерриктории
func main() {


	tmpl, err := template.New("template.html").Funcs(template.FuncMap{"IsNotDone": IsNotDone}).ParseFiles("template.html")
	if err != nil { // если не смог заэкспандить этот шаблон, то пишем ошибку
		log.Fatal("Can not expand template", err)
		return
	}

	todos := []Todo{ // если все ок, набрасаем список туду
		{"Выучить Go", false},
		{"Посетить лекцию по вебу", false},
		{"...", false},
		{"Profit", false},
	}

	// и на корневой пасс нашего приложения повесим экзекют этого шаблона. обработка этого шаблона
	// проверяет что туда не дан

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// читаем из urlencoded запроса
			param := r.FormValue("id")
			// преобразуем строку в int
			index, _ := strconv.ParseInt(param, 10, 0)
			todos[index].Done = true
		}

		// исполняем шаблон. если не смогли вернем 500ю
		err := tmpl.Execute(w, todos)
		if err != nil {
			// вернем 500 и напишем ошибку
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8081", nil)
}
