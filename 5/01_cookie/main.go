package main

import (
	"fmt"
	"net/http"
	"time"
)

var loginFormTmpl = `
<html>
	<body>
	<form action="/get_cookie" method="post">
		Login: <input type="text" name="login">
		Password: <input type="password" name="password">
		<input type="submit" value="Login">
	</form>
	</body>
</html>
`

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // ставим обработчик на главную страницу "\" и пытаемся получить куку
		sessionID, err := r.Cookie("session_id") // кука это метод хттп реквеста
		if err == http.ErrNoCookie {
			w.Write([]byte(loginFormTmpl)) // если куи нету выводим
			return
		} else if err != nil {
			PanicOnErr(err)
		}
		fmt.Fprint(w, "Welcome, "+sessionID.Value) // если кука вернулась - просто вывожу на экран
	})

	http.HandleFunc("/get_cookie", func(w http.ResponseWriter, r *http.Request) { // ввожу логин и пароль
		r.ParseForm()
		inputLogin := r.Form["login"][0] //читаю логин
		expiration := time.Now().Add(365 * 24 * time.Hour) // даат протухания куки
		cookie := http.Cookie{ // создаю новую куку
			Name:    "session_id",
			Value:   inputLogin,
			Expires: expiration,
		}
		http.SetCookie(w, &cookie) // пишему ответ сюда
		http.Redirect(w, r, "/", http.StatusFound)
	})

	http.ListenAndServe(":8081", nil) // стартуем сервер
}

//PanicOnErr panics on error
func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
