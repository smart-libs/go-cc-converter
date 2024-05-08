package converterdefault

import (
	"reflect"
	"sync"
)

type (
	fromToMap[T any] struct {
		matrix map[reflect.Type]map[reflect.Type]T
		sync.RWMutex
	}
)

func newFromToMap[T any]() fromToMap[T] {
	return fromToMap[T]{
		matrix: make(map[reflect.Type]map[reflect.Type]T),
	}
}

func (d *fromToMap[T]) Add(from, to reflect.Type, t T) *T {
	d.Lock()
	defer d.Unlock()
	currentTo, foundFrom := d.matrix[from]
	if !foundFrom {
		currentTo = make(map[reflect.Type]T)
		d.matrix[from] = currentTo
	}

	old, foundTo := currentTo[to]
	currentTo[to] = t
	if !foundTo {
		return nil
	}
	return &old
}

func (d *fromToMap[T]) Get(from, to reflect.Type) (t T, found bool) {
	d.RLock()
	defer d.RUnlock()
	if currentTo, foundFrom := d.matrix[from]; foundFrom {
		t, found = currentTo[to]
	}
	return
}
