package assignment3

//Set to be implemented by all sets
type Set interface {
	Union(s Set) Result
	Intersection(s Set) Result
	Difference(s Set) Result
	Complement(s Set) (Result, error)
}
