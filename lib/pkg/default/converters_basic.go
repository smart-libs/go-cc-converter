package converterdefault

import (
	"fmt"
	"github.com/smart-libs/go-cc-converter/lib/pkg"
	convertererror "github.com/smart-libs/go-cc-converter/lib/pkg/error"
	"reflect"
)

type (
	// basicConverters is a repository of conversion functions
	basicConverters struct {
		registry converter.Registry
	}
)

func NewBasicConverters(registry converter.Registry) converter.Converters {
	if registry == nil {
		panic(fmt.Errorf("%T cannot be nil", registry))
	}
	return &basicConverters{
		registry: registry,
	}
}

func (r *basicConverters) ConvertToType(from any, toType reflect.Type) (any, error) {
	targetValue := reflect.New(toType)
	err := r.Convert(from, targetValue.Interface())
	if err != nil {
		return nil, err
	}
	return targetValue.Elem().Interface(), nil
}

func (r *basicConverters) Convert(from any, to any) error {
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
	if fromToFunc := r.registry.GetConverter(fType, tType); fromToFunc != nil {
		return fromToFunc(from, to)
	}
	return convertererror.NewConversionNotFoundError(from, to)
}
