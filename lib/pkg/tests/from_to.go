package tests

import (
	"fmt"
	converter "github.com/smart-libs/go-cc-converter/lib/pkg"
	converterdefault "github.com/smart-libs/go-cc-converter/lib/pkg/default"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

type (
	argsFromTo[From any, To any] struct {
		from From
		to   *To
	}

	testCase[From any, To any] struct {
		name    string
		args    argsFromTo[From, To]
		wantErr assert.ErrorAssertionFunc
		wantTo  To
	}
)

func testFromTo[From any, To any](t *testing.T, testCases []testCase[From, To]) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			//fmt.Printf("DefaultConverters.Convert(%v, %v)\n", tt.args.from, tt.args.to)
			err := converterdefault.Converters.Convert(tt.args.from, tt.args.to)
			if tt.wantErr != nil {
				tt.wantErr(t, err)
			}
			if err == nil {
				if !assert.Equal(t, tt.wantTo, *tt.args.to) {
					valueOfWant := reflect.ValueOf(tt.wantTo)
					if valueOfWant.Kind() == reflect.Pointer {
						valueOfResult := reflect.ValueOf(*tt.args.to)
						assert.Equal(t, fmt.Sprintf("%v", valueOfWant.Elem().Interface()), fmt.Sprintf("%v", valueOfResult.Elem().Interface()))
					} else {
						assert.Equal(t, fmt.Sprintf("%v", tt.wantTo), fmt.Sprintf("%v", *tt.args.to))
					}
				}
			}
		})
	}
}

