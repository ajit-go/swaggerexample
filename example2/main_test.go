package main
import (
	"net/http"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
)

func handleServer(method string, path string, t *testing.T, f http.HandlerFunc, expectedStatus int) (responseRecorder *httptest.ResponseRecorder) {
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder = httptest.NewRecorder()

	f.ServeHTTP(responseRecorder, req)
	if status := responseRecorder.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, expectedStatus)
	}

	return
}

func TestArticle(t *testing.T) {
	rr := handleServer("GET", "/article", t, getArticle, http.StatusOK)
	expected := `{"aid":"","title":"sonos one","desc":"sonos one smart speaker","content":"Speakers with alexa and google home"}`
	
	got := rr.Body.String()
	if strings.Trim(got , " \n") != strings.Trim(expected, " \n") {
		t.Errorf("got %v want %v",
			strings.Trim(rr.Body.String(), " "), expected)
	}
	fmt.Println("---Test  Article---")
}