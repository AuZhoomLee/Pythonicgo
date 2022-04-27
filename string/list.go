package string

// InStringList judge whether element is in the list
// python: i in list
func InList(idea string, src []string) bool {
	for _, i := range src {
		if idea == i {
			return true
		}
	}
	return false
}

// LeastOneExist list whether empty
func LeastOneExist(src []string) bool {
	for _, i := range src {
		if i != "" {
			return true
		}
	}
	return false
}

// ToMap convert string list to a map
func ToMap(src []string) map[string]string {
	m := make(map[string]string, len(src))
	for _, i := range src {
		m[i] = ""
	}
	return m
}

// EqualIfPresent compared two strings for equality If both are present
// python: a && b && a==b
func EqualIfPresent(a string, b string) bool {
	if a == "" || b == "" {
		return false
	}
	return a == b
}
