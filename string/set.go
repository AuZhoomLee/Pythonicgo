package string

type Set struct {
	value map[string]struct{}
	len   int64
}

func (s *Set) ToList() []string {
	l := make([]string, 0, len(s.value))
	for k := range s.value {
		l = append(l, k)
	}
	return l
}

func (s *Set) ToString() (r string) {
	for k := range s.value {
		r += k
	}
	return
}

func (s *Set) Has(idea string) bool {
	_, ok := s.value[idea]
	return ok
}

func (s *Set) HasSome(sl []string) (has, hasnot []string) {
	has = make([]string, s.len/2)
	hasnot = make([]string, s.len/2)
	for _, i := range sl {
		if s.Has(i) {
			has = append(has, i)
			continue
		}
		hasnot = append(hasnot, i)
	}
	return
}

func (s *Set) HasAny(sl []string) bool {
	for _, i := range sl {
		if s.Has(i) {
			return true
		}
	}
	return false
}

func (s *Set) HasAll(sl []string) bool {
	for _, i := range sl {
		if !s.Has(i) {
			return false
		}
	}
	return true
}

func (s *Set) Add(idea string) *Set {
	if !s.Has(idea) {
		s.value[idea] = struct{}{}
		s.len += 1
	}
	return s
}

func (s *Set) AddList(sl []string) *Set {
	for _, i := range sl {
		s.Add(i)
	}
	return s
}

func (s *Set) Del(idea string) *Set {
	delete(s.value, idea)
	s.len -= 1
	return s
}

func (s *Set) UnionList(sl []string) *Set {
	for _, i := range sl {
		s.Add(i)
	}
	return s
}
