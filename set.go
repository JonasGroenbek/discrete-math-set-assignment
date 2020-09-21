package assignment3

type Set interface {
	Union(s Set) Result
	Intersection(s Set) Result
	Difference(s Set) Result
	Complement(s Set) (Result, error)
}
