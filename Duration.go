// Represents a time interval.
package duration

// Represents a time interval.
type Duration int64

// Represents the 1 tick.
// 1 tick is equal to 100ns.
const Tick = 1
const ticksPerMicrosecond = 10

// Represents the 1 mircosecond.
const Microsecond = ticksPerMicrosecond * Tick
const microsecondsPerSecond = 1000

// Represents the 1 millisecond.
const Millisecond = microsecondsPerSecond * Microsecond
const millisecondsPerSecond = 1000

// Represents the 1 second.
const Second = millisecondsPerSecond * Millisecond
const secondsPerMinute = 60

// Represents the 1 minute.
const Minute = secondsPerMinute * Second
const minutesPerHour = 60

// Represents the 1 hour.
const Hour = minutesPerHour * Minute
const hoursPerDay = 24

// Represents the 1 day.
const Day = hoursPerDay * Hour
const daysPerWeek = 7

// Represents the 1 week.
const Week = daysPerWeek * Day

// Represents the maximum Duration value.
// Maimum value is equal to 10675199.02:48:05.4775807. (Approx. 29277 years.)
const Maximum Duration = 1<<63 - 1

// Represents the minimum Duration value.
// Minimum value is equal to -10675199.02:48:05.4775808. (Approx. -29277 years.)
const Minimum Duration = -1 << 63

// Represents the zero Duration value.
const Zero Duration = 0
