package main
import (
	"net/http"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
	m "github.com/ajit-go/swaggerexample/mvc"
)

func handleServer(method string, path string, t *testing.T, f http.HandlerFunc, expectedStatus int) (rr *httptest.ResponseRecorder) {
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	//handler := http.HandlerFunc(f)

	f.ServeHTTP(rr, req)
	if status := rr.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, expectedStatus)
	}

	return
}

var cc = m.Controller{}

func TestHomePage(t *testing.T) {
	rr := handleServer("GET", "/", t, cc.Home, http.StatusOK)
	expected := "Home page"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	fmt.Println("TestHomePage")
}
func TestArticles(t *testing.T) {
	rr := handleServer("GET", "/as", t, cc.GetArticles, http.StatusOK)
	expected := `{"Articles":[{"title":"title 1","desc":"desc 1","content":"test content"},{"title":"title 2","desc":"desc 2","content":"test content"}]}`
	
	got := rr.Body.String()
	if strings.Trim(got , " \n") != strings.Trim(expected, " \n") {
		t.Errorf("got %v want %v",
			strings.Trim(rr.Body.String(), " "), expected)
	}
	fmt.Println("TestArticles")

}