package assignment3

import (
	"errors"
	"math"
	"sort"
)

type RangeSet struct {
	lowBoundary  float64
	highBoundary float64
}

func (this RangeSet) Union(s Set) CompositeSet {
	switch s.(type) {
	case FiniteSet:
		fs := s.(FiniteSet)
		for k, _ := range fs.set {
			if isBetween(k, this.lowBoundary, this.highBoundary) {
				delete(fs.set, k)
			}
		}
		return CompositeSet{[]Set{
			fs,
			this,
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
		return CompositeSet{[]Set{
			RangeSet{low, high},
		}}
	default:
		return CompositeSet{[]Set{
			NewInfiniteSet(),
		}}
	}
}

func (this RangeSet) Intersection(s Set) CompositeSet {
	switch s.(type) {
	case FiniteSet:
		fs := s.(FiniteSet)

		keys := make([]float64, 0)
		for k := range fs.set {
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
		if this.highBoundary > high {
			high = this.highBoundary
		}
		if this.lowBoundary < low {
			low = this.lowBoundary
		}
		return CompositeSet{[]Set{
			RangeSet{low, high},
		}}
	case RangeSet:
		rs := s.(RangeSet)
		var low float64
		var high float64
		//rs contains
		if rs.lowBoundary <= this.lowBoundary && rs.highBoundary >= this.highBoundary {
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
				//this lowest
			} else if this.lowBoundary < rs.lowBoundary {
				low = rs.lowBoundary
				high = this.highBoundary
				//rs lowest
			} else if rs.lowBoundary < this.lowBoundary {
				low = this.lowBoundary
				high = rs.highBoundary
			} else {
				return CompositeSet{}
			}
		}
		return CompositeSet{
			[]Set{
				RangeSet{
					low,
					high,
				},
			},
		}
	default:
		return CompositeSet{
			[]Set{
				NewInfiniteSet(),
			},
		}
	}
}

func (this RangeSet) Difference(s Set) CompositeSet {
	switch s.(type) {
	case FiniteSet:
		fs := s.(FiniteSet)
		//creating range differences and specific differences
		rangeDiffs := make([]RangeSet, 0)
		singleDiffs := make([]float64, 0)

		//sorting by finite keys
		keys := make([]float64, 0)
		for k := range fs.set {
			keys = append(keys, k)
		}

		sort.Float64s(keys)

		var prevKey float64
		for i, k := range keys {
			if isBetween(k, this.lowBoundary, this.highBoundary) {
				//no previous key
				if i == 0 {
					rangeDiffs = append(rangeDiffs, RangeSet{this.lowBoundary, k - 1})
				} else {
					//check if the previous key is 1 apart
					if prevKey != k-1 {
						//check if there is a gap of 1 between previous key
						if prevKey == k-2 {
							singleDiffs = append(singleDiffs, k-1)
							//append range since gap is larger than 2
						} else {
							if prevKey >= this.lowBoundary {
								rangeDiffs = append(rangeDiffs, RangeSet{prevKey + 1, k - 1})
							} else {
								rangeDiffs = append(rangeDiffs, RangeSet{this.lowBoundary, k - 1})
							}
						}
					}
				}
				//check if key does not equal boundaries, if that is the case then they both include and no difference
			} else if k > this.highBoundary {
				singleDiffs = append(singleDiffs, k)
				if this.highBoundary > prevKey {
					rangeDiffs = append(rangeDiffs, RangeSet{prevKey + 1, this.highBoundary})
				}
			} else if k < this.lowBoundary {
				singleDiffs = append(singleDiffs, k)
			}
			prevKey = k
		}
		if prevKey < this.highBoundary {
			rangeDiffs = append(rangeDiffs, RangeSet{prevKey + 1, this.highBoundary})
		}
		sets := make([]Set, 0)
		for _, rs := range rangeDiffs {
			sets = append(sets, Set(rs))
		}
		sets = append(sets, Set(NewFromSlice(singleDiffs)))
		return CompositeSet{sets}
	case RangeSet:
		rs := s.(RangeSet)
		//rs contains
		if rs.lowBoundary <= this.lowBoundary && rs.highBoundary >= this.highBoundary {
			return CompositeSet{[]Set{
				RangeSet{this.lowBoundary, rs.lowBoundary},
				RangeSet{this.highBoundary, rs.highBoundary},
			}}
			//this contains
		} else if this.lowBoundary <= rs.lowBoundary && this.highBoundary >= rs.highBoundary {
			return CompositeSet{[]Set{
				RangeSet{this.lowBoundary, rs.lowBoundary},
				RangeSet{rs.highBoundary, this.highBoundary},
			}}
			//they cross
		} else if isBetween(this.lowBoundary, rs.lowBoundary, rs.highBoundary) || isBetween(this.highBoundary, rs.lowBoundary, rs.highBoundary) {
			//equal low
			if this.lowBoundary == rs.lowBoundary {
				if this.highBoundary <= rs.highBoundary {
					return CompositeSet{[]Set{
						RangeSet{this.highBoundary, rs.highBoundary},
					}}
				} else {
					return CompositeSet{[]Set{
						RangeSet{rs.highBoundary, rs.highBoundary},
					}}
				}
				//equal high
			} else if this.highBoundary == rs.highBoundary {
				if this.lowBoundary <= rs.lowBoundary {
					return CompositeSet{[]Set{
						RangeSet{this.lowBoundary, rs.lowBoundary},
					}}
				} else {
					return CompositeSet{[]Set{
						RangeSet{rs.lowBoundary, this.lowBoundary},
					}}
				}
				//this lowest
			} else if this.lowBoundary < rs.lowBoundary {
				return CompositeSet{[]Set{
					RangeSet{this.lowBoundary, rs.lowBoundary},
					RangeSet{this.highBoundary, rs.highBoundary},
				}}
				//rs lowest
			} else if rs.lowBoundary < this.lowBoundary {
				return CompositeSet{[]Set{
					RangeSet{rs.lowBoundary, this.lowBoundary},
					RangeSet{rs.highBoundary, this.highBoundary},
				}}
			} else {
				return CompositeSet{}
			}
		}
	default:
		sets := make([]Set, 0)
		if this.lowBoundary == math.Inf(-1) && this.highBoundary == math.Inf(1) {
			return CompositeSet{}
		}
		if this.lowBoundary != math.Inf(-1) {
			sets = append(sets, RangeSet{math.Inf(-1), this.lowBoundary - 1})
		}
		if this.highBoundary != math.Inf(1) {
			sets = append(sets, RangeSet{this.highBoundary + 1, math.Inf(1)})
		}
		return CompositeSet{sets}
	}
	return CompositeSet{}
}

func (this RangeSet) Complement(s Set) (CompositeSet, error) {
	switch s.(type) {
	case FiniteSet:
		fs := s.(FiniteSet)
		keys := make([]float64, 0)
		for k := range fs.set {
			keys = append(keys, k)
		}
		sort.Float64s(keys)
		if keys[len(keys)-1] < this.highBoundary || keys[0] > this.lowBoundary {
			return CompositeSet{}, errors.New("the universal set does not include element")
		} else {
			singleDiffs := make([]float64, 0)
			for _, k := range keys {
				if !isBetween(k, this.lowBoundary, this.highBoundary) {
					singleDiffs = append(singleDiffs, k)
				}
			}
			return CompositeSet{[]Set{
				NewFromSlice(singleDiffs),
			}}, nil
		}
	case RangeSet:
		rs := s.(RangeSet)
		if this.lowBoundary < rs.lowBoundary || this.highBoundary > rs.highBoundary {
			return CompositeSet{}, errors.New("the universal set does not include element")
		} else {
			sets := make([]Set, 0)
			if this.highBoundary != rs.highBoundary {
				sets = append(sets, RangeSet{this.highBoundary + 1, rs.highBoundary})
			}
			if this.lowBoundary != rs.lowBoundary {
				sets = append(sets, RangeSet{rs.lowBoundary, this.lowBoundary - 1})
			}
			return CompositeSet{sets}, nil
		}
	default:
		return CompositeSet{}, nil
	}
}
