package webserver

import (
	"fmt"
	"net/http"
)

func LauncWebServer() {
	http.HandleFunc("/", getHandle)
    fmt.Println("Сервер запущен на http://localhost:9089")
	http.ListenAndServe(":80", nil)
}

func getHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Привет Егор"))
	}
}