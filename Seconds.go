package duration

// Gets the seconds component of the Duration.
//
// Returns:
//
//	int - The second component of this instance.
func (duration Duration) Seconds() int64 {
	return int64(duration / Second)
}

// Gets the value of the current Duration expressed in whole and fractional seconds.
//
// Returns:
//
//	float64 - The total number of seconds represented by this instance.
func (duration Duration) TotalSeconds() float64 {
	return float64(duration) / float64(Second)
}
