package samples

import (
	converterdefault "github.com/smart-libs/go-cc-converter/lib/pkg/default"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromArrayOfStringsToArrayOfConvertible(t *testing.T) {
	type ID string

	source := []string{"1", "2"}
	destination := []ID{}

	assert.NoError(t, converterdefault.Converters.Convert(source, &destination))
	assert.Equal(t, source[0], string(destination[0]))
	assert.Equal(t, source[1], string(destination[1]))
}
