package assignment3

type Result struct {
	sets []Set
}

func (this Result) Contains(target float64) bool {
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
