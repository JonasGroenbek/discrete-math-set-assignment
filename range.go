package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

type RangeSet struct {
	lowerBoundary float64
	upperBoundary float64
}

func (this RangeSet) New(l, u float64) RangeSet {
	return RangeSet{l, u}
}

func (this RangeSet) Union(s Set) UnionSet {
	var us UnionSet
	switch s.(type) {
	case FiniteSet:
		fs := s.(FiniteSet)
		rs := s.(RangeSet)
		for k := range fs.set {
			if _, ok := fs.set[k]; ok {
				delete(fs.set, k)
			}
		}
		us = UnionSet{this, rs}
	case RangeSet:
		rs := s.(RangeSet)
		var lower float64
		var upper float64
		if rs.lowerBoundary <= this.lowerBoundary {
			lower = rs.lowerBoundary
		} else {
			lower = this.lowerBoundary
		}
		if rs.upperBoundary >= this.upperBoundary {
			upper = rs.upperBoundary
		} else {
			upper = this.upperBoundary
		}
		us = UnionSet{RangeSet{lower, upper}, New()}
	case InfiniteSet:
		us = UnionSet{InfiniteSet{}, New()}
	}
	return us
}

func (this RangeSet) Intersection(s Set) IntersectionSet {
	var is IntersectionSet
	switch s.(type) {
	case FiniteSet:
		fs := s.(FiniteSet)
		thisLength := int(this.upperBoundary - this.lowerBoundary)
		fsLength := len(fs.set)
		intersection := New()
		if fsLength >= thisLength {
			for i := this.lowerBoundary; i < this.upperBoundary; i++ {
				if _, ok := fs.set[i]; ok {
					intersection.Add(i)
				}
			}
		} else {
			for k, _ := range fs.set {
				if isBetween(k, this.lowerBoundary, this.upperBoundary) {
					intersection.Add(k)
				}
			}
		}
		is = IntersectionSet{[]Set{intersection}}
	case RangeSet:
		fmt.Printf("RangeSet")
	case InfiniteSet:
		is = IntersectionSet{
			[]Set{
				InfiniteSet{},
			},
		}
	}
	return is
}

func (this RangeSet) Difference(s Set) DifferenceSet {
	var ds DifferenceSet
	switch s.(type) {
	case FiniteSet:
		fs := s.(FiniteSet)
		//creating range differences and specific differences
		rangeDifferences := make([]RangeSet, 0)
		finiteDifferences := make([]float64, 0)

		//sorting by finite keys
		keys := make([]float64, 0)
		for k := range fs.set {
			keys = append(keys, k)
		}

		sort.Float64s(keys)

		var previousKey float64
		for i, k := range keys {
			if isBetweenExcluding(k, this.lowerBoundary, this.upperBoundary) {
				//no previous key
				if i == 0 {
					rangeDifferences = append(rangeDifferences, RangeSet{this.lowerBoundary, k - 1})
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
			} else if k != this.lowerBoundary && k != this.upperBoundary {
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
	case RangeSet:
		fmt.Printf("RangeSet")
	case InfiniteSet:
		sets := make([]Set, 0)
		if this.lowerBoundary == math.Inf(-1) && this.upperBoundary == math.Inf(1) {
			ds = DifferenceSet{}
		} else {
			if this.lowerBoundary != math.Inf(-1) {
				sets = append(sets, RangeSet{math.Inf(-1), this.lowerBoundary - 1})
			}
			if this.upperBoundary != math.Inf(1) {
				sets = append(sets, RangeSet{this.upperBoundary + 1, math.Inf(1)})
			}
			ds = DifferenceSet{sets}
		}
	}
	return ds
}

func (this RangeSet) Complement(s Set) ComplementSet {
	var cs ComplementSet
	switch s.(type) {
	case FiniteSet:
		fs := s.(FiniteSet)
		keys := make([]float64, 0)
		for k := range fs.set {
			keys = append(keys, k)
		}
		sort.Float64s(keys)
		if keys[len(keys)-1] > this.upperBoundary || keys[0] < this.lowerBoundary {
			panic("the universal set does not include element")
		} else {
			rangeDifferences := make([]RangeSet, 0)
			finiteDifferences := make([]float64, 0)

			var previousKey float64
			for i, k := range keys {
				//no previous key
				if i == 0 && k != math.Inf(-1) {
					rangeDifferences = append(rangeDifferences, RangeSet{math.Inf(-1), k - 1})
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
				previousKey = k
			}
			sets := make([]Set, 0)
			for _, rs := range rangeDifferences {
				sets = append(sets, Set(rs))
			}
			sets = append(sets, Set(NewFromSlice(finiteDifferences)))
			cs = ComplementSet{sets}
		}
	case RangeSet:
		fmt.Printf("RangeSet")
	case InfiniteSet:
		cs = ComplementSet{}
	}
	return cs
}

func (this RangeSet) CompareTo(s Set) int {
	if reflect.DeepEqual(this, s) {
		return 0
	}
	return 1
}
