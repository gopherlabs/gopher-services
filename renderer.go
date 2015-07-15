package services

import (
	"net/http"

	f "github.com/gopherlabs/gopher-framework"
	"github.com/gopherlabs/gopher-providers-render"
)

type RenderProvider struct {
	render *render.Render
}

func (r RenderProvider) Register(c *f.Container, config interface{}) interface{} {
	r.render = render.New()
	return r
}

func (l RenderProvider) GetKey() string {
	return f.RENDERER
}

func getStatus(status []int) int {
	if len(status) > 0 {
		return status[0]
	} else {
		return http.StatusOK
	}
}

func (r RenderProvider) Data(rw http.ResponseWriter, data []byte, status ...int) {
	r.render.Data(rw, getStatus(status), data)
}

func (r RenderProvider) Text(rw http.ResponseWriter, data string, status ...int) {
	r.render.Text(rw, getStatus(status), data)
}

func (r RenderProvider) JSON(rw http.ResponseWriter, data interface{}, status ...int) {
	r.render.JSON(rw, getStatus(status), data)
}

func (r RenderProvider) JSONP(rw http.ResponseWriter, callback string, data interface{}, status ...int) {
	r.render.JSONP(rw, getStatus(status), callback, data)
}

func (r RenderProvider) XML(rw http.ResponseWriter, data interface{}, status ...int) {
	r.render.XML(rw, getStatus(status), data)
}

func (r RenderProvider) View(rw http.ResponseWriter, name string, binding interface{}, status ...int) {
	r.render.HTML(rw, getStatus(status), name, binding)
}
