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
	r.mux = mux.NewRouter()
	return r
}

func (r *RouteProvider) GetKey() string {
	return framework.ROUTER
}

func (r *RouteProvider) SubRouter(matcher f.GroupMatcher) f.Routable {
	sub := new(RouteProvider)
	group := r.mux.NewRoute()
	if matcher.PathPrefix != "" {
		group.PathPrefix(matcher.PathPrefix)
	}
	if matcher.Host != "" {
		group.Host(matcher.Host)
	}
	if len(matcher.Methods) > 0 {
		group.Methods(matcher.Methods...)
	}
	if len(matcher.Queries) > 0 {
		group.Queries(matcher.Queries...)
	}
	if len(matcher.Schemes) > 0 {
		group.Schemes(matcher.Schemes...)
	}
	sub.mux = group.Subrouter()
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
	for pathName, pathValue := range r.config.StaticDirs {
		pathPrefix := pathName
		if pathName != "/" {
			pathPrefix = pathPrefix + "/"
		}
		r.mux.PathPrefix(pathPrefix).Handler(http.StripPrefix(pathPrefix,
			http.FileServer(http.Dir(pathValue))))
	}
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

func (r *RouteProvider) Static(path string, dir string) {
	r.config.StaticDirs[path] = dir
}
