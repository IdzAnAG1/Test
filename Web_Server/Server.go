package webserver

import (
	"fmt"
	"net/http"
)

func LauncWebServer() {
	port:=":80"
	http.HandleFunc("/", getHandle)
    fmt.Println("Сервер запущен на http://localhost",port)
	http.ListenAndServe(port, nil)
}

func getHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Привет Егор"))
	}
}