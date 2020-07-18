package domain

import "testing"

func TestCalculatePartOfTheDay(t *testing.T) {
	t.Run("afternoon", func(t *testing.T) {
		// GIVEN
		givenTime := 13
		want := "afternoon"
		// WHEN
		got := calculatePartOfTheDay(givenTime)
		// THEN
		if want != got {
			t.Errorf("expected %q got %q", want, got)
		}
	})
	t.Run("morning", func(t *testing.T) {
		// GIVEN
		givenTime := 9
		want := "morning"
		// WHEN
		got := calculatePartOfTheDay(givenTime)
		// THEN
		if want != got {
			t.Errorf("expected %q got %q", want, got)
		}
	})
}
