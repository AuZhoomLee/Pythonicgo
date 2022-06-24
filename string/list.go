package string

import "sort"

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

func (l *List) Push(values ...interface{}) *List {
	p := func(v string) {
		*l = append(*l, v)
	}
	ps := func(v []string) {
		*l = append(*l, v...)
	}

	for _, v := range values {
		switch v := v.(type) {
		case string:
			p(v)
		case List:
			ps(v)
		case []string:
			ps(v)
		}
	}

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

func (l *List) Insert(index int, overflow bool, values ...interface{}) *List {
	if l.Len()-1 > index {
		if overflow {
			l.Push(values)
		}
		return l
	}
	// todo: insert in this list
	return l
}

func (l *List) Delete(index int, values ...interface{}) *List {
	// todo: delete from this list
	return l
}

type OrderList []string
