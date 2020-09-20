package main

import (
	"errors"
	"math"
	"sort"
)

type RangeSet struct {
	lowBoundary  float64
	highBoundary float64
}

func (this RangeSet) Union(s Set) Result {
	switch s.(type) {
	case FiniteSet:
		fs := s.(FiniteSet)
		rs := s.(RangeSet)
		for k := range fs.set {
			if _, ok := fs.set[k]; ok {
				delete(fs.set, k)
			}
		}
		return Result{[]Set{
			fs,
			rs,
		}}
	case RangeSet:
		rs := s.(RangeSet)
		var low float64
		var high float64
		if rs.lowBoundary <= this.lowBoundary {
			low = rs.lowBoundary
		} else {
			low = this.lowBoundary
		}
		if rs.highBoundary >= this.highBoundary {
			high = rs.highBoundary
		} else {
			high = this.highBoundary
		}
		return Result{[]Set{
			RangeSet{low, high},
		}}
	default:
		return Result{[]Set{
			NewInfiniteSet(),
		}}
	}
}

func (this RangeSet) Intersection(s Set) Result {
	var res Result
	switch s.(type) {
	case FiniteSet:
		fs := s.(FiniteSet)
		thisLength := int(this.highBoundary - this.lowBoundary)
		fsLength := len(fs.set)
		intersections := FiniteSet{}
		if fsLength >= thisLength {
			for i := this.lowBoundary; i < this.highBoundary; i++ {
				if _, ok := fs.set[i]; ok {
					intersections.Add(i)
				}
			}
		} else {
			for k, _ := range fs.set {
				if isBetween(k, this.lowBoundary, this.highBoundary) {
					intersections.Add(k)
				}
			}
		}
		res = Result{[]Set{
			intersections,
		}}
	case RangeSet:
		rs := s.(RangeSet)
		var low float64
		var high float64
		//rs contains
		if rs.lowBoundary <= this.lowBoundary && rs.highBoundary >= this.lowBoundary {
			low = this.lowBoundary
			high = this.highBoundary
			//this contains
		} else if this.lowBoundary <= rs.lowBoundary && this.highBoundary >= rs.lowBoundary {
			low = rs.lowBoundary
			high = rs.highBoundary
			//they cross
		} else if isBetween(this.lowBoundary, rs.lowBoundary, rs.highBoundary) || isBetween(this.highBoundary, rs.lowBoundary, rs.highBoundary) {
			//equal low
			if this.lowBoundary == rs.lowBoundary {
				if this.highBoundary <= rs.highBoundary {
					low = rs.lowBoundary
					high = this.highBoundary
				} else {
					low = rs.lowBoundary
					high = rs.highBoundary
				}
				//equal high
			} else if this.highBoundary == rs.highBoundary {
				if this.lowBoundary <= rs.lowBoundary {
					low = this.lowBoundary
					high = this.highBoundary
				} else {
					low = rs.lowBoundary
					high = this.highBoundary
				}
			}
			//this lowest
		} else if this.lowBoundary < rs.lowBoundary {
			low = rs.lowBoundary
			high = this.highBoundary
			//rs lowest
		} else if rs.lowBoundary < this.lowBoundary {
			low = this.lowBoundary
			high = rs.highBoundary
		} else {
			return Result{}
		}
		return Result{
			[]Set{
				RangeSet{
					low,
					high,
				},
			},
		}
	case infiniteSet:
		res = Result{
			[]Set{
				NewInfiniteSet(),
			},
		}
	}
	return res
}

func (this RangeSet) Difference(s Set) Result {
	var res Result
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

		var prevKey float64
		for i, k := range keys {
			if isBetweenExcluding(k, this.lowBoundary, this.highBoundary) {
				//no previous key
				if i == 0 {
					rangeDifferences = append(rangeDifferences, RangeSet{this.lowBoundary, k - 1})
				} else {
					//check if the previous key is 1 apart
					if prevKey != k-1 {
						//check if there is a gap of 1 between previous key
						if prevKey == k-2 {
							finiteDifferences = append(finiteDifferences, k-1)
							//append range since gap is larger than 2
						} else {
							rangeDifferences = append(rangeDifferences, RangeSet{prevKey + 1, k - 1})
						}
					}
				}
				//check if key does not equal boundaries, if that is the case then they both include and no difference
			} else if k != this.lowBoundary && k != this.highBoundary {
				finiteDifferences = append(finiteDifferences, k)
			}
			prevKey = k
		}
		sets := make([]Set, 0)
		for _, rs := range rangeDifferences {
			sets = append(sets, Set(rs))
		}
		sets = append(sets, Set(NewFromSlice(finiteDifferences)))
		res = Result{sets}
	case RangeSet:

	case infiniteSet:
		sets := make([]Set, 0)
		if this.lowBoundary == math.Inf(-1) && this.highBoundary == math.Inf(1) {
			res = Result{}
		} else {
			if this.lowBoundary != math.Inf(-1) {
				sets = append(sets, RangeSet{math.Inf(-1), this.lowBoundary - 1})
			}
			if this.highBoundary != math.Inf(1) {
				sets = append(sets, RangeSet{this.highBoundary + 1, math.Inf(1)})
			}
			res = Result{sets}
		}
	}
	return res
}

func (this RangeSet) Complement(s Set) (Result, error) {
	switch s.(type) {
	case FiniteSet:
		fs := s.(FiniteSet)
		keys := make([]float64, 0)
		for k := range fs.set {
			keys = append(keys, k)
		}
		sort.Float64s(keys)
		if keys[len(keys)-1] > this.highBoundary || keys[0] < this.lowBoundary {
			return Result{}, errors.New("the universal set does not include element")
		} else {
			rangeDifferences := make([]RangeSet, 0)
			finiteDifferences := make([]float64, 0)

			var prevKey float64
			for i, k := range keys {
				//no previous key
				if i == 0 && k != math.Inf(-1) {
					rangeDifferences = append(rangeDifferences, RangeSet{math.Inf(-1), k - 1})
				} else {
					//check if the previous key is 1 apart
					if prevKey != k-1 {
						//check if there is a gap of 1 between previous key
						if prevKey == k-2 {
							finiteDifferences = append(finiteDifferences, k-1)
							//append range since gap is larger than 2
						} else {
							rangeDifferences = append(rangeDifferences, RangeSet{prevKey + 1, k - 1})
						}
					}
				}
				//check if key does not equal boundaries, if that is the case then they both include and no difference
				prevKey = k
			}
			sets := make([]Set, 0)
			for _, rs := range rangeDifferences {
				sets = append(sets, Set(rs))
			}
			sets = append(sets, Set(NewFromSlice(finiteDifferences)))
			return Result{sets}, nil
		}
	case RangeSet:
		rs := s.(RangeSet)
		if this.lowBoundary < rs.lowBoundary || this.highBoundary > rs.highBoundary {
			return Result{}, errors.New("the universal set does not include element")
		} else {
			sets := make([]Set, 0)
			var low float64
			var high float64
			if rs.lowBoundary >= this.highBoundary {
				high = rs.highBoundary
			} else {
				high = this.highBoundary
			}

			if rs.lowBoundary >= this.lowBoundary {
				low = this.lowBoundary
			} else {
				low = rs.lowBoundary
			}
			if high != math.Inf(1) {
				sets = append(sets, RangeSet{high, math.Inf(1)})
			}
			if low != math.Inf(-1) {
				sets = append(sets, RangeSet{math.Inf(-1), low})
			}
			return Result{sets}, nil
		}
	default:
		return Result{}, nil
	}
}
