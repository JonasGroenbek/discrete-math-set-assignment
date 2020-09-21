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

func fail(msg string, t *testing.T) {
	t.Log(msg)
	t.Fail()
}

func TestUnionRangeSet(t *testing.T) {
	mInf, pInf, inf := rangeSets()
	low := float64(-15)
	high := float64(123)
	regular := RangeSet{low, high}
	result := mInf.Union(regular)
	if !result.Contains(math.Inf(-1)) {
		fail("Result does not include -infinity on union", t)
	}
	result = pInf.Union(regular)
	if !result.Contains(low) || !result.Contains(high) || !result.Contains(math.Inf(1)) {
		fail("failed on union on regular and positive infinite rangesets", t)
	}
	result = inf.Union(regular)
	if !result.Contains(math.Inf(-1)) || !result.Contains(low) || !result.Contains(high) || !result.Contains(math.Inf(1)) {
		fail("failed on union on regular and infinite rangesets", t)
	}
}

func TestUnionInfiniteSet(t *testing.T) {
	mInf, pInf, inf := rangeSets()
	is := NewInfiniteSet()
	result := mInf.Union(is)
	if !result.Contains(math.Inf(-1)) || !result.Contains(math.Inf(1)) {
		fail("failed on union with minus infinity rangeset and  InfiniteSet", t)
	}
	result = pInf.Union(is)
	if !result.Contains(math.Inf(-1)) || !result.Contains(math.Inf(1)) {
		fail("failed on union on positive infinity rangeset and infinite set", t)
	}
	result = inf.Union(is)
	if !result.Contains(math.Inf(-1)) || !result.Contains(math.Inf(1)) {
		fail("failed on union on regular and infinite rangesets", t)
	}
}

func TestUnionFiniteSet(t *testing.T) {
	mInf, pInf, inf := rangeSets()
	vals := []float64{
		-234234, -1, 8923942, 98234, 1231312, 4234,
	}
	fs := NewFromSlice(vals)

	result := mInf.Union(fs)
	if !result.Contains(math.Inf(-1)) || !result.Contains(vals[3]) || !result.Contains(vals[0]) {
		fail("failed on union with minus infinity rangeset and FiniteSet", t)
	}

	result = pInf.Union(fs)
	if !result.Contains(vals[3]) || !result.Contains(vals[1]) || !result.Contains(math.Inf(0)) {
		fail("failed on union on positive infinity rangeset and FiniteSet", t)
	}

	result = inf.Union(fs)
	if !result.Contains(math.Inf(-1)) || !result.Contains(math.Inf(1)) || !result.Contains(vals[3]) || !result.Contains(vals[1]) {
		fail("failed on union on regular and FiniteSet", t)
	}

}

func TestDifferenceRangeSet(t *testing.T) {
	mInf, pInf, inf := rangeSets()
	low := float64(-1325)
	high := float64(123)
	regular := RangeSet{low, high}
	result := mInf.Difference(regular)
	if !result.Contains(math.Inf(-1)) {
		fail("Result does not include -infinity on difference", t)
	}
	result = pInf.Difference(regular)
	if !result.Contains(low) || !result.Contains(math.Inf(1)) {
		fail("failed on difference on regular and positive infinite rangesets", t)
	}
	result = inf.Difference(regular)
	if !result.Contains(math.Inf(-1)) || !result.Contains(math.Inf(1)) {
		fail("failed on difference on regular and infinite rangesets", t)
	}
}

func TestDifferenceFiniteSet(t *testing.T) {
	mInf, pInf, inf := rangeSets()
	fs := NewFromSlice([]float64{
		-500000,
		500000,
	})

	result := mInf.Difference(fs)
	if !result.Contains(math.Inf(-1)) || !result.Contains(500000) {
		fail("Result does not contain difference elements when diff between minus inf RangeSet and FiniteSet", t)
	}

	result = pInf.Difference(fs)
	t.Log(result)
	if result.Contains(math.Inf(-500000)) || !result.Contains(math.Inf(1)) {
		fail("Result does not contain difference elements when diff between plus inf RangeSet and FiniteSet", t)
	}
	result = inf.Difference(fs)
	if !result.Contains(math.Inf(-1)) || !result.Contains(math.Inf(1)) {
		fail("Result does notcontain difference elements when diff between inf RangeSet and FiniteSet", t)
	}
}

func TestIntersectionRangeSet(t *testing.T) {
	mInf, pInf, inf := rangeSets()
	low := float64(-1325)
	high := float64(123)
	regular := RangeSet{low, high}
	result := mInf.Intersection(regular)
	if result.Contains(math.Inf(-1)) {
		fail("Result does not exclude -infinity on difference", t)
	}
	result = pInf.Intersection(regular)
	if !result.Contains(high) || result.Contains(math.Inf(1)) {
		fail("failed on intersection on regular and positive infinite rangesets", t)
	}
	result = inf.Intersection(regular)
	if result.Contains(math.Inf(-1)) || result.Contains(math.Inf(1)) {
		fail("failed on intersection on regular and infinite rangesets", t)
	}
}

func TestIntersectionFiniteSet(t *testing.T) {
	mInf, pInf, inf := rangeSets()
	fs := NewFromSlice([]float64{
		math.Inf(-1),
		-500000,
		500000,
		math.Inf(1),
	})
	result := mInf.Intersection(fs)
	if !result.Contains(math.Inf(-1)) || !result.Contains(math.Inf(-500000)) || !result.Contains(500000) || !result.Contains(math.Inf(1)) {
		fail("Result does not exclude FiniteSet elements", t)
	}
	result = pInf.Intersection(fs)
	if !result.Contains(math.Inf(-1)) || !result.Contains(math.Inf(-500000)) || !result.Contains(500000) || !result.Contains(math.Inf(1)) {
		fail("Result does not exclude FiniteSet elements when intersecting with ", t)
	}
	result = inf.Intersection(fs)
	if !result.Contains(math.Inf(-1)) || !result.Contains(math.Inf(-500000)) || !result.Contains(500000) || !result.Contains(math.Inf(1)) {
		fail("Result does not exclude FiniteSet elements", t)
	}
}

func TestComplementRangeSet(t *testing.T) {
	mInf, pInf, inf := rangeSets()
	low := float64(100)
	high := float64(1000)
	regular := RangeSet{low, high}
	result, err := mInf.Complement(inf)
	if result.Contains(math.Inf(-1)) || err != nil {
		fail("Complement include element from original set", t)
	}
	result, err = pInf.Complement(inf)
	if result.Contains(0) || result.Contains(math.Inf(1)) || err != nil {
		fail("Complement include its original elements", t)
	}
	result, err = regular.Complement(inf)
	if result.Contains(500) || !result.Contains(1001) || !result.Contains(99) || err != nil {
		fail("failed on complement on regular and infinite rangesets", t)
	}

	result, err = mInf.Complement(regular)
	if err == nil {
		fail("expected error on wrong universal set", t)
	}
}
