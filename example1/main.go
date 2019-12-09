package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("--------------in example 1-----------")
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:8082", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page - example 1")
}
