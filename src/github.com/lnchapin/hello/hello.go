package main

import (
	"io"
	"net/http"

	"github.com/lnchapin/string"
)

type MyHandler int

func (h MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/reverse":
		io.WriteString(res, string.Reverse("Hello, new gopher!"))
	case "/":
		io.WriteString(res, "Hello, new gopher!")
	}
}

func main() {
	var h MyHandler
	http.ListenAndServe(":9000", h)
}
