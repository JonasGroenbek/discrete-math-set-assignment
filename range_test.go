package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func testData() (RangeSet, RangeSet, RangeSet) {
	minusInf := RangeSet{math.Inf(-1), 15}
	regular := RangeSet{-32, 234}
	inf := RangeSet{-13, math.Inf(1)}
	return minusInf, regular, inf
}

func TestUnion(t *testing.T) {
	minusInf, regular, inf := testData()
	fmt.Println("testing union with rangeset")
	testRs := RangeSet{-15, 123}
	fmt.Println("union minusIf", minusInf.Union(testRs))
	fmt.Println("union regular", regular.Union(testRs))
	inf.Union(testRs)
	fmt.Println("testing union with infiniteset")
	fmt.Println("testing union with rangeset")
	assert.Equal(t, "", "", "The two words should be the same.")
}
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
