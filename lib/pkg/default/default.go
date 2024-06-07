package converterdefault

import converter "github.com/smart-libs/go-cc-converter/lib/pkg"

var (
	//Converters = NewConverters(Registry)
	Converters = NewBasicConverters(NewFallbackRegistry(Registry))
)

func To[T any](f any) (T, error) { return converter.To[T](Converters, f) }

func NewConverters(registry converter.Registry) converter.Converters {
	return NewBasicConverters(registry)
}
