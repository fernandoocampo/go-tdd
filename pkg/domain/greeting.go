package domain

// Hello returns a "Hello"
func Hello(username string, hour int) string {
	if username == "" && hour >= 0 && hour < 12 {
		return "Hello" + username + " Good " + calculatePartOfTheDay(hour)
	}
	if username == "" && hour >= 12 && hour <= 18 {
		return "Hello" + username + " Good " + calculatePartOfTheDay(hour)
	}

	if hour >= 0 && hour < 12 {
		return "Hello " + username + ", Good " + calculatePartOfTheDay(hour)
	}
	if hour >= 12 && hour <= 18 {
		return "Hello " + username + ", Good " + calculatePartOfTheDay(hour)
	}
	return "Hello " + username + ", Good " + calculatePartOfTheDay(hour)
}

func calculatePartOfTheDay(hour int) string {
	switch {
	case hour >= 0 && hour < 12:
		return "morning"
	case hour >= 12 && hour <= 18:
		return "afternoon"
	default:
		return "evening"
	}
}
