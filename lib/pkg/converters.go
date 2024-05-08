package converter

import (
	"reflect"
)

type (
	// Converters is an entity that knows how to convert one value type into another value type.
	Converters interface {
		Convert(from any, to any) error
		ConvertToType(from any, toType reflect.Type) (any, error)
	}
)
