package tests

import (
	"github.com/smart-libs/go-cc-converter/lib/pkg/funcs"
	convertertypes "github.com/smart-libs/go-cc-converter/lib/pkg/types"
	"testing"
)

func TestFromFlagStringListToStringArray(t *testing.T) {
	type args struct {
		list  convertertypes.FlattenedStringList
		array *[]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "if target array is nil, then returns error",
			wantErr: true,
		},
		{
			name: "if list is empty, then returns an empty array",
			args: args{
				list:  "",
				array: new([]string),
			},
			wantErr: false,
		},
		{
			name: "if list is 'a', then returns an array with a",
			args: args{
				list:  "a",
				array: &[]string{"a"},
			},
			wantErr: false,
		},
		{
			name: "if list is 'a,b', then returns an array with 'a' and 'b'",
			args: args{
				list:  "a,b",
				array: &[]string{"a", "b"},
			},
			wantErr: false,
		},
		{
			name: "if list is 'a, b', then returns an array with 'a' and ' b'",
			args: args{
				list:  "a, b",
				array: &[]string{"a", " b"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := converterfuncs.FromFlattenedStringListToStringArray(string(tt.args.list), tt.args.array); (err != nil) != tt.wantErr {
				t.Errorf("FromFlattenedStringListToStringArray() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