func testSuite[From any, To any](t *testing.T, converters converter.Converters, from From, expectedTo To) {
	var zero To
	var to To
	t.Run(fmt.Sprintf("func(F, *T) error => func(%v, %T) => %v", from, &to, expectedTo), func(t *testing.T) {
		assert.NoError(t, converters.Convert(from, &to))
		assert.Equal(t, expectedTo, to)
	})

	var ptTo *To
	t.Run(fmt.Sprintf("func(F, **T) error => func(%v, %T) => %v", from, &ptTo, expectedTo), func(t *testing.T) {
		if assert.NoError(t, converters.Convert(from, &ptTo)) {
			assert.Equal(t, expectedTo, *ptTo)
		}
	})

	t.Run(fmt.Sprintf("func(*F, *T) error => func(%v, %T) => %v", &from, &to, expectedTo), func(t *testing.T) {
		to = zero
		assert.NoError(t, converters.Convert(&from, &to))
		assert.Equal(t, expectedTo, to)
	})

	t.Run(fmt.Sprintf("func(*F, *T) error => func(%v, %T) => %v", nil, &to, time.Time{}), func(t *testing.T) {
		to = zero
		assert.NoError(t, converters.Convert((*From)(nil), &to))
		assert.Equal(t, zero, to)
	})

	t.Run(fmt.Sprintf("func(*F, **T) error => func(%v, %T) => %v", &from, &ptTo, expectedTo), func(t *testing.T) {
		ptTo = nil
		assert.NoError(t, converters.Convert(&from, &ptTo))
		assert.Equal(t, expectedTo, *ptTo)
	})

	t.Run(fmt.Sprintf("func(*F, **T) error => func(%v, %T) => %v", nil, &ptTo, nil), func(t *testing.T) {
		ptTo = nil
		assert.NoError(t, converters.Convert((*From)(nil), &ptTo))
		assert.Equal(t, (*To)(nil), ptTo)
	})

	var arrayTo []To
	expectedArrayTo := []To{expectedTo}
	t.Run(fmt.Sprintf("func(F, *[]T) error => func(%v, %T) => %v", from, &arrayTo, expectedArrayTo), func(t *testing.T) {
		assert.NoError(t, converters.Convert(from, &arrayTo))
		if assert.Equal(t, 1, len(arrayTo)) {
			assert.Equal(t, expectedArrayTo, arrayTo)
		}
	})

	arrayTo = nil
	t.Run(fmt.Sprintf("func(*F, *[]T) error => func(%v, %T) => %v", &from, &arrayTo, expectedArrayTo), func(t *testing.T) {
		assert.NoError(t, converters.Convert(&from, &arrayTo))
		if assert.Equal(t, 1, len(arrayTo)) {
			assert.Equal(t, expectedArrayTo, arrayTo)
		}
	})

	arrayTo = nil
	arrayFrom := []From{from}
	t.Run(fmt.Sprintf("func([]F, *[]T) error => func(%v, %T) => %v", arrayFrom, &arrayTo, expectedArrayTo), func(t *testing.T) {
		assert.NoError(t, converters.Convert(arrayFrom, &arrayTo))
		if assert.Equal(t, 1, len(arrayTo)) {
			assert.Equal(t, expectedArrayTo, arrayTo)
		}
	})

	var fromDest From
	t.Run(fmt.Sprintf("func(F, *F) error => func(%v, %T) => %v", from, &fromDest, from), func(t *testing.T) {
		assert.NoError(t, converters.Convert(from, &fromDest))
		assert.Equal(t, from, fromDest)
	})

	var ptFromDest *From
	t.Run(fmt.Sprintf("func(*F, **F) error => func(%v, %T) => %v", &from, &ptFromDest, from), func(t *testing.T) {
		assert.NoError(t, converters.Convert(&from, &ptFromDest))
		assert.Equal(t, from, *ptFromDest)
	})

	ptFromDest = nil
	t.Run(fmt.Sprintf("func(F, **F) error => func(%v, %T) => %v", from, &ptFromDest, from), func(t *testing.T) {
		assert.NoError(t, converters.Convert(from, &ptFromDest))
		assert.Equal(t, from, *ptFromDest)
	})

	var zeroFrom From
	fromDest = zeroFrom
	t.Run(fmt.Sprintf("func(*F, *F) error => func(%v, %T) => %v", &from, &fromDest, from), func(t *testing.T) {
		assert.NoError(t, converters.Convert(&from, &fromDest))
		assert.Equal(t, from, fromDest)
	})

	arrayFrom = nil
	t.Run(fmt.Sprintf("func(F, *[]F) error => func(%v, %T) => [%v]", from, &arrayFrom, from), func(t *testing.T) {
		assert.NoError(t, converters.Convert(from, &arrayFrom))
		if assert.Equal(t, 1, len(arrayFrom)) {
			assert.Equal(t, from, arrayFrom[0])
		}
	})

	arrayFrom = nil
	t.Run(fmt.Sprintf("func(*F, *[]F) error => func(%v, %T) => [%v]", &from, &arrayFrom, from), func(t *testing.T) {
		assert.NoError(t, converters.Convert(&from, &arrayFrom))
		if assert.Equal(t, 1, len(arrayFrom)) {
			assert.Equal(t, from, arrayFrom[0])
		}
	})

	var arrayFromDest []From
	arrayFrom = []From{from}
	t.Run(fmt.Sprintf("func([]F, *[]F) error => func(%v, %T) => [%v]", arrayFrom, &arrayFromDest, from), func(t *testing.T) {
		assert.NoError(t, converters.Convert(arrayFrom, &arrayFromDest))
		if assert.Equal(t, 1, len(arrayFromDest)) {
			assert.Equal(t, from, arrayFromDest[0])
		}
	})

	to = expectedTo
	var toDest To
	t.Run(fmt.Sprintf("func(T, *T) error => func(%v, %T) => %v", to, &toDest, to), func(t *testing.T) {
		assert.NoError(t, converters.Convert(to, &toDest))
		assert.Equal(t, to, toDest)
	})

	var ptToDest *To
	t.Run(fmt.Sprintf("func(*T, **T) error => func(%v, %T) => %v", &to, &ptToDest, to), func(t *testing.T) {
		assert.NoError(t, converters.Convert(&to, &ptToDest))
		assert.Equal(t, to, *ptToDest)
	})

	t.Run(fmt.Sprintf("func(T, **T) error => func(%v, %T) => %v", to, &ptToDest, to), func(t *testing.T) {
		assert.NoError(t, converters.Convert(to, &ptToDest))
		assert.Equal(t, to, *ptToDest)
	})

	t.Run(fmt.Sprintf("func(*T, *T) error => func(%v, %T) => %v", &to, &toDest, to), func(t *testing.T) {
		assert.NoError(t, converters.Convert(&to, &toDest))
		assert.Equal(t, to, toDest)
	})

	var arrayOfToDest []To
	t.Run(fmt.Sprintf("func(T, *[]T) error => func(%v, %T) => %v", to, &arrayOfToDest, expectedArrayTo), func(t *testing.T) {
		assert.NoError(t, converters.Convert(&to, &arrayOfToDest))
		assert.Equal(t, expectedArrayTo, arrayOfToDest)
	})

	t.Run(fmt.Sprintf("func(*T, *[]T) error => func(%v, %T) => %v", &to, &arrayOfToDest, expectedArrayTo), func(t *testing.T) {
		assert.NoError(t, converters.Convert(&to, &arrayOfToDest))
		assert.Equal(t, expectedArrayTo, arrayOfToDest)
	})

	arrayOfTo := []To{to}
	t.Run(fmt.Sprintf("func([]T, *[]T) error => func(%v, %T) => %v", arrayOfTo, &arrayOfToDest, expectedArrayTo), func(t *testing.T) {
		assert.NoError(t, converters.Convert(arrayOfTo, &arrayOfToDest))
		assert.Equal(t, expectedArrayTo, arrayOfToDest)
	})
}
