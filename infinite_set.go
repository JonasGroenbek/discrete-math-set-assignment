package assignment3

import (
	"errors"
)

type infiniteSet struct {
	min float64
	max float64
}

func NewInfiniteSet() infiniteSet {
	return infiniteSet{min: nInf(), max: inf()}
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
		if rs.lowBoundary == nInf() && rs.highBoundary == inf() {
			return CompositeSet{}
		} else {
			if rs.lowBoundary != nInf() {
				sets = append(sets, RangeSet{nInf(), rs.lowBoundary - 1})
			}
			if rs.highBoundary != inf() {
				sets = append(sets, RangeSet{rs.highBoundary + 1, inf()})
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
		if rs.lowBoundary == nInf() && rs.highBoundary == inf() {
			return CompositeSet{}, nil
		} else {
			return CompositeSet{}, errors.New("the universal set does not include element")
		}
	default:
		return CompositeSet{}, nil
	}
}
