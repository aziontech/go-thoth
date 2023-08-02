package unnamed

import (
	"math/rand"
	"strings"
	"time"
)

var rng *rand.Rand

// CreateRand initializes the random generator.
func createRand() {
	seed := time.Now().UnixNano()
	src := rand.NewSource(seed)
	rng = rand.New(src)
}

// random is a function that takes a slice of strings and returns a random string from the slice.
func random(el []string) string {
	createRand()
	return el[rng.Intn(len(el))]
}

// GenerateName generates a random name by combining random words from different categories.
// It returns the generated name as a string and any error encountered during the process.
func GenerateName() string {
	name := random(GetAdjectives()) + "_" + random(GetNouns())
	name = strings.ToLower(name)
	return name
}
