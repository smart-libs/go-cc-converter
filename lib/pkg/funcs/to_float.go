package converterfuncs

import (
	"fmt"
	"reflect"
	"strconv"
)

func NewFloatParser[T ~float32 | ~float64]() func(from string) (T, error) {
	bitSize := reflect.TypeFor[T]().Bits()
	return func(from string) (T, error) {
		i, err := strconv.ParseFloat(from, bitSize)
		if err != nil {
			err = fmt.Errorf("From(%T).To(%T): %w", from, T(i), err)
		}
		return T(i), err
	}
}
