package duration

// Returns a new Duration object whose value is the absolute value of this Duration instance.
// When current duration is equal to Minimum returned abolute value is Maximum. This error is negligibly small.
//
// Returns:
//
//	Duration - the absolute value of this Duration instance.
func (duration Duration) Abs() Duration {
	switch {
	case duration >= 0:
		return duration
	case duration == Minimum:
		return Maximum
	default:
		return -duration
	}
}
