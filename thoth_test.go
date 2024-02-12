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

	if !strings.Contains(name, "-") {
		t.Errorf("Generated name does not contain an underscore")
	}
}

func TestDuplicates(t *testing.T) {
	itemExists := (func(list []string, item string) bool {
		for _, el := range list {
			if el == item {
				return true
			}
		}
		return false
	})

	var duplicates []string
	nouns := GetNouns()
	adjs := GetAdjectives()
	
	for idx, el := range nouns {
		if itemExists(nouns[:idx], el) {
			duplicates = append(duplicates, el)
		}
	}

	for idx, el := range adjs {
		if itemExists(adjs[:idx], el) {
			duplicates = append(duplicates, el)
		}
	}

	if len(duplicates) > 0 {
		t.Errorf("has duplicate items: %v", duplicates)
	}
}
