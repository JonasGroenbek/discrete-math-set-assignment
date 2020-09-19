package main

import (
	"bytes"
	"fmt"
)

func main() {
	mySet := New()
	mySet.Add(1, 2)
	b := new(bytes.Buffer)
	for key, value := range mySet.set {
		fmt.Fprintf(b, "%d=\"%d\"\n", key, value)
	}
	fmt.Println(b.String())

	mySet.ChangeI(123)
	fmt.Println(mySet.i)
}
