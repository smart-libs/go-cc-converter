package converter

import (
	"errors"
	"fmt"
	convertererror "github.com/smart-libs/go-cc-converter/lib/pkg/error"
	"reflect"
)

type (
	// ConvertersList when you have more than one registry and want to use the first one that works
	ConvertersList []Converters
)

// NewConvertersList returns a Converters implementation that uses the given list of converters to find
// a function to convert values. If no Converters is provided, the the converter.Defauldefault is used
func NewConvertersList(converters ...Converters) Converters {
	if len(converters) == 0 {
		panic(fmt.Errorf("no Converters were provided"))
	}
	return ConvertersList(converters)
}

func (l ConvertersList) ConvertToType(from any, toType reflect.Type) (any, error) {
	var firstErr error
	for _, converterElem := range l {
		convertedValue, err := converterElem.ConvertToType(from, toType)
		if err == nil {
			return convertedValue, nil
		}
		if firstErr == nil {
			firstErr = err
		}
	}

	return firstErr, nil
}

func (l ConvertersList) Convert(from any, to any) error {
	var lastErr error
	for _, converterElem := range l {
		if err := converterElem.Convert(from, to); err == nil {
			return nil
		} else {
			var notFound convertererror.NotFoundError
			if !errors.As(err, &notFound) {
				lastErr = err
			}
		}
	}

	if lastErr == nil {
		return convertererror.NewConversionNotFoundError(from, to)
	}
	return lastErr
}
