package main

import (
	"log"
	"strings"
	"net/http"

	_ "github.com/ajit-go/swaggerexample/docs"
	"github.com/ajit-go/swaggerexample/docs"
	m "github.com/ajit-go/swaggerexample/mvc"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

var c = m.Controller{}

func handleRoutes(r *mux.Router) {
	r.HandleFunc("/", c.Home)
	r.HandleFunc("/articles", c.GetArticles)
	r.HandleFunc("/article/{aid}", c.GetArticle).Methods("GET")
	r.HandleFunc("/article", c.CreateArticle).Methods("POST")
	r.HandleFunc("/article/{aid}", c.UpdateArticle).Methods("PUT")
	httpSwagger.URL("swagger.json")

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

}



func main() {
	router := mux.NewRouter()
	router.Use(commonMiddleware)
	docs.SwaggerInfo.Host = "localhost:8082"
	docs.SwaggerInfo.BasePath = ""
	handleRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:8082", nil))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, "swagger"){
			w.Header().Add("Content-Type", "application/json")
		}
		next.ServeHTTP(w, r)
	})
}
