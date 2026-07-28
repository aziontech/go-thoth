package unnamed

import (
	"strings"
	"testing"
)

// FuzzRandom exercises random() over arbitrary element sets. The elements are
// carried in one seed string and split on NUL so the fuzzer can vary both how
// many there are and what they contain.
func FuzzRandom(f *testing.F) {
	f.Add("apple\x00banana\x00orange\x00grape")
	f.Add("")
	f.Add("\x00\x00")
	f.Add("-")

	f.Fuzz(func(t *testing.T, seed string) {
		// Split always yields at least one element, so random() never sees the
		// empty slice that would panic inside rng.Intn.
		el := strings.Split(seed, "\x00")

		result := random(el)

		for _, e := range el {
			if e == result {
				return
			}
		}
		t.Errorf("random(%q) = %q, which is not one of the elements", el, result)
	})
}
