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

func (r RenderProvider) View(rw http.ResponseWriter, status int, name string, binding interface{}) {
	r.render.HTML(rw, status, name, binding)
}
