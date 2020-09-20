package main

import (
	"fmt"
	"reflect"
)

//FiniteSet something
type (
	FiniteSet struct {
		set map[float64]nothing
	}
	nothing struct{}
)

func (this FiniteSet) Member(f float64) bool {
	return false
}

func (this FiniteSet) Union(s Set) UnionSet {
	var us UnionSet
	switch s.(type) {
	case FiniteSet:
		this2 := s.(FiniteSet)
		for k := range this2.set {
			if _, ok := this.set[k]; ok {
				delete(this.set, k)
			}
		}
		us = UnionSet{this, this2}
	case RangeSet:
		rs := s.(RangeSet)
		for k := range this.set {
			if isBetween(k, rs.lowerBoundary, rs.upperBoundary) {
				delete(this.set, k)
			}
		}
		us = UnionSet{this, rs}
	case InfiniteSet:
		is := s.(InfiniteSet)
		for k := range this.set {
			delete(this.set, k)
		}
		us = UnionSet{this, is}
	}
	return us
}

func (this FiniteSet) Intersection(s Set) IntersectionSet {
	var is IntersectionSet
	switch s.(type) {
	case FiniteSet:
		intersection := New()
		this2 := s.(FiniteSet)
		if len(this2.set) >= len(this.set) {
			for k := range this.set {
				if _, ok := this.set[k]; !ok {
					intersection.Add(k)
				}
			}
		} else {
			for k := range this.set {
				if _, ok := this2.set[k]; !ok {
					intersection.Add(k)
				}
			}
		}
		is = IntersectionSet{[]Set{intersection}}
	case RangeSet:
		rs := s.(RangeSet)
		rsLength := int(rs.upperBoundary - rs.lowerBoundary)
		thisLength := len(this.set)
		intersection := New()
		if thisLength >= rsLength {
			for i := rs.lowerBoundary; i < rs.upperBoundary; i++ {
				if _, ok := this.set[i]; ok {
					intersection.Add(i)
				}
			}
		} else {
			for k, _ := range this.set {
				if isBetween(k, rs.lowerBoundary, rs.upperBoundary) {
					intersection.Add(k)
				}
			}
		}
		is = IntersectionSet{[]Set{intersection}}
	case InfiniteSet:
		is = IntersectionSet{[]Set{s}}
	}
	return is
}

func (this FiniteSet) Difference(s Set) DifferenceSet {
	var ds DifferenceSet
	switch s.(type) {
	case FiniteSet:
		differences := New()
		fs := s.(FiniteSet)
		for k, _ := range fs.set {
			if _, ok := this.set[k]; ok {
				differences.Add(k)
			}
		}
		for k, _ := range this.set {
			if _, ok := fs.set[k]; ok {
				differences.Add(k)
			}
		}
	case RangeSet:
		fmt.Printf("RangeSet")
	case InfiniteSet:
		fmt.Printf("InfiniteSet")
	}
	return ds
}

func (this FiniteSet) Complement(s Set) ComplementSet {
	var cs ComplementSet
	switch s.(type) {
	case FiniteSet:
		fmt.Printf("FiniteSet")
	case RangeSet:
		fmt.Printf("RangeSet")
	case InfiniteSet:
		fmt.Printf("InfiniteSet")
	}
	return cs
}

func (this FiniteSet) CompareTo(s Set) int {
	var comparison int
	switch s.(type) {
	case FiniteSet:
		if reflect.DeepEqual(this, s) {
			return 0
		}
	case RangeSet:
		fmt.Printf("RangeSet")
	case InfiniteSet:
		fmt.Printf("InfiniteSet")
	}
	return comparison
}

func New() FiniteSet {
	return FiniteSet{set: make(map[float64]nothing)}
}

func (this FiniteSet) Add(k float64) {
	this.set[k] = nothing{}
}

func (this FiniteSet) Remove(k float64) error {
	_, exists := this.set[k]
	if !exists {
		return fmt.Errorf("Remove Error: Item doesn't exist in set")
	}
	delete(this.set, k)
	return nil
}

func (this FiniteSet) Size() int {
	return len(this.set)
}

func isBetween(target, lowerBoundary, upperBoundary float64) bool {
	return target >= lowerBoundary && target <= upperBoundary
}
