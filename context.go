package services

import (
	"net/http"

	f "github.com/gopherlabs/gopher-framework"
	"github.com/gopherlabs/gopher-providers-mux/context"
)

type ContextProvider struct {
}

func (p ContextProvider) Register(c *f.Container, config interface{}) interface{} {
	return p
}

func (p ContextProvider) GetKey() string {
	return f.CONTEXT
}

func (p ContextProvider) Set(r *http.Request, key, val interface{}) {
	context.Set(r, key, val)
}

func (p ContextProvider) Get(r *http.Request, key interface{}) interface{} {
	return context.Get(r, key)
}

func (p ContextProvider) GetOk(r *http.Request, key interface{}) (interface{}, bool) {
	return context.GetOk(r, key)
}

func (p ContextProvider) GetAll(r *http.Request) map[interface{}]interface{} {
	return context.GetAll(r)
}

func (p ContextProvider) GetAllOk(r *http.Request) (map[interface{}]interface{}, bool) {
	return context.GetAllOk(r)
}

func (p ContextProvider) Delete(r *http.Request, key interface{}) {
	return context.Delete(r, key)
}

func (p ContextProvider) Clear(r *http.Request) {
	context.Clear(r)
}

func (p ContextProvider) Purge(maxAge int) int {
	return context.Purge(maxAge)
}
