package fallback

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"testing"
)

func TestFromTypeEqualsToTypeFallback(t *testing.T) {
	var (
		fromString1 = "from1"
		fromString2 = "from2"
		toString    string
		ptrToString *string
	)
	type args struct {
		fValue  reflect.Value
		pTValue reflect.Value
	}
	tests := []struct {
		name string
		args args
		want func(t *testing.T, input args, fallbackFunc FromToFunc)
	}{
		{
			name: "From string to string",
			args: args{
				fValue:  reflect.ValueOf(fromString1),
				pTValue: reflect.ValueOf(&toString),
			},
			want: func(t *testing.T, input args, fallbackFunc FromToFunc) {
				if assert.NotNil(t, fallbackFunc) {
					if assert.NoError(t, fallbackFunc(input.fValue, input.pTValue)) {
						assert.Equal(t, fromString1, toString)
					}
					// there was a bug that when invoking by second time the value set was from the first call
					input.fValue = reflect.ValueOf(fromString2)
					if assert.NoError(t, fallbackFunc(input.fValue, input.pTValue)) {
						assert.Equal(t, fromString2, toString)
					}
				}
			},
		},
		{
			name: "From *string to string",
			args: args{
				fValue:  reflect.ValueOf(&fromString1),
				pTValue: reflect.ValueOf(&toString),
			},
			want: func(t *testing.T, input args, fallbackFunc FromToFunc) {
				if assert.NotNil(t, fallbackFunc) {
					if assert.NoError(t, fallbackFunc(input.fValue, input.pTValue)) {
						assert.Equal(t, fromString1, toString)
					}
					// there was a bug that when invoking by second time the value set was from the first call
					input.fValue = reflect.ValueOf(&fromString2)
					if assert.NoError(t, fallbackFunc(input.fValue, input.pTValue)) {
						assert.Equal(t, fromString2, toString)
					}
				}
			},
		},
		{
			name: "From *string to *string",
			args: args{
				fValue:  reflect.ValueOf(&fromString1),
				pTValue: reflect.ValueOf(&ptrToString),
			},
			want: func(t *testing.T, input args, fallbackFunc FromToFunc) {
				if assert.NotNil(t, fallbackFunc) {
					if assert.NoError(t, fallbackFunc(input.fValue, input.pTValue)) {
						assert.Equal(t, fromString1, *ptrToString)
					}
					// there was a bug that when invoking by second time the value set was from the first call
					input.fValue = reflect.ValueOf(&fromString2)
					if assert.NoError(t, fallbackFunc(input.fValue, input.pTValue)) {
						assert.Equal(t, fromString2, *ptrToString)
					}
				}
			},
		},
		{
			name: "From string to *string",
			args: args{
				fValue:  reflect.ValueOf(fromString1),
				pTValue: reflect.ValueOf(&ptrToString),
			},
			want: func(t *testing.T, input args, fallbackFunc FromToFunc) {
				if assert.NotNil(t, fallbackFunc) {
					if assert.NoError(t, fallbackFunc(input.fValue, input.pTValue)) {
						assert.Equal(t, fromString1, *ptrToString)
					}
					// there was a bug that when invoking by second time the value set was from the first call
					input.fValue = reflect.ValueOf(fromString2)
					if assert.NoError(t, fallbackFunc(input.fValue, input.pTValue)) {
						assert.Equal(t, fromString2, *ptrToString)
					}
				}
			},
		},
	}
	for _, tt := range tests {
		toString = ""
		t.Run(tt.name, func(t *testing.T) {
			fallbackFunc := FromTypeEqualsToTypeFallback(nil, tt.args.fValue.Type(), tt.args.pTValue.Type().Elem())
			tt.want(t, tt.args, fallbackFunc)
		})
	}
}

