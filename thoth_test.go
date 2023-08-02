package unnamed

import (
	"strings"
	"testing"
)

func TestRandom(t *testing.T) {
	el := []string{"apple", "banana", "orange", "grape"}

	result := random(el)

	found := false
	for _, e := range el {
		if e == result {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Expected result to be one of the elements in the slice, got %s", result)
	}
}

func TestGenerateName(t *testing.T) {
	name := GenerateName()

	if name == "" {
		t.Errorf("Generated name is empty")
	}

	if strings.Contains(name, " ") {
		t.Errorf("Generated name contains space")
	}

	if !strings.Contains(name, "_") {
		t.Errorf("Generated name does not contain an underscore")
	}
}
