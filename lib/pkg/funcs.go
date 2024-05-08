package converter

import (
	"fmt"
	error2 "github.com/smart-libs/go-cc-converter/lib/pkg/error"
	"reflect"
)

type (
	// FromToTypedFunc is a strongly typed conversion function from a given F type value to the given T type that
	// should be assigned to the given pointer to T. The given pointer cannot be nil since it should have to receive
	// the converted value.
	FromToTypedFunc[F any, T any] func(F, *T) error

	// ToTypedFunc is a partially typed conversion function from any type to a specific type
	ToTypedFunc[T any] func(any, *T) error

	// FromToFunc is a dynamic type conversion function
	FromToFunc func(any, any) error

	// ToFunc is a dynamic type conversion function
	ToFunc func(any) (any, error)
)

func assertNotNilConverter[F any](converter F) F {
	var asAny any = converter
	if asAny == nil {
		panic(fmt.Errorf("invalid nil conversion function=[%T]", converter))
	}
	return converter
}

func MakeFromToFunc[F any, T any](converter func(F, *T) error) FromToFunc {
	assertNotNilConverter(converter)
	var f F
	return func(fAny any, tAny any) (err error) {
		if f, err = CastOrErr[F](fAny); err != nil {
			return
		}

		var pt *T
		if pt, err = CastOrErr[*T](tAny); err != nil {
			return
		}
		return converter(f, pt)
	}
}

func MakeFromToFuncWithReturn[F any, T any](converter func(F) (T, error)) FromToFunc {
	assertNotNilConverter(converter)
	var f F
	return func(fAny any, tAny any) (err error) {
		if f, err = CastOrErr[F](fAny); err != nil {
			return
		}

		t, err := converter(f)
		if err != nil {
			return err
		}
		pt, err := CastOrErr[*T](tAny)
		if err != nil {
			return err
		}
		*pt = t
		return nil
	}
}

func MakeToFunc[F any, T any](converter func(F) (T, error)) ToFunc {
	assertNotNilConverter(converter)
	var f F
	return func(fAny any) (tAny any, err error) {
		if f, err = CastOrErr[F](fAny); err != nil {
			return
		}
		return converter(f)
	}
}

func MakeToFuncWithArg[F any, T any](converter func(F, *T) error) ToFunc {
	assertNotNilConverter(converter)
	var f F
	return func(fAny any) (tAny any, err error) {
		var pt *T
		if f, err = CastOrErr[F](fAny); err != nil {
			return
		}
		err = converter(f, pt)
		return *pt, err
	}
}

func CastOrErr[T any](value any) (result T, err error) {
	if value == nil {
		return
	}
	var ok bool
	if result, ok = value.(T); ok {
		return
	}
	if valueOf, fallBackOK := value.(reflect.Value); fallBackOK {
		return CastOrErr[T](valueOf.Interface())
	}

	err = error2.NewConversionError(value, result, "illegal cast operation")
	return
}

func To[T any](registry Converters, f any) (T, error) {
	var t T
	err := registry.Convert(f, &t)
	return t, err
}
