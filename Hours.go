package duration

// Gets the hours component of the Duration.
//
// Returns:
//
//	int - The hour component of this instance.
func (duration Duration) Hours() int {
	return int(duration / Hour)
}

// Gets the value of the current Duration expressed in whole and fractional hours.
//
// Returns:
//
//	float64 - The total number of hours represented by this instance.
func (duration Duration) TotalHours() float64 {
	return float64(duration) / float64(Hour)
}
