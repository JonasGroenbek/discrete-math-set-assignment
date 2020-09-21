package assignment3

import (
	"math"
	"testing"
)

func rangeSets() (RangeSet, RangeSet, RangeSet) {
	mInf := RangeSet{math.Inf(-1), 15}
	pInf := RangeSet{-13, math.Inf(1)}
	inf := RangeSet{math.Inf(-1), math.Inf(1)}
	return mInf, pInf, inf
}

func TestRangeSetUnionRangeSet(t *testing.T) {
	mInf, pInf, inf := rangeSets()
	low := float64(-15)
	high := float64(123)
	regular := RangeSet{low, high}
	cs := mInf.Union(regular)
	if !cs.Contains(math.Inf(-1)) {
		fail("cs does not include -infinity on union", t)
	}
	cs = pInf.Union(regular)
	if !cs.Contains(low) || !cs.Contains(high) || !cs.Contains(math.Inf(1)) {
		fail("failed on union on regular and positive infinite rangesets", t)
	}
	cs = inf.Union(regular)
	if !cs.Contains(math.Inf(-1)) || !cs.Contains(low) || !cs.Contains(high) || !cs.Contains(math.Inf(1)) {
		fail("failed on union on regular and infinite rangesets", t)
	}
}

func TestRangeSetUnionInfiniteSet(t *testing.T) {
	mInf, pInf, inf := rangeSets()
	is := NewInfiniteSet()
	cs := mInf.Union(is)
	if !cs.Contains(math.Inf(-1)) || !cs.Contains(math.Inf(1)) {
		fail("failed on union with minus infinity rangeset and  InfiniteSet", t)
	}
	cs = pInf.Union(is)
	if !cs.Contains(math.Inf(-1)) || !cs.Contains(math.Inf(1)) {
		fail("failed on union on positive infinity rangeset and infinite set", t)
	}
	cs = inf.Union(is)
	if !cs.Contains(math.Inf(-1)) || !cs.Contains(math.Inf(1)) {
		fail("failed on union on regular and infinite rangesets", t)
	}
}

func TestRangeSetUnionFiniteSet(t *testing.T) {
	mInf, pInf, inf := rangeSets()
	vals := []float64{
		-234234, -1, 8923942, 98234, 1231312, 4234,
	}
	fs := NewFromSlice(vals)

	cs := mInf.Union(fs)
	if !cs.Contains(math.Inf(-1)) || !cs.Contains(vals[3]) || !cs.Contains(vals[0]) {
		fail("failed on union with minus infinity rangeset and FiniteSet", t)
	}

	cs = pInf.Union(fs)
	if !cs.Contains(vals[3]) || !cs.Contains(vals[1]) || !cs.Contains(math.Inf(0)) {
		fail("failed on union on positive infinity rangeset and FiniteSet", t)
	}

	cs = inf.Union(fs)
	if !cs.Contains(math.Inf(-1)) || !cs.Contains(math.Inf(1)) || !cs.Contains(vals[3]) || !cs.Contains(vals[1]) {
		fail("failed on union on regular and FiniteSet", t)
	}

}

func TestRangeSetDifferenceRangeSet(t *testing.T) {
	mInf, pInf, inf := rangeSets()
	low := float64(-1325)
	high := float64(123)
	regular := RangeSet{low, high}
	cs := mInf.Difference(regular)
	if !cs.Contains(math.Inf(-1)) {
		fail("cs does not include -infinity on difference", t)
	}
	cs = pInf.Difference(regular)
	if !cs.Contains(low) || !cs.Contains(math.Inf(1)) {
		fail("failed on difference on regular and positive infinite rangesets", t)
	}
	cs = inf.Difference(regular)
	if !cs.Contains(math.Inf(-1)) || !cs.Contains(math.Inf(1)) {
		fail("failed on difference on regular and infinite rangesets", t)
	}
}

