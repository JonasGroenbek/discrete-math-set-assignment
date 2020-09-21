package assignment3

import (
	"math"
)

func inf() float64 {
	return math.Inf(1)
}

func nInf() float64 {
	return math.Inf(-1)
}

func isBetween(target, lowBoundary, highBoundary float64) bool {
	return target >= lowBoundary && target <= highBoundary
}

func isBetweenExcluding(target, lowBoundary, highBoundary float64) bool {
	return target > lowBoundary && target < highBoundary
}
