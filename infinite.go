package main

import (
	"math"
)

type InfiniteSet struct{}

func (this InfiniteSet) Union(s Set) UnionSet {
	return UnionSet{InfiniteSet{}, New()}
}

func (this InfiniteSet) Intersection(s Set) IntersectionSet {
	return IntersectionSet{
		[]Set{
			s,
		},
	}
}

func (this InfiniteSet) Difference(s Set) DifferenceSet {
	var ds DifferenceSet
	switch s.(type) {
	case RangeSet:
		rs := s.(RangeSet)
		sets := make([]Set, 0)
		if rs.lowerBoundary == math.Inf(-1) && rs.upperBoundary == math.Inf(1) {
			ds = DifferenceSet{}
		} else {
			if rs.lowerBoundary != math.Inf(-1) {
				sets = append(sets, RangeSet{math.Inf(-1), rs.lowerBoundary - 1})
			}
			if rs.upperBoundary != math.Inf(1) {
				sets = append(sets, RangeSet{rs.upperBoundary + 1, math.Inf(1)})
			}
			ds = DifferenceSet{sets}
		}
	default:
		ds = DifferenceSet{}
	}
	return ds
}

func (this InfiniteSet) Complement(s Set) ComplementSet {
	var cs ComplementSet
	switch s.(type) {
	case FiniteSet:
		panic("the universal set does not include element")
	case RangeSet:
		rs := s.(RangeSet)
		if rs.lowerBoundary == math.Inf(-1) && rs.upperBoundary == math.Inf(1) {
			cs = ComplementSet{}
		} else {
			panic("the universal set does not include element")
		}
	case InfiniteSet:
		cs = ComplementSet{}
	}
	return cs
}
