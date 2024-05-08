package converterfuncs

import (
	"fmt"
	convertertypes "github.com/smart-libs/go-cc-converter/lib/pkg/types"
	"reflect"
	"strconv"
)

type (
	integerTypeSet = interface {
		~int8 | ~int16 | ~int32 | ~int64
	}
	uIntegerTypeSet = interface {
		~uint8 | ~uint16 | ~uint32 | ~uint64
	}
	parserResultTypeSet = interface{ ~int64 | ~uint64 }
)

func newParserUsing[Base convertertypes.IntBase, From ~string, T integerTypeSet | uIntegerTypeSet, P parserResultTypeSet](parse func(string, int, int) (P, error)) func(from From) (T, error) {
	bitSize := reflect.TypeFor[T]().Bits()
	base := convertertypes.IntBaseCode[Base]()
	return func(from From) (T, error) {
		i, err := parse(string(from), base, bitSize)
		if err != nil {
			err = fmt.Errorf("From(%T).To(%T): %w", from, T(i), err)
		}
		return T(i), err
	}
}

func newParserFromAnyUsing[Base convertertypes.IntBase, T integerTypeSet | uIntegerTypeSet, P parserResultTypeSet](parse func(string, int, int) (P, error)) func(from any) (T, error) {
	bitSize := reflect.TypeFor[T]().Bits()
	base := convertertypes.IntBaseCode[Base]()
	return func(from any) (T, error) {
		i, err := parse(fmt.Sprintf("%v", from), base, bitSize)
		if err != nil {
			err = fmt.Errorf("From(%T).To(%T): %w", from, T(i), err)
		}
		return T(i), err
	}
}

func NewConverterFromStringToInt[Base convertertypes.IntBase, From ~string, IntType integerTypeSet]() func(from From) (IntType, error) {
	return newParserUsing[Base, From, IntType, int64](strconv.ParseInt)
}

func NewConverterFromStringToUint[Base convertertypes.IntBase, From ~string, UintType uIntegerTypeSet]() func(from From) (UintType, error) {
	return newParserUsing[Base, From, UintType, uint64](strconv.ParseUint)
}

func NewConverterFromAnyToInt[Base convertertypes.IntBase, IntType integerTypeSet]() func(from any) (IntType, error) {
	return newParserFromAnyUsing[Base, IntType, int64](strconv.ParseInt)
}

func NewConverterFromAnyToUint[Base convertertypes.IntBase, UintType uIntegerTypeSet]() func(from any) (UintType, error) {
	return newParserFromAnyUsing[Base, UintType, uint64](strconv.ParseUint)
}