func TestRangeSetDifferenceFiniteSet(t *testing.T) {
	mInf, pInf, inf := rangeSets()
	fs := NewFromSlice([]float64{
		-500000,
		500000,
	})

	cs := mInf.Difference(fs)
	if !cs.Contains(math.Inf(-1)) || !cs.Contains(500000) {
		fail("cs does not contain difference elements when diff between minus inf RangeSet and FiniteSet", t)
	}

	cs = pInf.Difference(fs)
	t.Log(cs)
	if cs.Contains(math.Inf(-500000)) || !cs.Contains(math.Inf(1)) {
		fail("cs does not contain difference elements when diff between plus inf RangeSet and FiniteSet", t)
	}
	cs = inf.Difference(fs)
	if !cs.Contains(math.Inf(-1)) || !cs.Contains(math.Inf(1)) {
		fail("cs does notcontain difference elements when diff between inf RangeSet and FiniteSet", t)
	}
}

func TestRangeSetIntersectionRangeSet(t *testing.T) {
	mInf, pInf, inf := rangeSets()
	low := float64(-1325)
	high := float64(123)
	regular := RangeSet{low, high}
	cs := mInf.Intersection(regular)
	if cs.Contains(math.Inf(-1)) {
		fail("cs does not exclude -infinity on difference", t)
	}
	cs = pInf.Intersection(regular)
	if !cs.Contains(high) || cs.Contains(math.Inf(1)) {
		fail("failed on intersection on regular and positive infinite rangesets", t)
	}
	cs = inf.Intersection(regular)
	if cs.Contains(math.Inf(-1)) || cs.Contains(math.Inf(1)) {
		fail("failed on intersection on regular and infinite rangesets", t)
	}
}

func TestRangeSetIntersectionFiniteSet(t *testing.T) {
	mInf, pInf, inf := rangeSets()
	fs := NewFromSlice([]float64{
		math.Inf(-1),
		-500000,
		500000,
		math.Inf(1),
	})
	cs := mInf.Intersection(fs)
	if !cs.Contains(math.Inf(-1)) || !cs.Contains(math.Inf(-500000)) || !cs.Contains(500000) || !cs.Contains(math.Inf(1)) {
		fail("cs does not exclude FiniteSet elements", t)
	}
	cs = pInf.Intersection(fs)
	if !cs.Contains(math.Inf(-1)) || !cs.Contains(math.Inf(-500000)) || !cs.Contains(500000) || !cs.Contains(math.Inf(1)) {
		fail("cs does not exclude FiniteSet elements when intersecting with ", t)
	}
	cs = inf.Intersection(fs)
	if !cs.Contains(math.Inf(-1)) || !cs.Contains(math.Inf(-500000)) || !cs.Contains(500000) || !cs.Contains(math.Inf(1)) {
		fail("cs does not exclude FiniteSet elements", t)
	}
}

func TestRangeSetComplementRangeSet(t *testing.T) {
	mInf, pInf, inf := rangeSets()
	low := float64(100)
	high := float64(1000)
	regular := RangeSet{low, high}
	cs, err := mInf.Complement(inf)
	if cs.Contains(math.Inf(-1)) || err != nil {
		fail("Complement include element from original set", t)
	}
	cs, err = pInf.Complement(inf)
	if cs.Contains(0) || cs.Contains(math.Inf(1)) || err != nil {
		fail("Complement include its original elements", t)
	}
	cs, err = regular.Complement(inf)
	if cs.Contains(500) || !cs.Contains(1001) || !cs.Contains(99) || err != nil {
		fail("failed on complement on regular and infinite rangesets", t)
	}

	cs, err = mInf.Complement(regular)
	if err == nil {
		fail("expected error on wrong universal set", t)
	}
}

func TestRangeSetComplementFiniteSet(t *testing.T) {

	rs := RangeSet{2, 4}
	rsError := RangeSet{0, 9}
	fs := NewFromSlice([]float64{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	})

	cs, err := rs.Complement(fs)
	t.Log(cs)
	if !cs.Contains(1) || !cs.Contains(5) || !cs.Contains(9) {
		fail("Does not include the complement elements when RangeSet complements Finite universal set", t)
	}
	_, err = rsError.Complement(fs)
	if err == nil {
		fail("should throw error when set contains elements not inside universal set", t)
	}
}
