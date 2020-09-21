package assignment3

type CompositeSet struct {
	sets []Set
}

func (this CompositeSet) Contains(target float64) bool {
	for _, set := range this.sets {
		switch set.(type) {
		case infiniteSet:
			return true
		case RangeSet:
			rs := set.(RangeSet)
			if target >= rs.lowBoundary && target <= rs.highBoundary {
				return true
			}
		case FiniteSet:
			fs := set.(FiniteSet)
			for el := range fs.set {
				if el == target {
					return true
				}
			}
		}
	}
	return false
}

func (this CompositeSet) ContainsMultiple(targets []float64) bool {
	for _, t := range targets {
		if !this.Contains(t) {
			return false
		}
		return true
	}
	//slice empty
	return true
}
