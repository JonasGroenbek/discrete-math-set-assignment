package assignment3

import (
	"math"
	"testing"
)

func infs() (float64, float64) {
	inf := math.Inf(+1)
	mInf := math.Inf(-1)
	return inf, mInf
}
func finiteSets() (FiniteSet, FiniteSet) {
	ifs := NewFromSlice([]float64{
		math.Inf(1),
		10000,
		-10000,
		math.Inf(-1),
	})
	fs := NewFromSlice([]float64{
		10000,
		1000,
		-10000,
		-1000,
	})
	return ifs, fs
}

func fail(msg string, t *testing.T) {
	t.Log(msg)
	t.Fail()
}

func TestInfiniteSetUnionRangeSet(t *testing.T) {
	ifs, fs := finiteSets()
	inf, mInf := infs()

	low := float64(-100000)
	high := float64(100)
	rs := RangeSet{low, high}

	res := ifs.Union(rs)
	if !res.ContainsMultiple([]float64{mInf, low, high, inf}) {
		fail("Result does not include -infinity on union", t)
	}
	res = fs.Union(rs)
	if !res.ContainsMultiple([]float64{low, high, 10000}) {
		fail("failed on union on regular and positive infinite rangesets", t)
	}
}

func TestUnionFiniteSet(t *testing.T) {
	ifs, fs := finiteSets()
	inf, mInf := infs()
	vals := []float64{
		-234234, -1, 8923942, 98234, 1231312, 4234,
	}
	fs2 := NewFromSlice(vals)

	keys := make([]float64, 0, len(fs.set))
	for k := range fs.set {
		keys = append(keys, k)
	}

	res := ifs.Union(fs2)
	if !res.ContainsMultiple(append(vals, mInf, inf)) {
		fail("failed on union with minus infinity rangeset and FiniteSet", t)
	}

	res = fs.Union(fs2)
	if !res.ContainsMultiple(append(vals, keys...)) {
		fail("failed on union on positive infinity rangeset and FiniteSet", t)
	}

}

/*
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

func TestComplementFiniteSet(t *testing.T) {

	rs := RangeSet{2, 4}
	rsError := RangeSet{0, 9}
	fs := NewFromSlice([]float64{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	})

	result, err := rs.Complement(fs)
	t.Log(result)
	if !result.Contains(1) || !result.Contains(5) || !result.Contains(9) {
		fail("Does not include the complement elements when RangeSet complements Finite universal set", t)
	}
	_, err = rsError.Complement(fs)
	if err == nil {
		fail("should throw error when set contains elements not inside universal set", t)
	}
}

*/
