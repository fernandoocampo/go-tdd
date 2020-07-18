package domain_test

import (
	"testing"

	"github.com/fernandoocampo/go-tdd/pkg/domain"
)

func TestIndividualGreeting(t *testing.T) {
	t.Run("without_name", func(t *testing.T) {
		// GIVEN
		givenName := ""
		want := "Hello Good morning"
		// WHEN
		got := domain.Hello(givenName, 0)
		// THEN
		assertHelloResult(t, want, got)
	})

	t.Run("with_trinity", func(t *testing.T) {
		// GIVEN
		givenName := "Trinity"
		want := "Hello Trinity, Good morning"
		// WHEN
		got := domain.Hello(givenName, 0)
		// THEN
		assertHelloResult(t, want, got)
	})

	t.Run("with_neo", func(t *testing.T) {
		// GIVEN
		givenName := "Neo"
		want := "Hello Neo, Good morning"
		// WHEN
		got := domain.Hello(givenName, 0)
		// THEN
		assertHelloResult(t, want, got)
	})
}

func TestHelloWithTestTables(t *testing.T) {
	// GIVEN
	tests := []struct {
		testcase  string
		givenName string
		givenHour int
		want      string
	}{
		{"with_trinity", "Trinity", 0, "Hello Trinity, Good morning"},
		{"with_neo", "Neo", 0, "Hello Neo, Good morning"},
		{"without_name", "", 0, "Hello Good morning"},
		{"without_name", "", 0, "Hello Good morning"},
		{"with_name_and_time", "Trinity", 0, "Hello Trinity, Good morning"},
		{"with_name_and_time", "Trinity", 13, "Hello Trinity, Good afternoon"},
	}

	for _, test := range tests {
		// WHEN
		got := domain.Hello(test.givenName, test.givenHour)
		assertHelloResult(t, test.want, got)
	}

}

func assertHelloResult(t *testing.T, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("expected %q got %q", want, got)
	}
}
