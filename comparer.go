package mint

import "reflect"

func getComparer(a, b interface{}, deeply bool) Comparer {
	if deeply {
		return deepComparer{}
	}
	switch reflect.ValueOf(a).Kind() {
	case reflect.Slice:
		return sliceComparer{}
	case reflect.Map:
		return mapComparer{}
	}
	return defaultComparer{}
}

type Comparer interface {
	Compare(a, b interface{}) bool
}

type defaultComparer struct{}

func (c defaultComparer) Compare(a, b interface{}) bool {
	return a == b
}

type deepComparer struct{}

func (c deepComparer) Compare(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

type mapComparer struct {
	deepComparer
}

type sliceComparer struct {
	deepComparer
}
