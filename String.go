package duration

// Converts the value of the current Duration object to its equivalent string representation.
// String returns a string representing the duration in the form [-][d.]hh:mm:ss[.fffffff].
//
// Returns:
//
//	string - current Duration string representation.
func (duration Duration) String() string {
	var array [26]byte
	index := duration.format(&array)
	return string(array[index:])
}

func (duration Duration) format(array *[26]byte) int {
	index := len(array)

	value := uint64(duration)

	negative := duration < Zero
	if negative {
		value = -value
	}

	if value%Second > 0 {
		index = fmtInt(array[:index], value%Second, 7)
		index--
		array[index] = '.'
	}

	value /= Second
	index = fmtInt(array[:index], value%secondsPerMinute, 2)
	index--
	array[index] = ':'

	value /= secondsPerMinute
	index = fmtInt(array[:index], value%minutesPerHour, 2)
	index--
	array[index] = ':'

	value /= minutesPerHour
	index = fmtInt(array[:index], value%hoursPerDay, 2)

	value /= hoursPerDay
	if value > 0 {
		index--
		array[index] = '.'
		index = fmtInt(array[:index], value, 1)
	}

	if negative {
		index--
		array[index] = '-'
	}

	return index
}

func fmtInt(array []byte, value uint64, digits int) int {
	index := len(array)
	for do := true; do; do = digits > 0 || value > 0 {
		index--
		digits--
		array[index] = byte(value%10) + '0'
		value /= 10
	}
	return index
}
