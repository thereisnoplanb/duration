package duration

// Gets the result of rounding current Duration value to the nearest multiple of m.
//
// The rounding behavior for halfway values is to round away from zero.
// If the result exceeds the maximum (or minimum) value that can be stored in a Duration, Round returns the maximum (or minimum) duration.
// If m <= Zero, Round returns current Duration unchanged.
//
// Parameters:
//
//	m Duration - Round seed.
//
// Returns:
//
//	Duration - result of rounding current Duration value to the nearest multiple of m.
func (duration Duration) Round(m Duration) Duration {
	if m <= 0 {
		return duration
	}
	r := duration % m
	if duration < 0 {
		r = -r
		if lessThanHalf(r, m) {
			return duration + r
		}
		if d1 := duration - m + r; d1 < duration {
			return d1
		}
		return Minimum // overflow
	}
	if lessThanHalf(r, m) {
		return duration - r
	}
	if d1 := duration + m - r; d1 > duration {
		return d1
	}
	return Maximum // overflow
}

func lessThanHalf(x, y Duration) bool {
	return uint64(x)+uint64(x) < uint64(y)
}
