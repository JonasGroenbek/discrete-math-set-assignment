package assignment3

import (
	"errors"
	"math"
)

type infiniteSet struct {
	min float64
	max float64
}

func NewInfiniteSet() infiniteSet {
	return infiniteSet{min: math.Inf(-1), max: math.Inf(1)}
}

func (this infiniteSet) Union(s Set) CompositeSet {
	return CompositeSet{
		[]Set{
			NewInfiniteSet(),
		},
	}
}

func (this infiniteSet) Intersection(s Set) CompositeSet {
	return CompositeSet{
		[]Set{
			s,
		},
	}
}

func (this infiniteSet) Difference(s Set) CompositeSet {
	switch s.(type) {
	case RangeSet:
		rs := s.(RangeSet)
		sets := make([]Set, 0)
		if rs.lowBoundary == math.Inf(-1) && rs.highBoundary == math.Inf(1) {
			return CompositeSet{}
		} else {
			if rs.lowBoundary != math.Inf(-1) {
				sets = append(sets, RangeSet{math.Inf(-1), rs.lowBoundary - 1})
			}
			if rs.highBoundary != math.Inf(1) {
				sets = append(sets, RangeSet{rs.highBoundary + 1, math.Inf(1)})
			}
			return CompositeSet{sets}
		}
	default:
		return CompositeSet{}
	}
}

func (this infiniteSet) Complement(s Set) (CompositeSet, error) {
	switch s.(type) {
	case FiniteSet:
		return CompositeSet{}, errors.New("the universal set does not include element")
	case RangeSet:
		rs := s.(RangeSet)
		if rs.lowBoundary == math.Inf(-1) && rs.highBoundary == math.Inf(1) {
			return CompositeSet{}, nil
		} else {
			return CompositeSet{}, errors.New("the universal set does not include element")
		}
	default:
		return CompositeSet{}, nil
	}
}
