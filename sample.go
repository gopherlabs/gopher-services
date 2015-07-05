package services

import (
	"github.com/gopherlabs/gopher-framework"
)

type SampleProvider struct {
	name string
	log  *framework.Loggable
}

func (p *SampleProvider) Register(config map[string]interface{}) interface{} {
	return p
}

func (p *SampleProvider) GetKey() string {
	return "SAMPLE"
}

func (p *SampleProvider) NewSample() framework.Samplable {
	return new(SampleProvider)
}

func (p *SampleProvider) GetName() string {
	return "||" + p.name
}

func (p *SampleProvider) SetName(name string) {
	p.name = name
}
