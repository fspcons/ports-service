package domain_test

import (
	"fmt"
	"github.com/fspcons/ports-service/src/domain"
	"testing"
)

// TODO more domain structs tests would be defined here

func TestPortIsValid(t *testing.T) {
	scenarios := []struct {
		p        domain.Port
		expected bool
	}{
		{domain.Port{}, false},
		{domain.Port{ID: "   "}, false},
		{domain.Port{ID: "I am a valid ID"}, true},
	}

	t.Log("Given a list of scenarios")
	for _, s := range scenarios {
		scenarioText := fmt.Sprintf("domain.Port for %q should be %t", s.p.ID, s.expected)
		t.Run(scenarioText, func(t *testing.T) {
			if got := s.p.IsValid(); got != s.expected {
				t.Errorf("unexpected IsValid result: expected %t, got %t", got, s.expected)
			}
		})
	}
}
