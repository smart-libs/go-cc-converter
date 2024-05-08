package convertererror

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type (
	NotFoundError struct {
		error
	}
)

var (
	NewConversionNotFoundError = func(from, to any) error {
		return NotFoundError{error: fmt.Errorf("no conversion function found from=[%T] to=[%T]", from, to)}
	}

	NewConversionError = func(from, to any, msg string) error {
		return fmt.Errorf("ConversionError: from %T=[%v] to %T=[%v] due to %s",
			extractValueOf(from), from, extractValueOf(to), to, msg)
	}
	NewConversionErrorWithCause = func(from, to any, errs ...error) error {
		const mask = "ConversionError: from %T=[%v] to %T=[%v] due to %w"
		switch len(errs) {
		case 0:
			return NewConversionError(from, to, "unknown reason")
		case 1:
			return fmt.Errorf(mask, extractValueOf(from), from, extractValueOf(to), to, errs[0])
		default:
			builder := strings.Builder{}
			builder.WriteString("multiple errors:\n")
			for _, e := range errs {
				builder.WriteString(e.Error())
				builder.WriteString("\n")
			}
			return fmt.Errorf(mask, extractValueOf(from), from, extractValueOf(to), to, errors.New(builder.String()))
		}
	}
)

func extractValueOf(v any) any {
	if valueOf, ok := (v).(reflect.Value); ok {
		return valueOf.Interface()
	}
	return v
}
