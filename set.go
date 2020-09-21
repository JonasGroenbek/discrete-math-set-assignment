package assignment3

type Set interface {
	Union(s Set) CompositeSet
	Intersection(s Set) CompositeSet
	Difference(s Set) CompositeSet
	Complement(s Set) (CompositeSet, error)
}
