package resource

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func handle() http.Handler {
	res := New()
	res.Add("some1.css", "text/css", []byte(`.some1{display:none}`))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if res.Response(w, r, func(w http.ResponseWriter, r *http.Request, i *OneResource) {
			w.Header().Set("Some-Header", "test")
			w.Header().Set("Cache-Control", "public, max-age=31536000")
		}, func(w http.ResponseWriter, r *http.Request, i *OneResource) {
			w.Write([]byte("\n\n/* Path: " + (*i).Path + " */"))
			w.Write([]byte("\n/* Ctype: " + (*i).Ctype + " */"))
		}) {
			return
		}
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`Hello World!`))
	})
}

func request(t *testing.T, file string) *httptest.ResponseRecorder {
	request, err := http.NewRequest("GET", file, nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handle().ServeHTTP(recorder, request)
	return recorder
}

func TestHtml(t *testing.T) {
	r := request(t, "/")
	if s := r.Code; s != http.StatusOK {
		t.Fatalf("handler return wrong status code: got (%v) want (%v)", s, http.StatusOK)
	}
	if c := r.Header().Get("Content-Type"); c != "text/html" {
		t.Fatalf("content type header not match: got (%v) want (%v)", c, "text/html")
	}
	if r.Body.String() != "Hello World!" {
		t.Fatalf("bad body response, not match")
	}
}

func TestResource(t *testing.T) {
	r := request(t, "/some1.css")
	if s := r.Code; s != http.StatusOK {
		t.Fatalf("handler return wrong status code: got (%v) want (%v)", s, http.StatusOK)
	}
	if c := r.Header().Get("Content-Type"); c != "text/css" {
		t.Fatalf("content type header not match: got (%v) want (%v)", c, "text/css")
	}
	if r.Body.String() != ".some1{display:none}\n\n/* Path: some1.css */\n/* Ctype: text/css */" {
		t.Fatalf("bad body response, not match")
	}
}
