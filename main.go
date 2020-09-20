package main

import (
	"fmt"
)

func main() {
	fs := NewFromSlice([]float64{
		1,
		2,
		3,
		4,
		8,
		10,
		12,
	})
	is := InfiniteSet{}
	//rs := RangeSet{math.Inf(-1), math.Inf(1)}
	fmt.Println(fs.Difference(is))
}
