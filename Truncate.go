package duration

// Gets the result of rounding current Duration value toward zero to a multiple of m.
//
// If m <= Zero, Round returns current Duration unchanged.
//
// Parameters:
//
//	m Duration - Round seed.
//
// Returns:
//
//	Duration - result of rounding current Duration value toward zero to a multiple of m.
func (duration Duration) Truncate(m Duration) Duration {
	if m <= 0 {
		return duration
	}
	return duration - duration%m
}
