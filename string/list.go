package string

import "sort"

const (
	TypeString = "string"
	TypeSlice  = "[]string"
	TypeList   = "List"
)

type List []string

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

func (l List) Index(v string) int {
	for no, i := range l {
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
		switch v := v.(type) {
		case string:
			fm[TypeString](v)
		case List:
			fm[TypeList](v)
		case []string:
			fm[TypeSlice](v)
		}
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

type OrderList []string
