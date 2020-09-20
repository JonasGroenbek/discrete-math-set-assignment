package main

import (
	"fmt"
	"reflect"
	"sort"
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
		rs := s.(RangeSet)
		//creating range differences and specific differences
		rangeDifferences := make([]RangeSet, 0)
		finiteDifferences := make([]float64, 0)

		//sorting by finite keys
		keys := make([]float64, 0)
		for k := range this.set {
			keys = append(keys, k)
		}
		sort.Float64s(keys)

		var previousKey float64
		for i, k := range keys {
			if isBetweenExcluding(k, rs.lowerBoundary, rs.upperBoundary) {
				//no previous key
				if i == 0 {
					rangeDifferences = append(rangeDifferences, RangeSet{rs.lowerBoundary, k - 1})
				} else {
					//check if the previous key is 1 apart
					if previousKey != k-1 {
						//check if there is a gap of 1 between previous key
						if previousKey == k-2 {
							finiteDifferences = append(finiteDifferences, k-1)
							//append range since gap is larger than 2
						} else {
							rangeDifferences = append(rangeDifferences, RangeSet{previousKey + 1, k - 1})
						}
					}
				}
				//check if key does not equal boundaries, if that is the case then they both include and no difference
			} else if k != rs.lowerBoundary && k != rs.upperBoundary {
				finiteDifferences = append(finiteDifferences, k)
			}
			previousKey = k
		}
		sets := make([]Set, 0)
		for _, rs := range rangeDifferences {
			sets = append(sets, Set(rs))
		}
		sets = append(sets, Set(NewFromSlice(finiteDifferences)))
		ds = DifferenceSet{sets}
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

func NewFromSlice(values []float64) FiniteSet {
	set := make(map[float64]nothing)
	for _, k := range values {
		set[k] = nothing{}
	}
	return FiniteSet{set}
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

func isBetweenExcluding(target, lowerBoundary, upperBoundary float64) bool {
	return target > lowerBoundary && target < upperBoundary
}
