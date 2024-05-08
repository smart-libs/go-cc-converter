package tests

import (
	converterdefault "github.com/smart-libs/go-cc-converter/lib/pkg/default"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromIntToFloat32(t *testing.T) {
	givenFloat32 := float32(10)
	type args[F any, T any] struct {
		f  F
		pT *T
	}
	type testCase[F any, T any] struct {
		name    string
		args    args[F, T]
		wantErr assert.ErrorAssertionFunc
		wantTo  *T
	}

	tests := []testCase[int, float32]{
		{
			name:    "from int to float32",
			args:    args[int, float32]{f: 10, pT: new(float32)},
			wantErr: assert.NoError,
			wantTo:  &givenFloat32,
		},
		{
			name: "from int to nil float32",
			args: args[int, float32]{f: 10},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if assert.NotNil(t, err) {
					assert.Equal(t, "ConversionError: from int=[10] to *float32=[<nil>] due to the to argument is not nil pointer", err.Error())
				}
				return true
			},
			wantTo: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := converterdefault.Converters.Convert(tt.args.f, tt.args.pT)
			if tt.wantErr != nil {
				tt.wantErr(t, err)
			}
			assert.Equal(t, tt.wantTo, tt.args.pT)
		})
	}
}
