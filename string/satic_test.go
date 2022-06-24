package string

import "testing"

func TestEqualIfPresent(t *testing.T) {
	var (
		a     = "a"
		empty = ""
		b     = "b"
	)
	cases := []struct {
		Name, A, B string
		Expected   bool
	}{
		{"both empty", empty, empty, false},
		{"left empty", empty, a, false},
		{"right empty", a, empty, false},
		{"equal", a, a, true},
		{"not equal", a, b, false},
	}

	for no, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if c.Expected != EqualIfPresent(c.A, c.B) {
				t.Errorf("No.%d name: %s , expected %v, got %v", no, c.Name, c.Expected, !c.Expected)
			}
		})
	}
}

func TestLeastOneExists(t *testing.T) {
	var (
		empty []string
		a     = []string{"a"}
	)
	cases := []struct {
		Name     string
		A        []string
		Expected bool
	}{
		{"empty", empty, false},
		{"not empty", a, true},
	}

	for no, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if c.Expected != LeastOneExist(c.A) {
				t.Errorf("No.%d name: %s, expected %v, got %v", no, c.Name, c.Expected, !c.Expected)
			}
		})
	}
}

func TestMapKeyEqual(t *testing.T) {
	var (
		empty = map[string]string{}
		a     = map[string]string{"a": ""}
		b     = map[string]string{"b": ""}
		ab    = map[string]string{"a": "", "b": ""}
	)
	cases := []struct {
		Name     string
		A, B     map[string]string
		Expected bool
	}{
		{"both empty", empty, empty, true},
		{"left empty", a, empty, false},
		{"right empty", empty, a, false},
		{"not any equal", a, b, false},
		{"left some not exist", a, ab, false},
		{"right some not exist", ab, a, false},
		{"equal", ab, ab, true},
	}
	for no, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if o := MapKeysEqual(c.A, c.B); o != c.Expected {
				t.Errorf("No.%d name: %s , expected %v, got %v", no, c.Name, c.Expected, o)
			}
		})

	}
}

func TestToMap(t *testing.T) {
	var (
		empty []string
		a     = []string{"a"}
	)
	cases := []struct {
		Name     string
		A        []string
		Expected map[string]string
	}{
		{"empty", empty, map[string]string{}},
		{"not empty", a, map[string]string{"a": ""}},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if o := ToMap(c.A); !MapKeysEqual(o, c.Expected) {
				t.Errorf("name: %s , expected %v, got %v", c.Name, c.Expected, o)
			}
		})
	}
}

func TestInList(t *testing.T) {
	var (
		empty        = ""
		emptyList    []string
		notEmptyList = []string{""}
		a            = "a"
		al           = []string{"a"}
	)
	cases := []struct {
		Name, A  string
		B        []string
		Expected bool
	}{
		{"empty in empty", empty, emptyList, false},
		{"empty in not empty", empty, notEmptyList, true},
		{"not in", a, notEmptyList, false},
		{"in", a, al, true},
	}

	for no, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if o := InList(c.A, c.B); o != c.Expected {
				t.Errorf("No.%d name: %s , expected %v, got %v", no, c.Name, c.Expected, o)
			}
		})
	}
}
