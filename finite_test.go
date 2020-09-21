package assignment3

import (
	"testing"
)

func finiteSets() (FiniteSet, FiniteSet) {
	ifs := NewFromSlice([]float64{
		inf(),
		10000,
		-10000,
		nInf(),
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

	low := float64(-100000)
	high := float64(100)
	rs := RangeSet{low, high}

	cs := ifs.Union(rs)
	if !cs.ContainsMultiple([]float64{nInf(), low, high, inf()}) {
		fail("cs does not include -infinity on union", t)
	}
	cs = fs.Union(rs)
	if !cs.ContainsMultiple([]float64{low, high, 10000}) {
		fail("failed on union on regular and positive infinite rangesets", t)
	}
}

func TestUnionFiniteSet(t *testing.T) {
	ifs, fs := finiteSets()
	vals := []float64{
		-234234, -1, 8923942, 98234, 1231312, 4234,
	}
	fs2 := NewFromSlice(vals)

	keys := make([]float64, 0, len(fs.set))
	for k := range fs.set {
		keys = append(keys, k)
	}

	cs := ifs.Union(fs2)
	if !cs.ContainsMultiple(append(vals, nInf(), inf())) {
		fail("failed on union with minus infinity rangeset and FiniteSet", t)
	}

	cs = fs.Union(fs2)
	if !cs.ContainsMultiple(append(vals, keys...)) {
		fail("failed on union on positive infinity rangeset and FiniteSet", t)
	}

}

/*
func TestDifferenceRangeSet(t *testing.T) {
	mInf, pInf, inf := rangeSets()
	low := float64(-1325)
	high := float64(123)
	regular := RangeSet{low, high}
	cs := mInf.Difference(regular)
	if !cs.Contains(nInf()) {
		fail("cs does not include -infinity on difference", t)
	}
	cs = pInf.Difference(regular)
	if !cs.Contains(low) || !cs.Contains(inf()) {
		fail("failed on difference on regular and positive infinite rangesets", t)
	}
	cs = inf.Difference(regular)
	if !cs.Contains(nInf()) || !cs.Contains(inf()) {
		fail("failed on difference on regular and infinite rangesets", t)
	}
}

func TestDifferenceFiniteSet(t *testing.T) {
	mInf, pInf, inf := rangeSets()
	fs := NewFromSlice([]float64{
		-500000,
		500000,
	})

	cs := mInf.Difference(fs)
	if !cs.Contains(nInf()) || !cs.Contains(500000) {
		fail("cs does not contain difference elements when diff between minus inf RangeSet and FiniteSet", t)
	}

	cs = pInf.Difference(fs)
	t.Log(cs)
	if cs.Contains(math.Inf(-500000)) || !cs.Contains(inf()) {
		fail("cs does not contain difference elements when diff between plus inf RangeSet and FiniteSet", t)
	}
	cs = inf.Difference(fs)
	if !cs.Contains(nInf()) || !cs.Contains(inf()) {
		fail("cs does notcontain difference elements when diff between inf RangeSet and FiniteSet", t)
	}
}

func TestIntersectionRangeSet(t *testing.T) {
	mInf, pInf, inf := rangeSets()
	low := float64(-1325)
	high := float64(123)
	regular := RangeSet{low, high}
	cs := mInf.Intersection(regular)
	if cs.Contains(nInf()) {
		fail("cs does not exclude -infinity on difference", t)
	}
	cs = pInf.Intersection(regular)
	if !cs.Contains(high) || cs.Contains(inf()) {
		fail("failed on intersection on regular and positive infinite rangesets", t)
	}
	cs = inf.Intersection(regular)
	if cs.Contains(nInf()) || cs.Contains(inf()) {
		fail("failed on intersection on regular and infinite rangesets", t)
	}
}

func TestIntersectionFiniteSet(t *testing.T) {
	mInf, pInf, inf := rangeSets()
	fs := NewFromSlice([]float64{
		nInf(),
		-500000,
		500000,
		inf(),
	})
	cs := mInf.Intersection(fs)
	if !cs.Contains(nInf()) || !cs.Contains(math.Inf(-500000)) || !cs.Contains(500000) || !cs.Contains(inf()) {
		fail("cs does not exclude FiniteSet elements", t)
	}
	cs = pInf.Intersection(fs)
	if !cs.Contains(nInf()) || !cs.Contains(math.Inf(-500000)) || !cs.Contains(500000) || !cs.Contains(inf()) {
		fail("cs does not exclude FiniteSet elements when intersecting with ", t)
	}
	cs = inf.Intersection(fs)
	if !cs.Contains(nInf()) || !cs.Contains(math.Inf(-500000)) || !cs.Contains(500000) || !cs.Contains(inf()) {
		fail("cs does not exclude FiniteSet elements", t)
	}
}

func TestComplementRangeSet(t *testing.T) {
	mInf, pInf, inf := rangeSets()
	low := float64(100)
	high := float64(1000)
	regular := RangeSet{low, high}
	cs, err := mInf.Complement(inf)
	if cs.Contains(nInf()) || err != nil {
		fail("Complement include element from original set", t)
	}
	cs, err = pInf.Complement(inf)
	if cs.Contains(0) || cs.Contains(inf()) || err != nil {
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

func TestComplementFiniteSet(t *testing.T) {

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

*/
