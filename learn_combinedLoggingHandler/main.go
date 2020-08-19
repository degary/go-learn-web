package main

import (
	"github.com/gorilla/handlers"
	"io"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", useCombinedLogginghandler(handler()))
	http.ListenAndServe(":1234", nil)
}

func handler() http.Handler {
	return http.HandlerFunc(myHandler)
}

func myHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	io.WriteString(rw, "Hello World")
}

func useCombinedLogginghandler(next http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, next)
}
