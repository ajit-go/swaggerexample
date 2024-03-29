package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/ajit-go/swaggerexample/docs"
	_ "github.com/ajit-go/swaggerexample/docs"
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

	//service swagger
	httpSwagger.URL("swagger.json")
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}

// @title Swagger Example API
// @version 1.0
// @description This is a example of swagger with mux.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.myproject.io/support
// @contact.email support@myproject.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8082
// @BasePath /v2

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://api.stage.context.cloud.sap/uaa/oauth/token
// @scope.view  Grants read access
// @scope.manage Grants read and write access 
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
		if !strings.Contains(r.URL.Path, "swagger") {
			w.Header().Add("Content-Type", "application/json")
		}
		next.ServeHTTP(w, r)
	})
}
