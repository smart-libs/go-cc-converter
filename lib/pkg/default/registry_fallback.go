package converterdefault

import (
	"fmt"
	"github.com/smart-libs/go-cc-converter/lib/pkg"
	convertererror "github.com/smart-libs/go-cc-converter/lib/pkg/error"
	"github.com/smart-libs/go-cc-converter/lib/pkg/fallback"
	"reflect"
)

type (

	// FallbackOptions suggest options to be considered by the fallback implementation
	FallbackOptions struct {
		// OnNoMoreFallbacks is used when the default Converters (converterRegistry) has no more fallbacks to use.
		// You can replace this behavior.
		OnNoMoreFallbacks converter.FromToFunc

		// ConversionFallbackList is the default list of TryNewConversionFunc used when there is no registered converter to use.
		ConversionFallbackList []fallback.TryNewConversionFunc

		DisableFallbackRegistry bool
	}

	// fallbackRegistryDecorator is a repository of conversion functions
	fallbackRegistryDecorator struct {
		decorated        converter.Registry
		options          FallbackOptions
		fallbackRegistry fromToMap[fallback.FromToFunc]
	}
)

var (
	defaultFallbackOptions = FallbackOptions{
		// OnNoMoreFallbacks default is to return error
		OnNoMoreFallbacks: func(from any, to any) error { return convertererror.NewConversionNotFoundError(from, to) },

		// ConversionFallbackList has the default list of TryNewConversionFunc
		ConversionFallbackList: []fallback.TryNewConversionFunc{
			fallback.FromTypeEqualsToTypeFallback,
			fallback.ConvertibleToFallback,
			fallback.ToArrayTypeFallback,
			fallback.PointerToPointerToTReflection,
			fallback.PointerFallback,
			fallback.FromAnyFallback, // search first for string type, last option is any
		},
	}
)

func NewFallbackRegistry(registry converter.Registry) converter.Registry {
	return NewFallbackRegistryWithOptions(registry, defaultFallbackOptions)
}

func NewFallbackRegistryWithOptions(registry converter.Registry, options FallbackOptions) converter.Registry {
	if registry == nil {
		panic(fmt.Errorf("%T cannot be nil", registry))
	}
	return &fallbackRegistryDecorator{
		decorated:        registry,
		options:          options,
		fallbackRegistry: newFromToMap[fallback.FromToFunc](),
	}
}

func (r *fallbackRegistryDecorator) Add(handler converter.Handler) *converter.Handler {
	return r.decorated.Add(handler)
}

func (d *fallbackRegistryDecorator) GetConverter(from, to reflect.Type) converter.FromToFunc {
	result := d.findConversionFor(from, to)
	if result != nil {
		return func(f any, t any) error {
			return result(getAsValueOf(f), getAsValueOf(t))
		}
	}
	return nil
}

func (d *fallbackRegistryDecorator) GetConverterTo(from, to reflect.Type) converter.ToFunc {
	result := d.decorated.GetConverter(from, to)
	if result != nil {
		return func(f any) (any, error) {
			ptrTo := reflect.New(to)
			err := result(f, ptrTo.Interface())
			return ptrTo.Elem().Interface(), err
		}
	}
	return nil
}

func (d *fallbackRegistryDecorator) findConversionFor(fType, tType reflect.Type) fallback.FromToFunc {
	if mainFunc := d.decorated.GetConverter(fType, tType); mainFunc != nil {
		return func(fromVal, ptrToVal reflect.Value) error {
			return mainFunc(fromVal.Interface(), ptrToVal.Interface())
		}
	}

	if fType != anyType { // for optimization we skip any type because it can exist only when registered in the registry repo.
		if !d.options.DisableFallbackRegistry {
			if fallbackFromToFunc, found := d.fallbackRegistry.Get(fType, tType); found {
				return fallbackFromToFunc
			}
		}
		for _, tryFallback := range d.options.ConversionFallbackList {
			if fallbackFromToFunc := tryFallback(d.findConversionFor, fType, tType); fallbackFromToFunc != nil {
				return fallbackFromToFunc
			}
		}
	}
	return nil
}
