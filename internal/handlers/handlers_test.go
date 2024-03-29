package handlers

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"net/url"
// 	"testing"
// )

// type postData struct {
// 	key   string
// 	value string
// }

// var theTests = []struct {
// 	name               string
// 	url                string
// 	method             string
// 	params             []postData
// 	expectedStatusCode int
// }{
// 	{"home", "/", "GET", []postData{}, http.StatusOK},
// 	{"about", "/about", "GET", []postData{}, http.StatusOK},
// 	{"generals", "/generals", "GET", []postData{}, http.StatusOK},
// 	{"majors", "/majors", "GET", []postData{}, http.StatusOK},
// 	{"search-availability", "/search-availability", "GET", []postData{}, http.StatusOK},
// 	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
// 	{"make-res", "/reservation", "GET", []postData{}, http.StatusOK},
// 	{"post-search-availability", "/search-availability", "Post", []postData{
// 		{key: "start", value: "2023-01-01"},
// 		{key: "end", value: "2023-01-02"},
// 	}, http.StatusOK},
// 	{"post-search-availability-json", "/search-availability-json", "Post", []postData{
// 		{key: "start", value: "2023-01-01"},
// 		{key: "end", value: "2023-01-02"},
// 	}, http.StatusOK},
// 	{"reservation", "/reservation ", "Post", []postData{
// 		{key: "first_name", value: "Simba"},
// 		{key: "last_name", value: "Yang"},
// 		{key: "email", value: "simba@email.com"},
// 		{key: "phone", value: "555-555-5555"},
// 	}, http.StatusOK},
// }

// func TestHandlers(t *testing.T) {
// 	routes := getRoutes()

// 	ts := httptest.NewTLSServer(routes)
// 	defer ts.Close()

// 	for _, e := range theTests {
// 		if e.method == "GET" {
// 			resp, err := ts.Client().Get(ts.URL + e.url)
// 			if err != nil {
// 				t.Fatal("Failed:", e.name, err)

// 			}

// 			if resp.StatusCode != e.expectedStatusCode {
// 				t.Errorf("for %s expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
// 			}
// 		} else {
// 			values := url.Values{}
// 			for _, x := range e.params {
// 				values.Add(x.key, x.value)
// 			}

// 			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
// 			if err != nil {
// 				t.Fatal("Failed:", e.name, err)
// 			}

// 			if resp.StatusCode != e.expectedStatusCode {
// 				t.Errorf("for %s expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
// 			}

// 		}
// 	}
// }
