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
	httpSwagger.URL("swagger.json")

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

}

// @title Swagger Example API
// @version 1.0
// @description This is a example of swagger with mux.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host github.com/ajit-go/swaggerexample
// @BasePath /v2
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
