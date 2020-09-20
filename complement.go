package main

type ComplementSet struct {
	sets []Set
}

func EmptyComplementSet() ComplementSet {
	return ComplementSet{make([]Set, 0)}
}
