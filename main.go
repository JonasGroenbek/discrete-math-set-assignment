package main

import (
	"fmt"
	"math"
)

func main() {
	rs1 := RangeSet{math.Inf(-1), math.Inf(1)}
	rs2 := RangeSet{15, 32}

	fmt.Println(rs1.Union(rs2))
}
