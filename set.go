package main

//Set to be implemented by all sets
type Set interface {
	Member(f float64) bool
	Union(s Set) UnionSet
	Intersection(s Set) IntersectionSet
	Difference(s Set) DifferenceSet
	Complement(s Set) ComplementSet
	CompareTo(s Set) int
}
