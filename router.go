package services

import (
	"net/http"

	"github.com/gopherlabs/gopher-framework"
	f "github.com/gopherlabs/gopher-framework"
	"github.com/gopherlabs/gopher-providers-mux"
)

type RouteProvider struct {
	http.Handler
	mux *mux.Router
}

func (r *RouteProvider) Register(config map[string]interface{}) interface{} {
	return r
}

func (r *RouteProvider) GetKey() string {
	return framework.ROUTER
}

func (r *RouteProvider) NewRouter() f.Routable {
	r.mux = mux.NewRouter()
	return r
}

func (r *RouteProvider) SubRouter() f.Routable {
	sub := new(RouteProvider)
	sub.mux = r.mux.PathPrefix("/products").Subrouter()
	return sub
}

func (r *RouteProvider) Get(path string, fn f.HandlerFn, mw ...f.MiddlewareHandler) {
	r.mux.HandleFunc(path, fn).Methods("GET")
}

func (r *RouteProvider) Head(path string, fn f.HandlerFn, mw ...f.MiddlewareHandler) {
	r.mux.HandleFunc(path, fn).Methods("HEAD")
}

func (r *RouteProvider) Post(path string, fn f.HandlerFn, mw ...f.MiddlewareHandler) {
	r.mux.HandleFunc(path, fn).Methods("POST")
}

func (r *RouteProvider) Put(path string, fn f.HandlerFn, mw ...f.MiddlewareHandler) {
	r.mux.HandleFunc(path, fn).Methods("PUT")
}

func (r *RouteProvider) Patch(path string, fn f.HandlerFn, mw ...f.MiddlewareHandler) {
	r.mux.HandleFunc(path, fn).Methods("PATCH")
}

func (r *RouteProvider) Delete(path string, fn f.HandlerFn, mw ...f.MiddlewareHandler) {
	r.mux.HandleFunc(path, fn).Methods("DELETE")
}

func (r *RouteProvider) Options(path string, fn f.HandlerFn, mw ...f.MiddlewareHandler) {
	r.mux.HandleFunc(path, fn).Methods("OPTIONS")
}

func (r *RouteProvider) Match(path string, fn f.HandlerFn, verbs []string, mw ...f.MiddlewareHandler) {
	r.mux.HandleFunc(path, fn).Methods(verbs...)
}

func (r *RouteProvider) All(path string, fn f.HandlerFn, mw ...f.MiddlewareHandler) {
	r.mux.HandleFunc(path, fn)
}

func (r *RouteProvider) NotFound(fn f.HandlerFn, mw ...f.MiddlewareHandler) {
	r.mux.NotFoundHandler = http.HandlerFunc(fn)
}

func (r *RouteProvider) Serve() {
	http.ListenAndServe("0.0.0.0:3000", r.mux)
}

func (r *RouteProvider) Vars(req *http.Request) map[string]string {
	return mux.Vars(req)
}

func (r *RouteProvider) Use(mw f.MiddlewareHandler, args ...interface{}) {
}
