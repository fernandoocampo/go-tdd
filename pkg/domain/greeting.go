package domain

import "time"

// Hello returns a hello following these rules.
// "Hello" -> if no name and hour is given
// "Hello Trinity" -> if a trinity name is given without an hour
func Hello(name string, now *time.Time) string {
	if name == "" && now == nil {
		return "Hello"
	}
	if now == nil {
		return "Hello " + name
	}
	partOfTheDay := calculatePartOfTheDay(*now)
	return "Hello " + name + ", Good " + partOfTheDay
}

// CalculatePartOfTheDay returns morning, afternoon, noon or evening
// depending of the given time.
func calculatePartOfTheDay(now time.Time) string {
	hour := now.Hour()
	minute := now.Minute()
	switch {
	case hour >= 0 && hour < 12:
		return "Morning"
	case hour == 12 && minute == 0:
		return "Noon"
	case hour >= 12 && hour <= 18:
		return "Afternoon"
	default:
		return "Evening"
	}
}
