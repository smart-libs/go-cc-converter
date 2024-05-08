package converterdefault

import (
	"github.com/smart-libs/go-cc-converter/lib/pkg"
	"reflect"
)

type (
	defaultRegistry struct {
		fromToMap[converter.Handler]
	}
)

var (
	Registry = NewRegistry()
)

func NewRegistry() converter.Registry {
	return &defaultRegistry{fromToMap: newFromToMap[converter.Handler]()}
}

func (d *defaultRegistry) Add(handler converter.Handler) (old *converter.Handler) {
	return d.fromToMap.Add(handler.GetFromType(), handler.GetToType(), handler)
}

func (d *defaultRegistry) GetConverter(from, to reflect.Type) converter.FromToFunc {
	if current, found := d.fromToMap.Get(from, to); found {
		return current.GetFromToFunc()
	}

	return nil
}

func (d *defaultRegistry) GetConverterTo(from, to reflect.Type) converter.ToFunc {
	if current, found := d.fromToMap.Get(from, to); found {
		return current.GetToFunc()
	}

	return nil
}

func AddHandler[From any, To any](fromToFunc func(From, *To) error) {
	converter.AddHandler[From, To](Registry, fromToFunc)
}

func AddHandlerWithReturn[From any, To any](toFunc func(From) (To, error)) {
	converter.AddHandlerWithReturn[From, To](Registry, toFunc)
}

func AddHandlerWithReturnNoError[From any, To any](toFunc func(From) To) {
	converter.AddHandlerWithReturnNoError[From, To](Registry, toFunc)
}
