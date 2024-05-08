package converter

import (
	"fmt"
	"github.com/smart-libs/go-cc-converter/lib/pkg/fallback"
	"reflect"
)

type (
	// Registry is the entity responsible for keep the FromToFunc instances classified by the `from` and `to` types.
	// Because Golang only supports the Generic feature per structure and not per structure method, we will use functions
	// to register the FromToFunc instances ensuring that Registry knows the `from` and `to` types.
	Registry interface {
		// Add adds a new converter to the registry. If there is already one, then the old one is returned and replaced by the new one.
		Add(handler Handler) *Handler

		// GetConverter returns the FromToFunc to be used to convert a `from` value to a `to` value pointer
		GetConverter(from, to reflect.Type) FromToFunc

		// GetConverterTo returns the ToFunc to be used to convert a `from` value to a `to` value
		GetConverterTo(from, to reflect.Type) ToFunc
	}
)

func assertNotNilRegistry(r Registry) Registry {
	if r == nil {
		panic(fmt.Errorf("invalid nil Registry"))
	}
	return r
}

func AddHandler[From any, To any](registry Registry, fromToFunc func(From, *To) error) {
	assertNotNilRegistry(registry).Add(CreateHandler[From, To](fromToFunc))
	AddFallbacks[From, To](registry, fromToFunc)
}

func AddHandlerWithReturn[From any, To any](registry Registry, toFunc func(From) (To, error)) {
	assertNotNilRegistry(registry).Add(CreateHandlerWithReturn[From, To](toFunc))
	AddFallbacks[From, To](registry, func(from From, to *To) (err error) {
		*to, err = toFunc(from)
		return err
	})
}

func AddHandlerWithReturnNoError[From any, To any](registry Registry, toFunc func(From) To) {
	assertNotNilRegistry(registry).Add(CreateHandlerWithReturnNoError[From, To](toFunc))
	AddFallbacks[From, To](registry, func(from From, to *To) (err error) {
		*to = toFunc(from)
		return nil
	})
}

func AddFallbacks[From any, To any](registry Registry, fromToFunc func(From, *To) error) {
	registry.Add(CreateHandler[From, *To](fallback.PointerToPointerToT(fromToFunc)))       // func(F, **T) error
	registry.Add(CreateHandler[*From, To](fallback.FromPointerToFToT(fromToFunc)))         // func(*F, *T) error
	registry.Add(CreateHandler[*From, *To](fallback.FromPointerToFToPointerT(fromToFunc))) // func(*F, **T) error
	registry.Add(CreateHandler[From, []To](fallback.FtoArrayOfT(fromToFunc)))              // func(F, *[]T) error
	registry.Add(CreateHandler[*From, []To](fallback.PointerToFtoArrayOfT(fromToFunc)))    // func(*F, *[]T) error
	registry.Add(CreateHandler[[]From, []To](fallback.ArrayOfFtoArrayOfT(fromToFunc)))     // func([]F, *[]T) error

	registry.Add(CreateHandler[From, From](fallback.FromFtoF[From]()))                    // func(F, *F) error
	registry.Add(CreateHandler[*From, *From](fallback.FromFtoF[*From]()))                 // func(*F, **F) error
	registry.Add(CreateHandler[From, *From](fallback.FromFtoPointerToF[From]()))          // func(F, **F) error
	registry.Add(CreateHandler[*From, From](fallback.FromPointerToFtoF[From]()))          // func(*F, *F) error
	registry.Add(CreateHandler[From, []From](fallback.FromFtoArrayOfF[From]()))           // func(F, *[]F) error
	registry.Add(CreateHandler[*From, []From](fallback.FromPointerToFtoArrayOfF[From]())) // func(*F, *[]F) error
	registry.Add(CreateHandler[[]From, []From](fallback.FromFtoF[[]From]()))              // func([]F, *[]F) error

	registry.Add(CreateHandler[To, To](fallback.FromFtoF[To]()))                    // func(F, *F) error
	registry.Add(CreateHandler[*To, *To](fallback.FromFtoF[*To]()))                 // func(*F, **F) error
	registry.Add(CreateHandler[To, *To](fallback.FromFtoPointerToF[To]()))          // func(F, **F) error
	registry.Add(CreateHandler[*To, To](fallback.FromPointerToFtoF[To]()))          // func(*F, *F) error
	registry.Add(CreateHandler[To, []To](fallback.FromFtoArrayOfF[To]()))           // func(F, *[]F) error
	registry.Add(CreateHandler[*To, []To](fallback.FromPointerToFtoArrayOfF[To]())) // func(*F, *[]F) error
	registry.Add(CreateHandler[[]To, []To](fallback.FromFtoF[[]To]()))              // func([]F, *[]F) error
}
