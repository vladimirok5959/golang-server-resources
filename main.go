package main

import (
	"net/http"

	"golang-server-resources/resource"
)

func main() {
	res := resource.New()
	res.Add("some1.css", "text/css", []byte(`.some1{display:none}`))
	res.Add("some2.css", "text/css", []byte(`.some2{display:none}`))
	res.Add("some3.css", "text/css", []byte(`.some3{display:none}`))
	res.Add("some4.css", "text/css", []byte(`.some4{display:none}`))
	res.Add("some5.css", "text/css", []byte(`.some5{display:none}`))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Resource response
		if res.Response(w, r, func(w http.ResponseWriter, r *http.Request, i *resource.Resource) {
			w.Header().Set("Some-Header", "test")
			w.Header().Set("Cache-Control", "public, max-age=31536000")
		}, func(w http.ResponseWriter, r *http.Request, i *resource.Resource) {
			w.Write([]byte("\n\n/* Path: " + (*i).Path + " */"))
			w.Write([]byte("\n/* Ctype: " + (*i).Ctype + " */"))
		}) {
			return
		}

		// Default response
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
			<div>Hello World!</div>
			<div><a href="/some1.css">/some1.css</a></div>
			<div><a href="/some2.css">/some2.css</a></div>
			<div><a href="/some3.css">/some3.css</a></div>
			<div><a href="/some4.css">/some4.css</a></div>
			<div><a href="/some5.css">/some5.css</a></div>
		`))
	})

	http.ListenAndServe(":8080", nil)
}
