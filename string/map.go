package string

type Map map[string]string

func (m Map) Have(k string) bool {
	_, ok := m[k]

	return ok
}

func (m Map) Get(k, v string) string {
	if v, ok := m[k]; ok {
		return v
	}
	return v
}

func (m *Map) Put(k, v string) *Map {
	(*m)[k] = v
	return m
}

func (m *Map) Update(n map[string]string) *Map {
	for k, v := range n {
		m.Put(k, v)
	}
	return m
}

func (m *Map) Extend(n map[string]string) *Map {
	for k, v := range n {
		if !m.Have(k) {
			m.Put(k, v)
		}
	}
	return m
}

func (m Map) Keys() (ks []string) {
	for k := range m {
		ks = append(ks, k)
	}
	return
}

func (m Map) Len() (l int) {
	for range m {
		l += 1
	}
	return
}

func (m Map) Values() (vs []string) {
	for _, v := range m {
		vs = append(vs, v)
	}
	return
}

func (m Map) DistinctValues() (vs []string) {
	n := make(Map)
	for _, v := range m {
		if !n.Have(v) {
			vs = append(vs, v)
			n.Put(v, "")
		}
	}
	return
}

func (m Map) Items() (is [][2]string) {
	for k, v := range m {
		is = append(is, [2]string{k, v})
	}
	return
}
