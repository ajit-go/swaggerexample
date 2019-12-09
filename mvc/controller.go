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


func (c *Controller) Home(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Hello{"home of api"})
}


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