func TestConvertibleToFallback(t *testing.T) {
	type AnyStringType string
	var (
		toString1        string
		toString2        string
		toAnyStringType1 AnyStringType
		toAnyStringType2 AnyStringType
	)
	type args struct {
		fValue  reflect.Value
		pTValue reflect.Value
	}
	tests := []struct {
		name string
		args args
		want func(t *testing.T, input args, fallbackFunc FromToFunc)
	}{
		{
			name: "From any string type to string",
			args: args{
				fValue:  reflect.ValueOf("from1"),
				pTValue: reflect.ValueOf(&toAnyStringType1),
			},
			want: func(t *testing.T, input args, fallbackFunc FromToFunc) {
				if assert.NotNil(t, fallbackFunc) {
					if assert.NoError(t, fallbackFunc(input.fValue, input.pTValue)) {
						assert.Equal(t, "from1", string(toAnyStringType1))
					}
					// there was a bug that when invoking by second time the value set was from the first call
					input.fValue = reflect.ValueOf("from2")
					input.pTValue = reflect.ValueOf(&toAnyStringType2)
					if assert.NoError(t, fallbackFunc(input.fValue, input.pTValue)) {
						assert.Equal(t, "from2", string(toAnyStringType2))
					}
				}
			},
		},
		{
			name: "From string to any string type",
			args: args{
				fValue:  reflect.ValueOf(AnyStringType("from1")),
				pTValue: reflect.ValueOf(&toString1),
			},
			want: func(t *testing.T, input args, fallbackFunc FromToFunc) {
				if assert.NotNil(t, fallbackFunc) {
					if assert.NoError(t, fallbackFunc(input.fValue, input.pTValue)) {
						assert.Equal(t, "from1", toString1)
					}
					// there was a bug that when invoking by second time the value set was from the first call
					input.fValue = reflect.ValueOf(AnyStringType("from2"))
					input.pTValue = reflect.ValueOf(&toString2)
					if assert.NoError(t, fallbackFunc(input.fValue, input.pTValue)) {
						assert.Equal(t, "from2", string(toAnyStringType2))
					}
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fallbackFunc := ConvertibleToFallback(nil, tt.args.fValue.Type(), tt.args.pTValue.Type().Elem())
			tt.want(t, tt.args, fallbackFunc)
		})
	}
}

func TestToArrayTypeFallback(t *testing.T) {
	var stringArray []string
	fromString1 := "1234"
	fromString2 := "xyz"
	int1 := 12345
	int2 := 9875
	type args struct {
		find      ConversionFinderFunc
		fValue    reflect.Value
		ptrTValue reflect.Value
	}
	tests := []struct {
		name string
		args args
		want func(t *testing.T, input args, fallbackFunc FromToFunc)
	}{
		{
			name: "We can convert from string to []string",
			args: args{
				fValue:    reflect.ValueOf("abc"),
				ptrTValue: reflect.ValueOf(&stringArray),
			},
			want: func(t *testing.T, input args, fallbackFunc FromToFunc) {
				if assert.NotNil(t, fallbackFunc) {
					assert.NoError(t, fallbackFunc(input.fValue, input.ptrTValue))
					assert.Equal(t, []string{"abc"}, input.ptrTValue.Elem().Interface())

					input.fValue = reflect.ValueOf(fromString2)
					assert.NoError(t, fallbackFunc(input.fValue, input.ptrTValue))
					assert.Equal(t, []string{fromString2}, input.ptrTValue.Elem().Interface())
				}
			},
		},
		{
			name: "We can convert from int to []string",
			args: args{
				find: func(from, to reflect.Type) FromToFunc {
					return func(fromVal, ptrToVal reflect.Value) error {
						str := strconv.Itoa(fromVal.Interface().(int))
						*(ptrToVal.Interface().(*string)) = str
						return nil
					}
				},
				fValue:    reflect.ValueOf(int(1234)),
				ptrTValue: reflect.ValueOf(&stringArray),
			},
			want: func(t *testing.T, input args, fallbackFunc FromToFunc) {
				if assert.NotNil(t, fallbackFunc) {
					assert.NoError(t, fallbackFunc(input.fValue, input.ptrTValue))
					assert.Equal(t, []string{"1234"}, input.ptrTValue.Elem().Interface())

					input.fValue = reflect.ValueOf(9876)
					assert.NoError(t, fallbackFunc(input.fValue, input.ptrTValue))
					assert.Equal(t, []string{"9876"}, input.ptrTValue.Elem().Interface())
				}
			},
		},
		{
			name: "We can convert from *string to []string",
			args: args{
				fValue:    reflect.ValueOf(&fromString1),
				ptrTValue: reflect.ValueOf(&stringArray),
			},
			want: func(t *testing.T, input args, fallbackFunc FromToFunc) {
				if assert.NotNil(t, fallbackFunc) {
					assert.NoError(t, fallbackFunc(input.fValue, input.ptrTValue))
					assert.Equal(t, []string{fromString1}, input.ptrTValue.Elem().Interface())

					input.fValue = reflect.ValueOf(&fromString2)
					assert.NoError(t, fallbackFunc(input.fValue, input.ptrTValue))
					assert.Equal(t, []string{fromString2}, input.ptrTValue.Elem().Interface())
				}
			},
		},
		{
			name: "We can convert from *int to []string",
			args: args{
				find: func(from, to reflect.Type) FromToFunc {
					return func(fromVal, ptrToVal reflect.Value) error {
						str := strconv.Itoa(fromVal.Interface().(int))
						*(ptrToVal.Interface().(*string)) = str
						return nil
					}
				},
				fValue:    reflect.ValueOf(&int1),
				ptrTValue: reflect.ValueOf(&stringArray),
			},
			want: func(t *testing.T, input args, fallbackFunc FromToFunc) {
				if assert.NotNil(t, fallbackFunc) {
					assert.NoError(t, fallbackFunc(input.fValue, input.ptrTValue))
					assert.Equal(t, []string{strconv.Itoa(int1)}, input.ptrTValue.Elem().Interface())

					input.fValue = reflect.ValueOf(&int2)
					assert.NoError(t, fallbackFunc(input.fValue, input.ptrTValue))
					assert.Equal(t, []string{strconv.Itoa(int2)}, input.ptrTValue.Elem().Interface())
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fromToFunc := ToArrayTypeFallback(tt.args.find, tt.args.fValue.Type(), tt.args.ptrTValue.Type().Elem())
			tt.want(t, tt.args, fromToFunc)
		})
	}
}

func MakeFromToFuncForTest[From any, To any](conversion func(From, *To) error) FromToFunc {
	return func(fromVal, ptrToVal reflect.Value) error {
		if fromVal.Type() == reflect.TypeFor[From]() || reflect.TypeFor[any]() == reflect.TypeFor[From]() {
			if ptrToVal.Type().Elem() == reflect.TypeFor[To]() {
				return conversion(fromVal.Interface().(From), ptrToVal.Interface().(*To))
			}
		}
		return fmt.Errorf("invalid input from=[%T] to=[%T]", fromVal.Interface(), ptrToVal.Interface())
	}
}

func TestFromAnyFallback(t *testing.T) {
	type args struct {
		find  ConversionFinderFunc
		fType reflect.Type
		tType reflect.Type
	}
	tests := []struct {
		name string
		args args
		want func(t *testing.T, input args, fallbackFunc FromToFunc)
	}{
		{
			name: "If from value is any, then no function must be returned",
			args: args{
				find: func(from, to reflect.Type) FromToFunc {
					assert.Fail(t, "should not be invoked")
					return nil
				},
				fType: reflect.TypeFor[any](),
				tType: reflect.TypeFor[string](),
			},
			want: func(t *testing.T, input args, fallbackFunc FromToFunc) {
				assert.Nil(t, fallbackFunc)
			},
		},
		{
			name: "If `from` is int and `to` string, then find is used and conversion should be performed",
			args: args{
				find: func(from, to reflect.Type) FromToFunc {
					return MakeFromToFuncForTest(func(v any, to *string) error {
						*to = fmt.Sprintf("%v", v)
						return nil
					})
				},
				fType: reflect.TypeFor[int](),
				tType: reflect.TypeFor[string](),
			},
			want: func(t *testing.T, input args, fallbackFunc FromToFunc) {
				if assert.NotNil(t, fallbackFunc) {
					var str string
					assert.NoError(t, fallbackFunc(reflect.ValueOf(123), reflect.ValueOf(&str)))
					assert.Equal(t, "123", str)

					assert.NoError(t, fallbackFunc(reflect.ValueOf(123.56), reflect.ValueOf(&str)))
					assert.Equal(t, "123.56", str)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fromToFunc := FromAnyFallback(tt.args.find, tt.args.fType, tt.args.tType)
			tt.want(t, tt.args, fromToFunc)
		})
	}
}

func TestPointerFallback(t *testing.T) {
	type args struct {
		find  ConversionFinderFunc
		fType reflect.Type
		tType reflect.Type
	}
	tests := []struct {
		name string
		args args
		want func(t *testing.T, input args, fallbackFunc FromToFunc)
	}{
		{
			name: "If `from` is int and `to` string, then find is used and conversion should be performed",
			args: args{
				find: func(from, to reflect.Type) FromToFunc {
					return MakeFromToFuncForTest(func(v any, to *string) error {
						*to = fmt.Sprintf("%v", v)
						return nil
					})
				},
				fType: reflect.TypeFor[*int](),
				tType: reflect.TypeFor[string](),
			},
			want: func(t *testing.T, input args, fallbackFunc FromToFunc) {
				if assert.NotNil(t, fallbackFunc) {
					var str string
					var intVal int = 123
					assert.NoError(t, fallbackFunc(reflect.ValueOf(&intVal), reflect.ValueOf(&str)))
					assert.Equal(t, "123", str)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fromToFunc := PointerFallback(tt.args.find, tt.args.fType, tt.args.tType)
			tt.want(t, tt.args, fromToFunc)
		})
	}
}

func TestPointerToPointerToTReflection(t *testing.T) {
	type args struct {
		find  ConversionFinderFunc
		fType reflect.Type
		tType reflect.Type
	}
	tests := []struct {
		name string
		args args
		want func(t *testing.T, input args, fallbackFunc FromToFunc)
	}{
		{
			name: "If `from` is int and `to` *string, then find is used and conversion should be performed",
			args: args{
				find: func(from, to reflect.Type) FromToFunc {
					return MakeFromToFuncForTest(func(v any, to *string) error {
						*to = fmt.Sprintf("%v", v)
						return nil
					})
				},
				fType: reflect.TypeFor[int](),
				tType: reflect.TypeFor[*string](),
			},
			want: func(t *testing.T, input args, fallbackFunc FromToFunc) {
				if assert.NotNil(t, fallbackFunc) {
					var str *string
					assert.NoError(t, fallbackFunc(reflect.ValueOf(123), reflect.ValueOf(&str)))
					if assert.NotNil(t, str) {
						assert.Equal(t, "123", *str)
					}
				}
			},
		},
		{
			name: "If `from` is *int and `to` *string, then find is used and conversion should be performed",
			args: args{
				find: func(from, to reflect.Type) FromToFunc {
					return MakeFromToFuncForTest(func(v any, to *string) error {
						*to = fmt.Sprintf("%v", v)
						return nil
					})
				},
				fType: reflect.TypeFor[*int](),
				tType: reflect.TypeFor[*string](),
			},
			want: func(t *testing.T, input args, fallbackFunc FromToFunc) {
				if assert.NotNil(t, fallbackFunc) {
					var str *string
					var i *int

					assert.NoError(t, fallbackFunc(reflect.ValueOf(i), reflect.ValueOf(&str)))
					assert.Equal(t, (*string)(nil), str)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fromToFunc := PointerToPointerToTReflection(tt.args.find, tt.args.fType, tt.args.tType)
			tt.want(t, tt.args, fromToFunc)
		})
	}
}
