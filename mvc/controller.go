package mvc

import (
	"encoding/json"
	"fmt"
	. "github.com/ajit-go/swaggerexample/model"
	"net/http"
)

//Controller struct to hold all rest handler fucntions
type Controller struct{}

var data = Data{}

//Home page
// @Summary Home page
// @Description Home page
// @Accept  json
// @Produce  json
// @Param q query string false "name search by q"
// @Success 200 {object} model.Hello
// @Header 200 {string} Token "qwerty"
// Failure 400 {object}
// @Failure 404 {object} model.Error
// Failure 500 {object}
// @Router / [get]
func (c *Controller) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

//GetArticles page
// @Summary list Articles
// @Description get Articles
// @Accept  json
// @Produce  json
// @Param name header string true "name search by q"
// @Param q query string true "name search by q"
// @Success 200 {array} model.Article
// @Header 200 {string} Token "qwerty"
// Failure 400 {object}
// @Failure 404 {string} model.Error
// Failure 500 {object}
// @Router /articles [get]
// @Security BasicAuth
func (c *Controller) GetArticles(w http.ResponseWriter, r *http.Request) {
	initialize()
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(data)
}

func initialize() {
	data.Articles = []Article{
		Article{Title: "title 1", Desc: "desc 1", Content: "test content"},
		Article{Title: "title 2", Desc: "desc 2", Content: "test content"},
	}
}
