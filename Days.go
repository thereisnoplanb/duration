package duration

// Gets the days component of the Duration.
//
// Returns:
//
//	int - The day component of this instance.
func (duration Duration) Days() int {
	return int(duration / Day)
}

// Gets the value of the current Duration expressed in whole and fractional days.
//
// Returns:
//
//	float64 - The total number of days represented by this instance.
func (duration Duration) TotalDays() float64 {
	return float64(duration) / float64(Day)
}
