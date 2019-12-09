package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("--------------in example 2-----------")
	router := mux.NewRouter()
	router.Use(commonMiddleware)
	router.HandleFunc("/", home)

	router.HandleFunc("/article/{aid}", getArticle).Methods("GET")
	router.HandleFunc("/article", createArticle).Methods("POST")
	router.HandleFunc("/article/{aid}", updateArticle).Methods("PUT")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:8082", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page - example 2")
}














func getArticle(w http.ResponseWriter, r *http.Request) {
	aid := mux.Vars(r)["aid"]
	a := article{AId: aid, Title: "sonos one", Desc: "sonos one smart speaker", Content: "Speakers with alexa and google home"}
	json.NewEncoder(w).Encode(a)
}
func createArticle(w http.ResponseWriter, r *http.Request) {
	a := article{AId: "a2", Title: "sonos two", Desc: "sonos one smart speaker", Content: "Speakers with alexa and google home"}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(a)
}
func updateArticle(w http.ResponseWriter, r *http.Request) {
	a := article{AId: "a1", Title: "sonos one SL", Desc: "sonos one smart speaker", Content: "Speakers with alexa and google home"}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(a)
}

type article struct {
	AId     string `json:"aid"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
