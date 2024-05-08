package tests

import (
	converterdefault "github.com/smart-libs/go-cc-converter/lib/pkg/default"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromNil(t *testing.T) {
	str := "123"
	var from any
	assert.NoError(t, converterdefault.Converters.Convert(nil, &str))
	assert.Equal(t, "", str)

	assert.NoError(t, converterdefault.Converters.Convert(from, &str))
	assert.Equal(t, "", str)

	ptr := &str
	assert.NoError(t, converterdefault.Converters.Convert(nil, &ptr))
	assert.Equal(t, (*string)(nil), ptr)

	assert.NoError(t, converterdefault.Converters.Convert(from, &ptr))
	assert.Equal(t, (*string)(nil), ptr)
}
