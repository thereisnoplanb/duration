package duration

// Gets the milliseconds component of the Duration.
//
// Returns:
//
//	int64 - The millisecond component of this instance.
func (duration Duration) Millisecond() int64 {
	return int64(duration / Millisecond)
}

// Gets the value of the current Duration expressed in whole and fractional milliseconds.
//
// Returns:
//
//	float64 - The total number of milliseconds represented by this instance.
func (duration Duration) TotalMillisecond() float64 {
	return float64(duration) / float64(Millisecond)
}
