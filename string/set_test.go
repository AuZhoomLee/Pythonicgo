package string

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	var (
		a = "a"
		s Set
	)

	s.Add(a)
	v := reflect.ValueOf(s)
	if !(v.IsValid() &&
		v.FieldByName("len").Int() == s.Len() &&
		v.FieldByName("value").MapIndex(reflect.ValueOf(a)).String() == reflect.ValueOf("").String()) {
		t.Errorf("Name: Set.Add, Expected %v, got %v", s, v)
	}
}

func TestDel(t *testing.T) {
	var (
		a = "a"
		s Set
	)

	reflect.ValueOf(s).Set()

}
