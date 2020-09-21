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

}

/*
func TestDifference(t *testing.T) {
	minusInf, regular, inf := testData()
	assert.Equal(t, "", "", "The two words should be the same.")
}

func TestIntersection(t *testing.T) {
	minusInf, regular, inf := testData()
	assert.Equal(t, "", "", "The two words should be the same.")
}
func TestComplement(t *testing.T) {
	minusInf, regular, inf := testData()
	assert.Equal(t, "", "", "The two words should be the same.")
}
*/
