package main

import (
	"fmt"
	"math"
)

func main() {
	fs := NewFromSlice([]float64{
		1,
		2,
		3,
		4,
		8,
		10,
	})
	rs := RangeSet{math.Inf(-1), math.Inf(1)}
	fmt.Println(fs.Difference(rs))
}
