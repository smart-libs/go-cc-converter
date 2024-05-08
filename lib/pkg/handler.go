package converter

import (
	"reflect"
)

type (
	// Handler is a data structure that provides input type definitions to invoke a FromToFunc function
	Handler struct {
		fromType, toType reflect.Type
		fromToFunc       FromToFunc
		toFunc           ToFunc
	}
)

func (h Handler) GetFromType() reflect.Type { return h.fromType }
func (h Handler) GetToType() reflect.Type   { return h.toType }
func (h Handler) GetFromToFunc() FromToFunc { return h.fromToFunc }
func (h Handler) GetToFunc() ToFunc         { return h.toFunc }

func CreateHandler[From any, To any](converter func(From, *To) error) Handler {
	var (
		f From
		t To
	)
	return Handler{
		fromType:   reflect.TypeOf(f),
		toType:     reflect.TypeOf(t),
		fromToFunc: MakeFromToFunc(converter),
		toFunc:     MakeToFuncWithArg(converter),
	}
}

func CreateHandlerWithReturn[From any, To any](converter func(From) (To, error)) Handler {
	var (
		f From
		t To
	)
	return Handler{
		fromType:   reflect.TypeOf(f),
		toType:     reflect.TypeOf(t),
		fromToFunc: MakeFromToFuncWithReturn(converter),
		toFunc:     MakeToFunc(converter),
	}
}

func CreateHandlerWithReturnNoError[From any, To any](converter func(From) To) Handler {
	return CreateHandlerWithReturn[From, To](func(from From) (To, error) { return converter(from), nil })
}
