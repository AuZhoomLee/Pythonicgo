package string

import "sort"

const (
	TypeString = "string"
	TypeSlice  = "[]string"
	TypeList   = "List"
	TypeOther  = "other"
)

type List []string

func ToList(s []string) (l *List) {
	return l.Push(s)
}

func typeOf(v interface{}) string {
	switch v.(type) {
	case string:
		return TypeString
	case []string:
		return TypeSlice
	case List:
		return TypeList
	}
	return TypeOther
}

func (l List) Len() int {
	return len(l)
}

func (l List) Less(i, j int) bool {
	return l[i] <= l[j]
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) IsEmpty() bool {
	return l.Len() < 1
}

func (l *List) RemoveFirst() *List {
	switch {
	case l.Len() == 1:
		l.Clear()
	case l.Len() > 1:
		*l = (*l)[1:]
	}
	return l
}

func (l *List) RemoveLast() *List {
	switch {
	case l.Len() == 1:
		l.Clear()
	case l.Len() > 1:
		*l = (*l)[:l.Len()-1]
	}
	return l
}

func (l List) Get(index int) (string, bool) {
	if index >= 0 && index < l.Len() {
		return l[index], true
	}
	return "", false
}

func (l List) IndexFirst(v string, index int) int {
	for no, i := range l[index:] {
		if i == v {
			return no
		}
	}
	return -1
}

func (l List) Sort() List {
	sort.Sort(l)
	return l
}

func (l *List) Sorted() *List {
	sort.Sort(*l)
	return l
}

func (l *List) Swapped(i, j int) *List {
	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
	return l
}

func (l *List) DealWithFunc(fm map[string]func(v interface{}), values ...interface{}) *List {
	for _, v := range values {
		fm[typeOf(v)](v)
	}
	return l
}

func (l *List) Push(values ...interface{}) *List {
	l.DealWithFunc(map[string]func(v interface{}){
		TypeString: func(v interface{}) {
			*l = append(*l, v.(string))
		},
		TypeList: func(v interface{}) {
			*l = append(*l, v.(List)...)
		},
		TypeSlice: func(v interface{}) {
			*l = append(*l, v.([]string)...)
		},
		TypeOther: func(v interface{}) {},
	}, values)

	return l
}

func (l *List) Pop() *List {
	if p := l.Len(); p > 0 {
		*l = (*l)[0 : p-1]
	}
	return l
}

func (l *List) PopN(n int) *List {
	for i := 0; i < n || l.Len() > 0; i++ {
		l.Pop()
	}
	return l
}

func (l *List) Clear() *List {
	if l.Len() > 0 {
		*l = (*l)[:0]
	}
	return l
}

func (l *List) Insert(index int, overflow bool, values ...interface{}) *List {
	switch true {
	case index >= l.Len():
		if overflow {
			l.Push(values)
		}
	case 0 < index && index < l.Len():
		tail := (*l)[index:]
		head := (*l)[:index]
		return head.Push(values).Push(tail)
	case index == 0:
		cp := (*l)[:]
		l.Clear().Push(values).Push(cp)
	}
	return l
}

func (l *List) Remove(index int) *List {
	switch {
	case l.Len() == 1:
		if index == 0 {
			l.Clear()
		}
	case l.Len() > 1:
		switch {
		case index == 0:
			l.RemoveFirst()
		case index == l.Len()-1:
			l.RemoveLast()
		case index > 0 && index < l.Len():
			*l = append((*l)[:index], (*l)[index+1:]...)
		}
	}
	return l
}

func (l List) ToPlainMap() map[string]struct{} {
	s := make(map[string]struct{})
	for _, i := range l {
		s[i] = struct{}{}
	}
	return s
}

func (l List) ToMapWithIndex() map[string]int {
	s := make(map[string]int)
	for no, i := range l {
		s[i] = no
	}
	return s
}

func (l *List) Delete(index int, values ...interface{}) *List {
	f := func(v interface{}) {
		var vm map[string]struct{}
		switch v := v.(type) {
		case []string:
			vm = ToList(v).ToPlainMap()
		case List:
			vm = v.ToPlainMap()
		}

		left := make(List, len((*l)[index:]))
		for _, i := range (*l)[index:] {
			if _, ok := vm[i]; !ok {
				left.Push(i)
			}
		}

		*l = append((*l)[:index], left...)
	}

	fm := map[string]func(v interface{}){
		TypeString: func(v interface{}) {
			pos := (*l)[index:].IndexFirst(v.(string), 0)
			l.Remove(pos)
		},
		TypeSlice: f,
		TypeList:  f,
		TypeOther: func(v interface{}) {},
	}

	l.DealWithFunc(fm, values)

	return l
}

type OrderList struct {
	*List
}

func ToOrderList(s []string) (o *OrderList) {
	return o.Push(s)
}

func (o *OrderList) Push(values ...interface{}) *OrderList {
	var todo List
	if length := len(values); length < 1 {
		return o
	} else {
		todo = make(List, 0, length)
	}

	todo.Push(values).Sorted()
	if o.IsEmpty() {
		o.List = &todo
		return o
	}
	backup := *(o.List)
	o.Clear()

	p, q := 0, 0
	for p < backup.Len() || q < todo.Len() {
		if backup[p] <= todo[q] {
			*(o.List) = append(*(o.List), backup[p])
			p++
		} else {
			*(o.List) = append(*(o.List), todo[q])
			q++
		}
	}

	if p < backup.Len() {
		*(o.List) = append(*(o.List), backup[p:]...)
		return o
	}

	*(o.List) = append(*(o.List), todo[q:]...)
	return o
}
