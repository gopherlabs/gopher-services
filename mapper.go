package services

import (
	"github.com/gopherlabs/gopher-framework"
	cmap "github.com/gopherlabs/gopher-providers-map"
)

type MapProvider struct {
	cmap cmap.ConcurrentMap
}

func (p MapProvider) Register(config map[string]interface{}) interface{} {
	p.cmap = cmap.New()
	return p
}

func (p MapProvider) GetKey() string {
	return framework.MAPPER
}

func (p MapProvider) Get(key string) (value interface{}) {
	value, _ = p.cmap.Get(key)
	return
}

func (p MapProvider) Has(key string) bool {
	return p.cmap.Has(key)
}

func (p MapProvider) Set(key string, value interface{}) {
	p.cmap.Set(key, value)
}

func (p MapProvider) Remove(key string) {
	p.cmap.Remove(key)
}
