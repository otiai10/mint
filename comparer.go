package mint

import "reflect"

func getComparer(a, b interface{}, deeply bool) Comparer {
	if deeply {
		return deepComparer{}
	}
	if reflect.ValueOf(a).Kind() == reflect.Slice {
		return sliceComparer{}
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
