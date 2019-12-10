package mvc

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/ajit-go/swaggerexample/model"
	"github.com/gorilla/mux"
)

//Controller struct to hold all rest handler fucntions
type Controller struct{}

var data = Data{}

// Home page
// @Summary Home page
// @Description Home page
// @Accept  json
// @Produce  json
// @Param name query string false "search by name"
// @Success 200 {object} model.Hello
// @Header 200 {string} Token "qwerty"
// Failure 400 {object} model.Error
// @Router / [get]
func (c *Controller) Home(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Hello{"home of api"})
}

//GetArticles page
// @Summary list Articles
// @Description get Articles
// @Accept  json
// @Produce  json
// @Param tenant header string true "tenant name"
// @Success 200 {array} model.Article
// @Header 200 {string} Token "qwerty"
// Failure 400 {object} model.Error
// @Failure 404 {string} model.Error
// @Failure 500 {object} model.Error
// @Router /articles [get]
// @Security BasicAuth
func (c *Controller) GetArticles(w http.ResponseWriter, r *http.Request) {
	initialize()
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(data)
}

func (c *Controller) GetArticle(w http.ResponseWriter, r *http.Request) {
	aid := mux.Vars(r)["aid"]
	a := Article{AId: aid, Title: "sonos one", Desc: "sonos one smart speaker", Content: "Speakers with alexa and google home"}
	json.NewEncoder(w).Encode(a)
}
func (c *Controller) CreateArticle(w http.ResponseWriter, r *http.Request) {
	a := Article{AId: "a2", Title: "sonos two", Desc: "sonos one smart speaker", Content: "Speakers with alexa and google home"}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(a)
}
func (c *Controller) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	a := Article{AId: "a1", Title: "sonos one SL", Desc: "sonos one smart speaker", Content: "Speakers with alexa and google home"}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(a)
}

func initialize() {
	data.Articles = []Article{
		Article{AId: "a1", Title: "title 1", Desc: "desc 1", Content: "test content"},
		Article{AId: "a2", Title: "title 2", Desc: "desc 2", Content: "test content"},
	}
}
