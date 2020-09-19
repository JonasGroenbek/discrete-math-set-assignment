package main

import (
	"fmt"
)

func main() {
	mySet := New()
	mySet.Add(1)
	mySet.Add(2)
	mySet.Add(5)
	fmt.Println(mySet)
}
