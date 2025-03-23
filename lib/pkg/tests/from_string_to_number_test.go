package tests

import (
	"fmt"
	converterdefault "github.com/smart-libs/go-cc-converter/lib/pkg/default"
	convertertypes "github.com/smart-libs/go-cc-converter/lib/pkg/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromStringToInt8(t *testing.T) {
	t.Run("StringInt[Binary] to int8", func(t *testing.T) {
		type BinInt8 = convertertypes.StringInt[convertertypes.Binary]
		testSuite[BinInt8, int8](t, converterdefault.Converters, BinInt8("101"), 5)
		testSuite[BinInt8, int8](t, converterdefault.Converters, "-101", -5)
	})
	t.Run("String to Implied Int8", func(t *testing.T) {
		type ImpliedInt8 = convertertypes.Int8[convertertypes.ImpliedBase]
		testSuite[string, ImpliedInt8](t, converterdefault.Converters, "0b101", ImpliedInt8(5))
		testSuite[string, ImpliedInt8](t, converterdefault.Converters, "0o101", ImpliedInt8(65))
		testSuite[string, ImpliedInt8](t, converterdefault.Converters, "101", ImpliedInt8(101))
		testSuite[string, ImpliedInt8](t, converterdefault.Converters, "0x71", ImpliedInt8(113))
		testSuite[string, ImpliedInt8](t, converterdefault.Converters, "-0b101", ImpliedInt8(-5))
	})
	t.Run("String to BinInt8", func(t *testing.T) {
		type BinInt8 = convertertypes.Int8[convertertypes.Binary]
		testSuite[string, BinInt8](t, converterdefault.Converters, "101", BinInt8(5))
		testSuite[string, BinInt8](t, converterdefault.Converters, "-101", BinInt8(-5))
	})
	t.Run("String to OctalInt8", func(t *testing.T) {
		type OctalInt8 = convertertypes.Int8[convertertypes.Octal]
		testSuite[string, OctalInt8](t, converterdefault.Converters, "101", OctalInt8(65))
	})
	t.Run("String to DecimalInt8", func(t *testing.T) {
		type DecimalInt8 = convertertypes.Int8[convertertypes.Decimal]
		testSuite[string, DecimalInt8](t, converterdefault.Converters, "101", DecimalInt8(101))
	})
	t.Run("String to HexadecimalInt8", func(t *testing.T) {
		type HexadecimalInt8 = convertertypes.Int8[convertertypes.Hexadecimal]
		testSuite[string, HexadecimalInt8](t, converterdefault.Converters, "71", HexadecimalInt8((7*16)+1))
	})
}

func TestFromStringToUint8(t *testing.T) {
	t.Run("String to BinUint8", func(t *testing.T) {
		type BinUint8 = convertertypes.Uint8[convertertypes.Binary]
		testSuite[string, BinUint8](t, converterdefault.Converters, "10000101", BinUint8(128+5))
	})
	t.Run("String to OctalUint8", func(t *testing.T) {
		type OctalUint8 = convertertypes.Uint8[convertertypes.Octal]
		testSuite[string, OctalUint8](t, converterdefault.Converters, "201", OctalUint8(129))
	})
	t.Run("String to DecimalUint8", func(t *testing.T) {
		type DecimalUint8 = convertertypes.Uint8[convertertypes.Decimal]
		testSuite[string, DecimalUint8](t, converterdefault.Converters, "201", DecimalUint8(201))
	})
	t.Run("String to HexadecimalUint8", func(t *testing.T) {
		type HexadecimalUint8 = convertertypes.Uint8[convertertypes.Hexadecimal]
		testSuite[string, HexadecimalUint8](t, converterdefault.Converters, "81", HexadecimalUint8((8*16)+1))
	})
}

func TestFromStringToInt8Binary(t *testing.T) {
	type BinInt8 = convertertypes.Int8[convertertypes.Binary]
	var to *BinInt8
	testFromTo(t, []testCase[string, *BinInt8]{
		{
			name: "string(\"121\") to *BinInt8(5)",
			args: argsFromTo[string, *BinInt8]{from: "121", to: &to},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				expected := "From(string).To(convertertypes.Int8[github.com/smart-libs/go-cc-converter/lib/pkg/types.Binary]): strconv.ParseInt: parsing \"121\": invalid syntax"
				fmt.Printf("to=[%v], type=[%T]\n", to, to)
				fmt.Printf("err=[%v], type=[%T]\n", err, err)
				return assert.Equal(t, expected, err.Error())
			},
			wantTo: nil,
		},
	})
}

func TestFromStringToUint8Binary(t *testing.T) {
	type BinUint8 = convertertypes.Uint8[convertertypes.Binary]
	var to *BinUint8
	testFromTo(t, []testCase[string, *BinUint8]{
		{
			name: "string(\"121\") to *BinUint8(5)",
			args: argsFromTo[string, *BinUint8]{from: "121", to: &to},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				expected := "From(string).To(convertertypes.Uint8[github.com/smart-libs/go-cc-converter/lib/pkg/types.Binary]): strconv.ParseUint: parsing \"121\": invalid syntax"
				return assert.Equal(t, expected, err.Error())
			},
			wantTo: nil,
		},
	})
}

func TestFromStringToInt8Octal(t *testing.T) {
	type OctInt8 = convertertypes.Int8[convertertypes.Octal]
	d121 := OctInt8(81)
	var to *OctInt8
	testFromTo(t, []testCase[string, *OctInt8]{
		{
			name:    "string(\"121\") to *OctInt8(121)",
			args:    argsFromTo[string, *OctInt8]{from: "121", to: &to},
			wantErr: assert.NoError,
			wantTo:  &d121,
		},
	})
}

func TestFromStringToInt8Decimal(t *testing.T) {
	type DecInt8 = convertertypes.Int8[convertertypes.Decimal]
	d121 := DecInt8(121)
	var to *DecInt8
	testFromTo(t, []testCase[string, *DecInt8]{
		{
			name:    "string(\"121\") to *DecInt8(121)",
			args:    argsFromTo[string, *DecInt8]{from: "121", to: &to},
			wantErr: assert.NoError,
			wantTo:  &d121,
		},
	})
}

func TestFromStringToInt16Hexadecimal(t *testing.T) {
	type HexaInt8 = convertertypes.Int8[convertertypes.Hexadecimal]
	d24 := HexaInt8(36)
	var to *HexaInt8
	testFromTo(t, []testCase[string, *HexaInt8]{
		{
			name:    "string(\"24\") to *HexaInt8(36)",
			args:    argsFromTo[string, *HexaInt8]{from: "24", to: &to},
			wantErr: assert.NoError,
			wantTo:  &d24,
		},
	})
}
