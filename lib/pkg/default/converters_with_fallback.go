package converterdefault

import (
	"fmt"
	"github.com/smart-libs/go-cc-converter/lib/pkg"
	convertererror "github.com/smart-libs/go-cc-converter/lib/pkg/error"
	"github.com/smart-libs/go-cc-converter/lib/pkg/fallback"
	"reflect"
)

type (
	// defaultConverters is a repository of conversion functions
	defaultConverters struct {
		options          FallbackOptions
		registry         converter.Registry
		fallbackRegistry fromToMap[fallback.FromToFunc]
	}
)

// not planned to be exposed so far
func newConvertersWithFallback(registry converter.Registry) converter.Converters {
	return newConvertersWithFallbackOptions(registry, defaultFallbackOptions)
}

// not planned to be exposed so far
func newConvertersWithFallbackOptions(registry converter.Registry, options FallbackOptions) converter.Converters {
	if registry == nil {
		panic(fmt.Errorf("%T cannot be nil", registry))
	}
	return &defaultConverters{
		registry:         registry,
		options:          options,
		fallbackRegistry: newFromToMap[fallback.FromToFunc](),
	}
}

func (r *defaultConverters) ConvertToType(from any, toType reflect.Type) (any, error) {
	targetValue := reflect.New(toType)
	err := r.Convert(from, targetValue.Interface())
	if err != nil {
		return nil, err
	}
	return targetValue.Elem().Interface(), nil
}

func (r *defaultConverters) Convert(from any, to any) error {
	var newError = convertererror.NewConversionError
	if to == nil {
		return newError(from, to, "the 'to' argument is nil")
	}

	var (
		fValue  = getAsValueOf(from)
		pTValue = getAsValueOf(to)
	)

	if pTValue.Kind() != reflect.Pointer {
		return newError(from, to, "the to argument is not a pointer")
	}
	if pTValue.IsNil() {
		return newError(from, to, "the to argument is not nil pointer")
	}

	if from == nil {
		// init the pointer to 'To'
		pTValue.Elem().Set(reflect.Zero(pTValue.Type().Elem()))
		return nil
	}

	fType := fValue.Type()
	tType := pTValue.Type().Elem()
	if conversionFunc := r.findConversionFor(fType, tType); conversionFunc != nil {
		return conversionFunc(fValue, pTValue)
	}
	return r.options.OnNoMoreFallbacks(from, to)
}

var anyType = reflect.TypeFor[any]()

func (r *defaultConverters) findConversionFor(fType, tType reflect.Type) fallback.FromToFunc {
	if tType.Kind() == reflect.Interface {
		tType = tType.Elem()
	}
	if fromToFunc := r.registry.GetConverter(fType, tType); fromToFunc != nil {
		return func(fromVal, toVal reflect.Value) error {
			return fromToFunc(fromVal.Interface(), toVal.Interface())
		}
	}
	if fType != anyType { // for optimization we skip any type because it can exist only when registered in the registry repo.
		if !r.options.DisableFallbackRegistry {
			if fallbackFromToFunc, found := r.fallbackRegistry.Get(fType, tType); found {
				return fallbackFromToFunc
			}
		}
		for _, tryFallback := range r.options.ConversionFallbackList {
			if fallbackFromToFunc := tryFallback(r.findConversionFor, fType, tType); fallbackFromToFunc != nil {
				return fallbackFromToFunc
			}
		}
	}
	return nil
}

func getAsValueOf(v any) reflect.Value {
	if asValueOf, ok := v.(reflect.Value); ok {
		return asValueOf
	}
	return reflect.ValueOf(v)
}
