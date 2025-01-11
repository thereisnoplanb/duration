package duration

// Gets the weeks component of the Duration.
//
// Returns:
//
//	int - The week component of this instance.
func (duration Duration) Weeks() int {
	return int(duration / Week)
}

// Gets the value of the current Duration expressed in whole and fractional weeks.
//
// Returns:
//
//	float64 - The total number of weeks represented by this instance.
func (duration Duration) TotalWeeks() float64 {
	return float64(duration) / float64(Week)
}
