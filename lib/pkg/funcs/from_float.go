package converterfuncs

import (
	"fmt"
)

func NewConverterFroFloatToString[T ~float32 | ~float64]() func(T) string {
	return func(t T) string {
		return fmt.Sprintf("%f", t)
	}
}
