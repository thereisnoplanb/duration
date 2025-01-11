package duration

// Gets the minutes component of the Duration.
//
// Returns:
//
//	int - The minute component of this instance.
func (duration Duration) Minutes() int64 {
	return int64(duration / Minute)
}

// Gets the value of the current Duration expressed in whole and fractional minutes.
//
// Returns:
//
//	float64 - The total number of minutes represented by this instance.
func (duration Duration) TotalMinutes() float64 {
	return float64(duration) / float64(Minute)
}
