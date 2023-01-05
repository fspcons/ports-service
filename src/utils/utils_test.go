package utils_test

import (
	"fmt"
	"github.com/fspcons/ports-service/src/utils"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	scenarios := map[string]bool{"    ": true, "": true, "test": false}

	t.Log("Given a list of scenarios")
	for s, v := range scenarios {
		scenarioText := fmt.Sprintf("IsEmpty for %q should be %t", s, v)
		t.Run(scenarioText, func(t *testing.T) {
			result := utils.IsEmpty(s)

			if result != v {
				t.Errorf("unexpected IsEmpty result: expected %t, got %t", result, v)
			}
		})
	}
}
