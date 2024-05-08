package tests

import (
	converterdefault "github.com/smart-libs/go-cc-converter/lib/pkg/default"
	"testing"
	"time"
)

func TestFromStringToTimePtrPtr(t *testing.T) {
	d20210101 := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	testSuite[string, time.Time](t, converterdefault.Converters, "20210101", d20210101)
}
