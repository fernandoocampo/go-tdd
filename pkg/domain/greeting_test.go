package domain_test

import (
	"testing"
	"time"

	"github.com/fernandoocampo/go-tdd/pkg/domain"
)

func TestHello(t *testing.T) {
	t.Run("without_time", func(t *testing.T) {
		tests := []struct {
			testcase  string
			givenName string
			givenTime *time.Time
			want      string
		}{
			{"no_user", "", nil, "Hello"},
			{"with_trinity", "Trinity", nil, "Hello Trinity"},
			{"with_niobe", "Niobe", nil, "Hello Niobe"},
			{"with_avantica", "Avantica", nil, "Hello Avantica"},
			{"with_matrix", "matrix", nil, "Hello matrix"},
			{"with_trinity_morning", "Trinity", newTime(t, 5, 0), "Hello Trinity, Good Morning"},
			{"with_niobe_morning", "Niobe", newTime(t, 9, 0), "Hello Niobe, Good Morning"},
			{"with_trinity_afternoon", "Trinity", newTime(t, 13, 0), "Hello Trinity, Good Afternoon"},
			{"with_trinity_noon", "Trinity", newTime(t, 12, 0), "Hello Trinity, Good Noon"},
		}
		// WHEN
		for _, test := range tests {
			got := domain.Hello(test.givenName, test.givenTime)
			assertHelloResult(t, test.testcase, test.want, got)
		}
	})
}

// assertHelloResult helper function to check hello result.
func assertHelloResult(t *testing.T, testcase, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("test %q: expected: %q but got: %q", testcase, want, got)
	}
}

func NewInt(value int) *int {
	return &value
}

func newTime(t *testing.T, hour, minute int) *time.Time {
	t.Helper()
	newtime := time.Date(2019, time.December, 17, hour, minute, 0, 0, time.UTC)
	return &newtime
}
