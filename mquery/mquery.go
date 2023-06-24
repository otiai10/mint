package mquery

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Query(m any, q string) any {
	return query(m, strings.Split(q, "."))
}

func query(m any, qs []string) any {
	t := reflect.TypeOf(m)
	switch t.Kind() {
	case reflect.Map:
		return queryMap(m, t, qs)
	case reflect.Slice:
		return querySlice(m, t, qs)
	default:
		return m
	}
}

func queryMap(m any, t reflect.Type, qs []string) any {
	if len(qs) == 0 {
		return m
	}
	switch t.Key().Kind() {
	case reflect.String:
		next := reflect.ValueOf(m).MapIndex(reflect.ValueOf(qs[0])).Interface()
		return query(next, qs[1:])
	case reflect.Int:
		i, err := strconv.Atoi(qs[0])
		if err != nil {
			return fmt.Errorf("cannot access map with keyword: %s: %v", qs[0], err)
		}
		next := reflect.ValueOf(m).MapIndex(reflect.ValueOf(i)).Interface()
		return query(next, qs[1:])
	}
	return nil
}

func querySlice(m any, t reflect.Type, qs []string) any {
	if len(qs) == 0 {
		return m
	}
	v := reflect.ValueOf(m)
	if v.Len() == 0 {
		return nil
	}
	i, err := strconv.Atoi(qs[0])
	if err != nil {
		return fmt.Errorf("cannot access slice with keyword: %s: %v", qs[0], err)
	}
	if v.Len() <= i {
		return nil
	}
	next := v.Index(i).Interface()
	return query(next, qs[1:])
}
