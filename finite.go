package assignment3

import (
	"errors"
	"math"
	"reflect"
	"sort"
)

type (
	FiniteSet struct {
		set map[float64]nothing
	}
	nothing struct{}
)

func (this FiniteSet) Union(s Set) Result {
	switch s.(type) {
	case FiniteSet:
		rs := s.(FiniteSet)
		for k := range rs.set {
			if _, ok := this.set[k]; ok {
				delete(this.set, k)
			}
		}
		return Result{[]Set{
			this,
			s,
		}}
	case RangeSet:
		rs := s.(RangeSet)

		keys := make([]float64, 0)
		for k := range this.set {
			keys = append(keys, k)
		}
		sort.Float64s(keys)

		var high float64
		var low float64
		for i, k := range keys {
			if i == 0 {
				high = k
				low = k
			}
			if k > high {
				high = k
			}
			if k < low {
				low = k
			}
		}
		if rs.highBoundary > high {
			high = rs.highBoundary
		}
		if rs.lowBoundary < low {
			low = rs.lowBoundary
		}
		return Result{[]Set{
			RangeSet{low, high},
		}}
	default:
		is := s.(infiniteSet)
		for k := range this.set {
			delete(this.set, k)
		}
		return Result{[]Set{
			this,
			is,
		}}
	}
}

func (this FiniteSet) Intersection(s Set) Result {
	switch s.(type) {
	case FiniteSet:
		intersections := FiniteSet{}
		rs := s.(FiniteSet)
		if len(rs.set) >= len(this.set) {
			for k := range this.set {
				if _, ok := this.set[k]; !ok {
					intersections.Add(k)
				}
			}
		} else {
			for k := range this.set {
				if _, ok := rs.set[k]; !ok {
					intersections.Add(k)
				}
			}
		}
		return Result{[]Set{
			intersections,
		}}
	case RangeSet:
		rs := s.(RangeSet)
		rsLength := int(rs.highBoundary - rs.lowBoundary)
		thisLength := len(this.set)
		intersections := FiniteSet{}
		if thisLength >= rsLength {
			for i := rs.lowBoundary; i < rs.highBoundary; i++ {
				if _, ok := this.set[i]; ok {
					intersections.Add(i)
				}
			}
		} else {
			for k, _ := range this.set {
				if isBetween(k, rs.lowBoundary, rs.highBoundary) {
					intersections.Add(k)
				}
			}
		}
		return Result{[]Set{
			intersections,
		}}
	default:
		return Result{[]Set{
			s,
		}}
	}
}

func (this FiniteSet) Difference(s Set) Result {
	switch s.(type) {
	case FiniteSet:
		diff := FiniteSet{}
		fs := s.(FiniteSet)
		for k, _ := range fs.set {
			if _, ok := this.set[k]; ok {
				diff.Add(k)
			}
		}
		for k, _ := range this.set {
			if _, ok := fs.set[k]; ok {
				diff.Add(k)
			}
		}
		return Result{
			[]Set{
				diff,
			}}
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

		var prevKey float64
		for i, k := range keys {
			if isBetweenExcluding(k, rs.lowBoundary, rs.highBoundary) {
				//no previous key
				if i == 0 {
					rangeDifferences = append(rangeDifferences, RangeSet{rs.lowBoundary, k - 1})
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
			} else if k != rs.lowBoundary && k != rs.highBoundary {
				finiteDifferences = append(finiteDifferences, k)
			}
			prevKey = k
		}
		sets := make([]Set, 0)
		for _, rs := range rangeDifferences {
			sets = append(sets, Set(rs))
		}
		sets = append(sets, Set(NewFromSlice(finiteDifferences)))
		return Result{sets}
	default:
		rangeDifferences := make([]RangeSet, 0)
		finiteDifferences := make([]float64, 0)

		//sorting by finite keys
		keys := make([]float64, 0)
		for k := range this.set {
			keys = append(keys, k)
		}
		sort.Float64s(keys)

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
		return Result{sets}
	}
}

//assuming the other set is the universal set
func (this FiniteSet) Complement(s Set) (Result, error) {
	switch s.(type) {
	case FiniteSet:
		fs := s.(FiniteSet)
		complements := make([]float64, 0)
		if reflect.DeepEqual(Set(this), fs) {
			return Result{}, nil
		} else {
			//checks if all keys exist in universal set
			for k, _ := range fs.set {
				if _, ok := this.set[k]; !ok {
					return Result{}, errors.New("the universal set does not include element")
				}
			}
			for k, _ := range this.set {
				if _, ok := this.set[k]; !ok {
					complements = append(complements, k)
				}
			}
		}
		sets := make([]Set, 0)
		sets = append(sets, Set(NewFromSlice(complements)))
		return Result{sets}, nil
	case RangeSet:
		rs := s.(RangeSet)
		keys := make([]float64, 0)
		for k := range this.set {
			keys = append(keys, k)
		}
		sort.Float64s(keys)
		if keys[len(keys)-1] > rs.highBoundary || keys[0] < rs.lowBoundary {
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
	default:
		keys := make([]float64, 0)
		for k := range this.set {
			keys = append(keys, k)
		}
		sort.Float64s(keys)

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

func (this FiniteSet) Remove(k float64) bool {
	if _, ok := this.set[k]; ok {
		return false
	}
	delete(this.set, k)
	return true
}

func isBetween(target, lowBoundary, highBoundary float64) bool {
	return target >= lowBoundary && target <= highBoundary
}

func isBetweenExcluding(target, lowBoundary, highBoundary float64) bool {
	return target > lowBoundary && target < highBoundary
}
