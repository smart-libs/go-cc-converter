package fallback

import (
	"reflect"
)

type (
	// FromToFunc is a conversion function where the from and to values are provided as wrapped by reflection.Value instances.
	FromToFunc func(fromVal, ptrToVal reflect.Value) error

	// ConversionFinderFunc is the signature of the find argument provided to TryNewConversionFunc implementations.
	// In theory is should be used to get a conversion function registered in the main repo or in the fallback repo.
	// The implementation shall return nil if no conversion function is found.
	ConversionFinderFunc = func(from, to reflect.Type) FromToFunc

	// TryNewConversionFunc is a function invoked by Converters when no registered converter function is found.
	// The 'from' and 'ptrTo' arguments are the original 'from' and `to` values wrapped as reflect.Value instances
	// which were already created by Converters in its try to find a conversion function.
	// The TryNewConversionFunc function must return nil if it was not possible to derive a conversion function using the
	// given input.
	// To additional validation that should have been performed by the Converters implementation, the TryNewConversionFunc
	// shall assume the from argument is not nil or the default reflection.Value instance, and the ptrTo argument is a
	// pointer to a memory to hold the converted value. The find argument never is nil and can be used to get previously
	// stored conversion functions.
	TryNewConversionFunc func(find ConversionFinderFunc, fromType reflect.Type, ptrToType reflect.Type) FromToFunc
)

// ConvertibleToFallback like (from string to *string) or (from *string to *string)
func ConvertibleToFallback(_ ConversionFinderFunc, fType, tType reflect.Type) FromToFunc {
	if fType.ConvertibleTo(tType) {
		return func(from reflect.Value, ptrTo reflect.Value) error {
			// tempVar := from.(*To)
			convertedValue := from.Convert(tType)
			ptrTo.Elem().Set(convertedValue)
			return nil
		}
	}
	// if tType is a pointer, so try convertible to the target type without pointer
	if tType.Kind() == reflect.Pointer && fType.ConvertibleTo(tType.Elem()) {
		return func(from reflect.Value, ptrTo reflect.Value) error {
			// tempVar := from.(To)
			convertedValue := from.Convert(tType.Elem())
			ptrTo.Elem().Set(convertedValue)
			return nil
		}
	}
	return nil
}

// ToArrayTypeFallback like (from string to []string) or (from *string to *string)
func ToArrayTypeFallback(find ConversionFinderFunc, fType, tType reflect.Type) FromToFunc {
	// from string to []string
	tKind := tType.Kind()
	if tKind == reflect.Slice || tKind == reflect.Array {
		arrayElementType := tType.Elem()
		if arrayElementType == fType { // from type == array element type
			return func(fValue reflect.Value, ptrToValue reflect.Value) error {
				array := reflect.New(tType)
				array = reflect.Append(array.Elem(), fValue)
				ptrToValue.Elem().Set(array)
				return nil
			}
		}
		if fType.Kind() == reflect.Pointer {
			if fallbackFunc := ToArrayTypeFallback(find, fType.Elem(), tType); fallbackFunc != nil {
				return func(fromVal, ptrToVal reflect.Value) error {
					return fallbackFunc(fromVal.Elem(), ptrToVal)
				}
			}
		}
		if conversionFunc := find(fType, arrayElementType); conversionFunc != nil { // conv(*from,element type)
			return func(fValue reflect.Value, ptrToValue reflect.Value) error {
				convertedPtr := reflect.New(arrayElementType)
				if err := conversionFunc(fValue, convertedPtr); err != nil {
					return err
				}
				array := reflect.New(tType)
				array = reflect.Append(array.Elem(), convertedPtr.Elem())
				ptrToValue.Elem().Set(array)
				return nil
			}
		}
	}
	return nil
}

// FromTypeEqualsToTypeFallback from and to are similar
func FromTypeEqualsToTypeFallback(_ ConversionFinderFunc, fType, tType reflect.Type) FromToFunc {
	// from string to string or *string to *string
	if fType == tType {
		return func(fValue reflect.Value, ptrToValue reflect.Value) error {
			ptrToValue.Elem().Set(fValue)
			return nil
		}
	}
	if fType.Kind() == reflect.Pointer {
		// *string to string
		if fType.Elem() == tType {
			return func(fValue reflect.Value, ptrToValue reflect.Value) error {
				ptrToValue.Elem().Set(fValue.Elem())
				return nil
			}
		}
	}
	if tType.Kind() == reflect.Pointer {
		// from string to *string
		if fType == tType.Elem() {
			return func(fValue reflect.Value, ptrToValue reflect.Value) error {
				if fValue.CanAddr() {
					ptrToValue.Elem().Set(fValue.Addr())
				} else {
					newValue := reflect.New(fValue.Type())
					newValue.Elem().Set(fValue)
					ptrToValue.Elem().Set(newValue)
				}
				return nil
			}
		}
	}
	return nil
}

// FromAnyFallback tries to find a conversion from type any to the desired given to type.
func FromAnyFallback(find ConversionFinderFunc, fType, tType reflect.Type) FromToFunc {
	anyType := reflect.TypeFor[any]()
	if fType == anyType {
		return nil // or we will have a infinite loop
	}

	if converter := find(anyType, tType); converter != nil {
		return func(fValue reflect.Value, ptrTValue reflect.Value) error {
			return converter(fValue, ptrTValue)
		}
	}

	return nil
}

// PointerToPointerToTReflection if wants func(F, **T) tries func(F, *T)
func PointerToPointerToTReflection(getConverter ConversionFinderFunc, fType, tType reflect.Type) FromToFunc {
	if tType.Kind() == reflect.Pointer {
		// in this case:
		// ptrTValue.Type() is **T
		// ptrTValue.Type().Elem() is *T
		// ptrTValue.Type().Elem().Elem() is T
		// we need to find a converter from F to T
		ptrToTType := tType.Elem()
		if fromToFunc := getConverter(fType, ptrToTType); fromToFunc != nil {
			zero := reflect.Zero(tType)
			return func(fValue, ptrTValue reflect.Value) error {
				if isNil(fValue) {
					ptrTValue.Elem().Set(zero) // *T = nil
					return nil
				}
				valueOfPt := reflect.New(ptrToTType) // new(T) that returns *T
				if err := fromToFunc(fValue, valueOfPt); err != nil {
					return err
				}
				// valueOfpT => *T
				// ptrTValue => **T, ptrTValue.Elem() => *T
				ptrTValue.Elem().Set(valueOfPt)
				return nil
			}
		}
	}
	return nil
}

func PointerFallback(getConverter ConversionFinderFunc, fType, tType reflect.Type) FromToFunc {
	if fType.Kind() == reflect.Pointer {
		if fromToFunc := getConverter(fType.Elem(), tType); fromToFunc != nil {
			return func(fValue, ptrTValue reflect.Value) error {
				if fValue.IsNil() {
					ptrTValue.Elem().Set(reflect.Zero(tType))
					return nil
				}
				return fromToFunc(fValue.Elem(), ptrTValue)
			}
		}
	}
	return nil
}

func isNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}
