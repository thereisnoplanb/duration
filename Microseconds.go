package duration

// Gets the microseconds component of the Duration.
//
// Returns:
//
//	int64 - The microsecond component of this instance.
func (duration Duration) Microseconds() int64 {
	return int64(duration / Microsecond)
}

// Gets the value of the current Duration expressed in whole and fractional microseconds.
//
// Returns:
//
//	float64 - The total number of microseconds represented by this instance.
func (duration Duration) TotalMicroseconds() float64 {
	return float64(duration) / float64(Microsecond)
}
