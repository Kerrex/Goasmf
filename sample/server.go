package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type MyHandler struct {

}

func (m MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if strings.HasPrefix(request.URL.Path, "/static") || strings.HasPrefix(request.URL.Path, "/templates") {
		http.FileServer(http.Dir("sample/")).ServeHTTP(writer, request)
		return
	}
	file, err := ioutil.ReadFile("sample/index.html")
	if err != nil {
		panic(err)
	}

	writer.WriteHeader(200)
	writer.Write(file)
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", MyHandler{}))
}
