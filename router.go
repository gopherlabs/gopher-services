package services

import (
	"net/http"

	"strconv"

	"github.com/gopherlabs/gopher-framework"
	f "github.com/gopherlabs/gopher-framework"
	"github.com/gopherlabs/gopher-providers-mux"
)

type RouteProvider struct {
	http.Handler
	mux       *mux.Router
	config    f.ConfigRouter
	container *f.Container
}

func (r *RouteProvider) Register(c *f.Container, config interface{}) interface{} {
	r.container = c
	r.config = config.(f.ConfigRouter)
	//	c.Log.Info("|   > Host: %s", r.config.Host)
	//	c.Log.Info("|   > Port: %d", r.config.Port)
	r.mux = mux.NewRouter()
	return r
}

func (r *RouteProvider) GetKey() string {
	return framework.ROUTER
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
	http.ListenAndServe(r.config.Host+":"+strconv.Itoa(r.config.Port), r.mux)
}

func (r *RouteProvider) Vars(req *http.Request) map[string]string {
	return mux.Vars(req)
}

func (r *RouteProvider) Var(req *http.Request, name string) string {
	return mux.Vars(req)[name]
}

func (r *RouteProvider) Use(mw f.MiddlewareHandler, args ...interface{}) {
}
