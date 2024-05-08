package tests

import (
	converter "github.com/smart-libs/go-cc-converter/lib/pkg"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func TestMakeToFuncWithArg(t *testing.T) {
	toFunc := func(s string) (io.Reader, error) {
		return strings.NewReader(s), nil
	}

	fromToFunc := converter.MakeFromToFuncWithReturn(toFunc)
	var dest io.Reader
	source := "test"
	assert.NoError(t, fromToFunc(source, &dest))
	b, err := io.ReadAll(dest)
	assert.NoError(t, err)
	assert.Equal(t, source, string(b))
}
