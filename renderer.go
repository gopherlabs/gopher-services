package services

import (
	"net/http"

	"github.com/gopherlabs/gopher-framework"
	"github.com/gopherlabs/gopher-providers-render"
)

type RenderProvider struct {
	render *render.Render
}

func (r RenderProvider) Register(config map[string]interface{}) interface{} {
	r.render = render.New()
	return r
}

func (l RenderProvider) GetKey() string {
	return framework.RENDERER
}

func (r RenderProvider) Data(rw http.ResponseWriter, status int, data []byte) {
	r.render.Data(rw, status, data)
}

func (r RenderProvider) Text(rw http.ResponseWriter, data string, status ...int) {
	if len(status) > 0 {
		r.render.Text(rw, status[0], data)
	} else {
		r.render.Text(rw, http.StatusOK, data)
	}
}

func (r RenderProvider) JSON(rw http.ResponseWriter, status int, data interface{}) {
	r.render.JSON(rw, status, data)
}

func (r RenderProvider) JSONP(rw http.ResponseWriter, status int, callback string, data interface{}) {
	r.render.JSONP(rw, status, callback, data)
}

func (r RenderProvider) XML(rw http.ResponseWriter, status int, data interface{}) {
	r.render.XML(rw, status, data)
}

func (r RenderProvider) View(rw http.ResponseWriter, status int, name string, binding interface{}) {
	r.render.HTML(rw, status, name, binding)
}
