package main

import (
	"fmt"
	"reflect"
)

type InfiniteSet struct{}

func (is InfiniteSet) Member(f float64) bool {
	return true
}

func (is InfiniteSet) Union(s Set) UnionSet {
	var us UnionSet
	switch s.(type) {
	case FiniteSet:
		fmt.Printf("FiniteSet")
	case RangeSet:
		fmt.Printf("RangeSet")
	case InfiniteSet:
		fmt.Printf("InfiniteSet")
	}
	return us
}

func (is InfiniteSet) Intersection(s Set) IntersectionSet {
	var iss IntersectionSet
	switch s.(type) {
	case FiniteSet:
		fmt.Printf("FiniteSet")
	case RangeSet:
		fmt.Printf("RangeSet")
	case InfiniteSet:
		fmt.Printf("InfiniteSet")
	}
	return iss
}

func (is InfiniteSet) Difference(s Set) DifferenceSet {
	var ds DifferenceSet
	switch s.(type) {
	case FiniteSet:
		fmt.Printf("FiniteSet")
	case RangeSet:
		fmt.Printf("RangeSet")
	case InfiniteSet:
		fmt.Printf("InfiniteSet")
	}
	return ds
}

func (is InfiniteSet) Complement(s Set) ComplementSet {
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

func (is InfiniteSet) CompareTo(s Set) int {
	if reflect.DeepEqual(is, s) {
		return 0
	}
	return 1
}
