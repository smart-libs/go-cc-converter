package tests

import (
	converterdefault "github.com/smart-libs/go-cc-converter/lib/pkg/default"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestFromStringPtrToString(t *testing.T) {
	t.Run("String to string", func(t *testing.T) {
		testSuite[string, string](t, converterdefault.Converters, "abcdef", "abcdef")
	})

	givenString := "any"
	type args struct {
		f      any
		toType reflect.Type
	}
	type testCase struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
		wantTo  any
	}

	tests := []testCase{
		{
			name:    "string(\"any\") to string(\"any\")",
			args:    args{f: givenString, toType: reflect.TypeFor[string]()},
			wantErr: assert.NoError,
			wantTo:  givenString,
		},
		{
			name:    "string(\"any\") to *string(\"any\")",
			args:    args{f: givenString, toType: reflect.TypeFor[*string]()},
			wantErr: assert.NoError,
			wantTo:  &givenString,
		},
		{
			name:    "*string(\"any\") to string(\"any\")",
			args:    args{f: &givenString, toType: reflect.TypeFor[string]()},
			wantErr: assert.NoError,
			wantTo:  givenString,
		},
		{
			name:    "*string(\"any\") to *string(\"any\")",
			args:    args{f: &givenString, toType: reflect.TypeFor[*string]()},
			wantErr: assert.NoError,
			wantTo:  &givenString,
		},
		{
			name:    "string(\"any\") to []string{\"any\"}",
			args:    args{f: givenString, toType: reflect.TypeFor[[]string]()},
			wantErr: assert.NoError,
			wantTo:  []string{givenString},
		},
		{
			name:    "*string(\"any\") to []*string{\"any\"}",
			args:    args{f: &givenString, toType: reflect.TypeFor[[]*string]()},
			wantErr: assert.NoError,
			wantTo:  []*string{&givenString},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			targetType := reflect.TypeOf(tt.wantTo)
			convertedValue, err := converterdefault.Converters.ConvertToType(tt.args.f, targetType)
			if tt.wantErr != nil {
				tt.wantErr(t, err)
			}
			assert.Equalf(t, tt.wantTo, convertedValue, "From type=[%T] to type=[%s]", tt.args.f, targetType)
		})
	}
}
