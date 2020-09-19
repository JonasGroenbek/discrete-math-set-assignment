package main

import (
	"fmt"
	"reflect"
)

//Set something
type Set struct {
	container map[int]bool
}

//New insantiates a set with a map
func New() Set {
	return Set{container: make(map[int]bool)}
}

//Member check is a element is present in set
func (s Set) Member(k int) bool {
	_, ok := s.container[k]
	return ok
}

//Add adds an element to the set
func (s Set) Add(k int) {
	s.container[k] = true
}

func (s Set) Comparator(set Set){
	if reflect.DeepEqual(s, set) return 0
}

A ⊂ B: -1
• A = B: 0
• A ⊃ B: 1
• Undeterminable: 2
• A * B ∧ B * A: -2

//Remove removes an element from the set
func (s Set) Remove(key int) error {
	_, exists := s.container[key]
	if !exists {
		return fmt.Errorf("Remove Error: Item doesn't exist in set")
	}
	delete(s.container, key)
	return nil
}

//Size returns the amount of elements in the set
func (s Set) Size() int {
	return len(s.container)
}
