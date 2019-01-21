package resource

import (
	"net/http"
)

type Resource struct {
	Path  string
	Ctype string
	bytes []byte
}

type resource struct {
	list []Resource
}

func New() *resource {
	r := resource{}
	r.list = make([]Resource, 0)
	return &r
}

func (this *resource) Add(path string, ctype string, bytes []byte) {
	this.list = append(this.list, Resource{
		Path:  path,
		Ctype: ctype,
		bytes: bytes,
	})
}

func (this *resource) Response(w http.ResponseWriter, r *http.Request, before func(w http.ResponseWriter, r *http.Request, i *Resource), after func(w http.ResponseWriter, r *http.Request, i *Resource)) bool {
	for _, value := range this.list {
		if r.URL.Path == "/"+value.Path {
			if before != nil {
				before(w, r, &value)
			}
			w.Header().Set("Content-Type", value.Ctype)
			w.Write(value.bytes)
			if after != nil {
				after(w, r, &value)
			}
			return true
		}
	}
	return false
}
