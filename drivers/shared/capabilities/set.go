package capabilities

import (
	"sort"
	"strings"
)

type nothing struct{}

var null = nothing{}

// Set represents a group linux capabilities, implementing some useful set
// operations, taking care of name normalization.
type Set struct {
	data map[string]nothing
}

func New(caps []string) *Set {
	m := make(map[string]nothing, len(caps))
	for _, c := range caps {
		m[normalize(c)] = null
	}
	return &Set{data: m}
}

//func (s *Set) Union(b *Set) *Set {
//	data := make(map[string]nothing)
//	for c := range s.data {
//		data[c] = null
//	}
//	for c := range b.data {
//		data[c] = null
//	}
//	return &Set{data: data}
//}

// Add cap into s.
func (s *Set) Add(cap string) {
	name := normalize(cap)
	if name == "all" {
		sys := Supported()
		s.data = sys.data
		return
	}
	s.data[normalize(cap)] = null
}

// Remove caps from s.
func (s *Set) Remove(caps []string) {
	for _, c := range caps {
		name := normalize(c)
		if name == "all" {
			s.data = make(map[string]nothing)
			return
		}
		delete(s.data, name)
	}
}

// Difference returns the Set of elements of b not in s.
func (s *Set) Difference(b *Set) *Set {
	data := make(map[string]nothing)
	for c := range b.data {
		if _, exists := s.data[c]; !exists {
			data[c] = null
		}
	}
	return &Set{data: data}
}

// Empty return true if no capabilities exist in s.
func (s *Set) Empty() bool {
	return len(s.data) == 0
}

// String returns the normalized and sorted string representation of s.
func (s *Set) String() string {
	return strings.Join(s.Slice(), ",")
}

// Slice returns a sorted slice of capabilities in s.
func (s *Set) Slice() []string {
	caps := make([]string, 0, len(s.data))
	for c := range s.data {
		caps = append(caps, c)
	}
	sort.Strings(caps)
	return caps
}

// linux capabilities are often named in 4 possible ways - upper or lower case,
// and with or without a CAP_ prefix
//
// since we must do comparisons on cap names, always normalize the names before
// letting them into the Set data-structure
func normalize(name string) string {
	lower := strings.ToLower(name)
	trim := strings.TrimPrefix(lower, "cap_")
	return trim
}
