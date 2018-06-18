package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", write)
	http.HandleFunc("/newID", newID)
	fmt.Println("listening on: 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func write(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	io.WriteString(w, r.URL.Path+"\n")
}

func newID(w http.ResponseWriter, r *http.Request) {
	newID := uuid.Must(uuid.NewV4())

	fmt.Println(newID)
	io.WriteString(w, newID.String()+"\n")
}
