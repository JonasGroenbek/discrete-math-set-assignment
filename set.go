package main

//Set something
type Set struct {
	set map[int]int
	i   int
}

//New insantiates a set with a map
func New() Set {
	return Set{make(map[int]int), 100}
}

//Add adds an element to the set
func (s Set) Add(key, value int) {
	s.set[key] = value
}

//ChangeI stuff
func (s *Set) ChangeI(i int) {
	s.i = 1000
}
