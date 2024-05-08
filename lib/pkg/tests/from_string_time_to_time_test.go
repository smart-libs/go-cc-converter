package tests

import (
	converterdefault "github.com/smart-libs/go-cc-converter/lib/pkg/default"
	convertertypes "github.com/smart-libs/go-cc-converter/lib/pkg/types"
	"testing"
	"time"
)

func TestFromStringTimeToTimePtrPtr(t *testing.T) {
	t.Run("RFC3339Year", func(t *testing.T) {
		d20210101 := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
		type stringTimeYearType = convertertypes.StringTime[convertertypes.RFC3339Year]
		stringTimeYear := stringTimeYearType("2021")
		testSuite[stringTimeYearType, time.Time](t, converterdefault.Converters, stringTimeYear, d20210101)
	})
	t.Run("RFC3339Year to int", func(t *testing.T) {
		type stringTimeYearType = convertertypes.StringTime[convertertypes.RFC3339Year]
		stringTimeYear := stringTimeYearType("2021")
		testSuite[stringTimeYearType, int](t, converterdefault.Converters, stringTimeYear, 2021)
	})
}
