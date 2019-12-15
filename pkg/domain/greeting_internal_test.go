package domain

import (
	"testing"
	"time"
)

func TestCalculatePartOfTheDay(t *testing.T) {
	tests := []struct {
		testcase  string
		givenTime time.Time
		want      string
	}{
		{"morning", newTime(t, 8, 0), "Morning"},
		{"afternoon", newTime(t, 13, 0), "Afternoon"},
		{"evening", newTime(t, 23, 0), "Evening"},
		{"noon", newTime(t, 12, 0), "Noon"},
	}

	for _, test := range tests {
		got := calculatePartOfTheDay(test.givenTime)
		if test.want != got {
			t.Errorf("%q expected %q, but got %q", test.testcase, test.want, got)
		}
	}
}

func newTime(t *testing.T, hour, minute int) time.Time {
	t.Helper()
	return time.Date(2019, time.December, 17, hour, minute, 0, 0, time.UTC)
}
