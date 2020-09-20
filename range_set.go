package main

import (
	"fmt"
	"reflect"
)

type RangeSet struct {
	lowerBoundary float64
	upperBoundary float64
}

func (rs RangeSet) New(l, u float64) RangeSet {
	return RangeSet{l, u}
}

func (rs RangeSet) Member(f float64) bool {
	if f >= rs.lowerBoundary && f <= rs.upperBoundary {
		return true
	}
	return false
}

func (rs RangeSet) Union(s Set) UnionSet {
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

func (rs RangeSet) Intersection(s Set) IntersectionSet {
	var is IntersectionSet
	switch s.(type) {
	case FiniteSet:
		fmt.Printf("FiniteSet")
	case RangeSet:
		fmt.Printf("RangeSet")
	case InfiniteSet:
		fmt.Printf("InfiniteSet")
	}
	return is
}

func (rs RangeSet) Difference(s Set) DifferenceSet {
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

func (rs RangeSet) Complement(s Set) ComplementSet {
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

func (rs RangeSet) CompareTo(s Set) int {
	if reflect.DeepEqual(rs, s) {
		return 0
	}
	return 1
}