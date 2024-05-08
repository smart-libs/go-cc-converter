package convertertypes

type (
	// These types are used as base to number conversions to string or from string

	// ImpliedBase is used when the number format in string should be used to identify the base
	ImpliedBase struct{}
	Binary      struct{}
	Octal       struct{}
	Decimal     struct{}
	Hexadecimal struct{}
	// IntBase is a number base constraint
	IntBase interface {
		Binary | Octal | Decimal | Hexadecimal | ImpliedBase
	}

	Int8[Base IntBase]   int8
	Int16[Base IntBase]  int16
	Int32[Base IntBase]  int32
	Int64[Base IntBase]  int64
	UInt[Base IntBase]   uint
	Uint8[Base IntBase]  uint8
	Uint16[Base IntBase] uint16
	Uint32[Base IntBase] uint32
	Uint64[Base IntBase] uint64

	// StringInt is used when you want to identify the base to be used to convert the string to int.
	// Example `converterdefault.Converters.Convert(StringInt[Decimal]("23"), &myInt8)`
	StringInt[Base IntBase] string
)

func IntBaseMask[B IntBase]() string {
	var (
		b    B
		base any = b
	)
	switch base.(type) {
	case Binary:
		return "0b%b"
	case Octal:
		return "%O"
	case Decimal:
		return "%d"
	case Hexadecimal:
		return "%x"
	default:
		return "%d"
	}
}

func IntBaseCode[B IntBase]() int {
	var (
		b    B
		base any = b
	)
	switch base.(type) {
	case Binary:
		return 2
	case Octal:
		return 8
	case Decimal:
		return 10
	case Hexadecimal:
		return 16
	default:
		return 0
	}
}
