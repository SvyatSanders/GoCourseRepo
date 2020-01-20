// Дополните функцию hello() http-сервера так, чтобы принять и отобразить на странице один GET-параметр, например name.
// При этом в браузере запрос будет выглядеть так: http://localhost/hello&name=World.
// Значение этого параметра надо получить в функции и отобразить при выводе html-кода

// запрос: http://localhost/hello?name=fuckers

package main

import (
	"fmt"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	queryValues := req.URL.Query()
	fmt.Fprintf(res, "hello, %s!\n", queryValues.Get("name"))
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":80", nil)
}
