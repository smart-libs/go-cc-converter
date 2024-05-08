package converterfuncs

import (
	"fmt"
	convertertypes "github.com/smart-libs/go-cc-converter/lib/pkg/types"
)

func NewConverterFromIntToString[T integerTypeSet | uIntegerTypeSet]() func(T) string {
	return NewConverterFromIntBaseToString[convertertypes.ImpliedBase, T]()
}

func NewConverterFromIntBaseToString[Base convertertypes.IntBase, T integerTypeSet | uIntegerTypeSet]() func(T) string {
	mask := convertertypes.IntBaseMask[Base]()
	return func(t T) string {
		return fmt.Sprintf(mask, t)
	}
}
