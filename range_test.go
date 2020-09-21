package assignment3

import (
	"testing"
)

func rangeSets() (RangeSet, RangeSet, RangeSet) {
	nInfRs := RangeSet{nInf(), 15}
	pInfRs := RangeSet{-13, inf()}
	infRs := RangeSet{nInf(), inf()}
	return nInfRs, pInfRs, infRs
}

func TestRangeSetUnionRangeSet(t *testing.T) {
	nInfRs, pInfRs, infRs := rangeSets()
	low := float64(-15)
	high := float64(123)
	regular := RangeSet{low, high}
	cs := nInfRs.Union(regular)
	if !cs.Contains(nInf()) {
		fail("cs does not include -infinity on union", t)
	}
	cs = pInfRs.Union(regular)
	if !cs.Contains(low) || !cs.Contains(high) || !cs.Contains(inf()) {
		fail("failed on union on regular and positive infinite rangesets", t)
	}
	cs = infRs.Union(regular)
	if !cs.Contains(nInf()) || !cs.Contains(low) || !cs.Contains(high) || !cs.Contains(inf()) {
		fail("failed on union on regular and infinite rangesets", t)
	}
}

func TestRangeSetUnionInfiniteSet(t *testing.T) {
	nInfRs, pInfRs, infRs := rangeSets()
	is := NewInfiniteSet()
	cs := nInfRs.Union(is)
	if !cs.Contains(nInf()) || !cs.Contains(inf()) {
		fail("failed on union with minus infinity rangeset and  InfiniteSet", t)
	}
	cs = pInfRs.Union(is)
	if !cs.Contains(nInf()) || !cs.Contains(inf()) {
		fail("failed on union on positive infinity rangeset and infinite set", t)
	}
	cs = infRs.Union(is)
	if !cs.Contains(nInf()) || !cs.Contains(inf()) {
		fail("failed on union on regular and infinite rangesets", t)
	}
}

func TestRangeSetUnionFiniteSet(t *testing.T) {
	nInfRs, pInfRs, infRs := rangeSets()
	vals := []float64{
		-234234, -1, 8923942, 98234, 1231312, 4234,
	}
	fs := NewFromSlice(vals)

	cs := nInfRs.Union(fs)
	if !cs.Contains(nInf()) || !cs.Contains(vals[3]) || !cs.Contains(vals[0]) {
		fail("failed on union with minus infinity rangeset and FiniteSet", t)
	}

	cs = pInfRs.Union(fs)
	if !cs.Contains(vals[3]) || !cs.Contains(vals[1]) || !cs.Contains(inf()) {
		fail("failed on union on positive infinity rangeset and FiniteSet", t)
	}

	cs = infRs.Union(fs)
	if !cs.Contains(nInf()) || !cs.Contains(inf()) || !cs.Contains(vals[3]) || !cs.Contains(vals[1]) {
		fail("failed on union on regular and FiniteSet", t)
	}

}

func TestRangeSetDifferenceRangeSet(t *testing.T) {
	nInfRs, pInfRs, infRs := rangeSets()
	low := float64(-1325)
	high := float64(123)
	regular := RangeSet{low, high}
	cs := nInfRs.Difference(regular)
	if !cs.Contains(nInf()) {
		fail("cs does not include -infinity on difference", t)
	}
	cs = pInfRs.Difference(regular)
	if !cs.Contains(low) || !cs.Contains(inf()) {
		fail("failed on difference on regular and positive infinite rangesets", t)
	}
	cs = infRs.Difference(regular)
	if !cs.Contains(nInf()) || !cs.Contains(inf()) {
		fail("failed on difference on regular and infinite rangesets", t)
	}
}

func TestRangeSetDifferenceFiniteSet(t *testing.T) {
	nInfRs, pInfRs, infRs := rangeSets()
	fs := NewFromSlice([]float64{
		-500000,
		500000,
	})

	cs := nInfRs.Difference(fs)
	if !cs.Contains(nInf()) || !cs.Contains(500000) {
		fail("cs does not contain difference elements when diff between minus inf RangeSet and FiniteSet", t)
	}

	cs = pInfRs.Difference(fs)
	if cs.Contains(-500000) || !cs.Contains(inf()) {
		fail("cs does not contain difference elements when diff between plus inf RangeSet and FiniteSet", t)
	}
	cs = infRs.Difference(fs)
	if !cs.Contains(nInf()) || !cs.Contains(inf()) {
		fail("cs does notcontain difference elements when diff between inf RangeSet and FiniteSet", t)
	}
}

func TestRangeSetIntersectionRangeSet(t *testing.T) {
	nInfRs, pInfRs, infRs := rangeSets()
	low := float64(-1325)
	high := float64(123)
	regular := RangeSet{low, high}
	cs := nInfRs.Intersection(regular)
	if cs.Contains(nInf()) {
		fail("cs does not exclude -infinity on difference", t)
	}
	cs = pInfRs.Intersection(regular)
	if !cs.Contains(high) || cs.Contains(inf()) {
		fail("failed on intersection on regular and positive infinite rangesets", t)
	}
	cs = infRs.Intersection(regular)
	if cs.Contains(nInf()) || cs.Contains(inf()) {
		fail("failed on intersection on regular and infinite rangesets", t)
	}
}

func TestRangeSetIntersectionFiniteSet(t *testing.T) {
	nInfRs, pInfRs, infRs := rangeSets()
	fs := NewFromSlice([]float64{
		nInf(),
		-500000,
		500000,
		inf(),
	})
	cs := nInfRs.Intersection(fs)
	if !cs.Contains(nInf()) || !cs.Contains(-500000) || !cs.Contains(500000) || !cs.Contains(inf()) {
		fail("cs does not exclude FiniteSet elements", t)
	}
	cs = pInfRs.Intersection(fs)
	if !cs.Contains(nInf()) || !cs.Contains(-500000) || !cs.Contains(500000) || !cs.Contains(inf()) {
		fail("cs does not exclude FiniteSet elements when intersecting with ", t)
	}
	cs = infRs.Intersection(fs)
	if !cs.Contains(nInf()) || !cs.Contains(-500000) || !cs.Contains(500000) || !cs.Contains(inf()) {
		fail("cs does not exclude FiniteSet elements", t)
	}
}

func TestRangeSetComplementRangeSet(t *testing.T) {
	nInfRs, pInfRs, infRs := rangeSets()
	low := float64(100)
	high := float64(1000)
	regular := RangeSet{low, high}
	cs, err := nInfRs.Complement(infRs)
	if cs.Contains(nInf()) || err != nil {
		fail("Complement include element from original set", t)
	}
	cs, err = pInfRs.Complement(infRs)
	if cs.Contains(0) || cs.Contains(inf()) || err != nil {
		fail("Complement include its original elements", t)
	}
	cs, err = regular.Complement(infRs)
	if cs.Contains(500) || !cs.Contains(1001) || !cs.Contains(99) || err != nil {
		fail("failed on complement on regular and infinite rangesets", t)
	}

	cs, err = nInfRs.Complement(regular)
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
	if !cs.Contains(1) || !cs.Contains(5) || !cs.Contains(9) {
		fail("Does not include the complement elements when RangeSet complements Finite universal set", t)
	}
	_, err = rsError.Complement(fs)
	if err == nil {
		fail("should throw error when set contains elements not inside universal set", t)
	}
}
