package main

func main() {
	fs := New()
	fs.Add(1)
	fs.Add(3)
	fs.Add(2)
	s := Set(fs)
	s.CompareTo(Set(InfiniteSet{}))
}
