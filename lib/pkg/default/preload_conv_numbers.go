package converterdefault

import (
	funcs "github.com/smart-libs/go-cc-converter/lib/pkg/funcs"
	convertertypes "github.com/smart-libs/go-cc-converter/lib/pkg/types"
	"strconv"
)

func init() {
	type (
		Binary      = convertertypes.Binary
		Octal       = convertertypes.Octal
		Decimal     = convertertypes.Decimal
		Hexadecimal = convertertypes.Hexadecimal
		ImpliedBase = convertertypes.ImpliedBase
		Base        = convertertypes.IntBase
	)

	AddHandlerWithReturnNoError(strconv.Itoa)
	AddHandlerWithReturnNoError(funcs.NewConverterFroFloatToString[float32]())
	AddHandlerWithReturnNoError(funcs.NewConverterFroFloatToString[float64]())

	AddHandlerWithReturnNoError(funcs.NewConverterFromIntToString[int8]())
	AddHandlerWithReturnNoError(funcs.NewConverterFromIntToString[int16]())
	AddHandlerWithReturnNoError(funcs.NewConverterFromIntToString[int32]())
	AddHandlerWithReturnNoError(funcs.NewConverterFromIntToString[int64]())

	AddHandlerWithReturnNoError(funcs.NewConverterFromIntBaseToString[Binary, convertertypes.Int8[Binary]]())
	AddHandlerWithReturnNoError(funcs.NewConverterFromIntBaseToString[Binary, convertertypes.Int16[Binary]]())
	AddHandlerWithReturnNoError(funcs.NewConverterFromIntBaseToString[Binary, convertertypes.Int32[Binary]]())
	AddHandlerWithReturnNoError(funcs.NewConverterFromIntBaseToString[Binary, convertertypes.Int64[Binary]]())

	AddHandlerWithReturnNoError(funcs.NewConverterFromIntBaseToString[Octal, convertertypes.Int8[Octal]]())
	AddHandlerWithReturnNoError(funcs.NewConverterFromIntBaseToString[Octal, convertertypes.Int16[Octal]]())
	AddHandlerWithReturnNoError(funcs.NewConverterFromIntBaseToString[Octal, convertertypes.Int32[Octal]]())
	AddHandlerWithReturnNoError(funcs.NewConverterFromIntBaseToString[Octal, convertertypes.Int64[Octal]]())

	AddHandlerWithReturnNoError(funcs.NewConverterFromIntBaseToString[Decimal, convertertypes.Int8[Decimal]]())
	AddHandlerWithReturnNoError(funcs.NewConverterFromIntBaseToString[Decimal, convertertypes.Int16[Decimal]]())
	AddHandlerWithReturnNoError(funcs.NewConverterFromIntBaseToString[Decimal, convertertypes.Int32[Decimal]]())
	AddHandlerWithReturnNoError(funcs.NewConverterFromIntBaseToString[Decimal, convertertypes.Int64[Decimal]]())

	AddHandlerWithReturnNoError(funcs.NewConverterFromIntBaseToString[Hexadecimal, convertertypes.Int8[Hexadecimal]]())
	AddHandlerWithReturnNoError(funcs.NewConverterFromIntBaseToString[Hexadecimal, convertertypes.Int16[Hexadecimal]]())
	AddHandlerWithReturnNoError(funcs.NewConverterFromIntBaseToString[Hexadecimal, convertertypes.Int32[Hexadecimal]]())
	AddHandlerWithReturnNoError(funcs.NewConverterFromIntBaseToString[Hexadecimal, convertertypes.Int64[Hexadecimal]]())

	AddHandlerWithReturn(strconv.Atoi)

	AddHandlerWithReturn(funcs.NewFloatParser[float32]())
	AddHandlerWithReturn(funcs.NewFloatParser[float64]())

	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[ImpliedBase, convertertypes.StringInt[ImpliedBase], int8]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[ImpliedBase, convertertypes.StringInt[ImpliedBase], int16]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[ImpliedBase, convertertypes.StringInt[ImpliedBase], int32]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[ImpliedBase, convertertypes.StringInt[ImpliedBase], int64]())

	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Binary, convertertypes.StringInt[Binary], int8]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Binary, convertertypes.StringInt[Binary], int16]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Binary, convertertypes.StringInt[Binary], int32]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Binary, convertertypes.StringInt[Binary], int64]())

	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Octal, convertertypes.StringInt[Octal], int8]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Octal, convertertypes.StringInt[Octal], int16]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Octal, convertertypes.StringInt[Octal], int32]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Octal, convertertypes.StringInt[Octal], int64]())

	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Decimal, convertertypes.StringInt[Decimal], int8]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Decimal, convertertypes.StringInt[Decimal], int16]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Decimal, convertertypes.StringInt[Decimal], int32]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Decimal, convertertypes.StringInt[Decimal], int64]())

	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Hexadecimal, convertertypes.StringInt[Hexadecimal], int8]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Hexadecimal, convertertypes.StringInt[Hexadecimal], int16]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Hexadecimal, convertertypes.StringInt[Hexadecimal], int32]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Hexadecimal, convertertypes.StringInt[Hexadecimal], int64]())

	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[ImpliedBase, convertertypes.StringInt[ImpliedBase], uint8]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[ImpliedBase, convertertypes.StringInt[ImpliedBase], uint16]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[ImpliedBase, convertertypes.StringInt[ImpliedBase], uint32]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[ImpliedBase, convertertypes.StringInt[ImpliedBase], uint64]())

	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Binary, convertertypes.StringInt[Binary], uint8]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Binary, convertertypes.StringInt[Binary], uint16]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Binary, convertertypes.StringInt[Binary], uint32]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Binary, convertertypes.StringInt[Binary], uint64]())

	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Octal, convertertypes.StringInt[Octal], uint8]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Octal, convertertypes.StringInt[Octal], uint16]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Octal, convertertypes.StringInt[Octal], uint32]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Octal, convertertypes.StringInt[Octal], uint64]())

	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Decimal, convertertypes.StringInt[Decimal], uint8]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Decimal, convertertypes.StringInt[Decimal], uint16]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Decimal, convertertypes.StringInt[Decimal], uint32]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Decimal, convertertypes.StringInt[Decimal], uint64]())

	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Hexadecimal, convertertypes.StringInt[Hexadecimal], uint8]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Hexadecimal, convertertypes.StringInt[Hexadecimal], uint16]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Hexadecimal, convertertypes.StringInt[Hexadecimal], uint32]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Hexadecimal, convertertypes.StringInt[Hexadecimal], uint64]())

	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[ImpliedBase, string, convertertypes.Int8[ImpliedBase]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[ImpliedBase, string, convertertypes.Int16[ImpliedBase]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[ImpliedBase, string, convertertypes.Int32[ImpliedBase]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[ImpliedBase, string, convertertypes.Int64[ImpliedBase]]())

	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[ImpliedBase, string, convertertypes.Uint8[ImpliedBase]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[ImpliedBase, string, convertertypes.Uint16[ImpliedBase]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[ImpliedBase, string, convertertypes.Uint32[ImpliedBase]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[ImpliedBase, string, convertertypes.Uint64[ImpliedBase]]())

	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Binary, string, convertertypes.Int8[Binary]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Binary, string, convertertypes.Int16[Binary]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Binary, string, convertertypes.Int32[Binary]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Binary, string, convertertypes.Int64[Binary]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Octal, string, convertertypes.Int8[Octal]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Octal, string, convertertypes.Int16[Octal]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Octal, string, convertertypes.Int32[Octal]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Octal, string, convertertypes.Int64[Octal]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Decimal, string, convertertypes.Int8[Decimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Decimal, string, convertertypes.Int16[Decimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Decimal, string, convertertypes.Int32[Decimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Decimal, string, convertertypes.Int64[Decimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Hexadecimal, string, convertertypes.Int8[Hexadecimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Hexadecimal, string, convertertypes.Int16[Hexadecimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Hexadecimal, string, convertertypes.Int32[Hexadecimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToInt[Hexadecimal, string, convertertypes.Int64[Hexadecimal]]())

	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Binary, string, convertertypes.Uint8[Binary]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Binary, string, convertertypes.Uint16[Binary]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Binary, string, convertertypes.Uint32[Binary]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Binary, string, convertertypes.Uint64[Binary]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Octal, string, convertertypes.Uint8[Octal]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Octal, string, convertertypes.Uint16[Octal]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Octal, string, convertertypes.Uint32[Octal]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Octal, string, convertertypes.Uint64[Octal]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Decimal, string, convertertypes.Uint8[Decimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Decimal, string, convertertypes.Uint16[Decimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Decimal, string, convertertypes.Uint32[Decimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Decimal, string, convertertypes.Uint64[Decimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Hexadecimal, string, convertertypes.Uint8[Hexadecimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Hexadecimal, string, convertertypes.Uint16[Hexadecimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Hexadecimal, string, convertertypes.Uint32[Hexadecimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromStringToUint[Hexadecimal, string, convertertypes.Uint64[Hexadecimal]]())

	AddHandlerWithReturn(funcs.NewConverterFromAnyToInt[Binary, convertertypes.Int8[Binary]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToInt[Binary, convertertypes.Int16[Binary]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToInt[Binary, convertertypes.Int32[Binary]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToInt[Binary, convertertypes.Int64[Binary]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToInt[Octal, convertertypes.Int8[Octal]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToInt[Octal, convertertypes.Int16[Octal]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToInt[Octal, convertertypes.Int32[Octal]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToInt[Octal, convertertypes.Int64[Octal]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToInt[Decimal, convertertypes.Int8[Decimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToInt[Decimal, convertertypes.Int16[Decimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToInt[Decimal, convertertypes.Int32[Decimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToInt[Decimal, convertertypes.Int64[Decimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToInt[Hexadecimal, convertertypes.Int8[Hexadecimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToInt[Hexadecimal, convertertypes.Int16[Hexadecimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToInt[Hexadecimal, convertertypes.Int32[Hexadecimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToInt[Hexadecimal, convertertypes.Int64[Hexadecimal]]())

	AddHandlerWithReturn(funcs.NewConverterFromAnyToUint[Binary, convertertypes.Uint8[Binary]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToUint[Binary, convertertypes.Uint16[Binary]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToUint[Binary, convertertypes.Uint32[Binary]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToUint[Binary, convertertypes.Uint64[Binary]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToUint[Octal, convertertypes.Uint8[Octal]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToUint[Octal, convertertypes.Uint16[Octal]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToUint[Octal, convertertypes.Uint32[Octal]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToUint[Octal, convertertypes.Uint64[Octal]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToUint[Decimal, convertertypes.Uint8[Decimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToUint[Decimal, convertertypes.Uint16[Decimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToUint[Decimal, convertertypes.Uint32[Decimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToUint[Decimal, convertertypes.Uint64[Decimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToUint[Hexadecimal, convertertypes.Uint8[Hexadecimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToUint[Hexadecimal, convertertypes.Uint16[Hexadecimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToUint[Hexadecimal, convertertypes.Uint32[Hexadecimal]]())
	AddHandlerWithReturn(funcs.NewConverterFromAnyToUint[Hexadecimal, convertertypes.Uint64[Hexadecimal]]())
}
