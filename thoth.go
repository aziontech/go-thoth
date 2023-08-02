package unnamed

import (
	"math/rand"
	"strings"
	"time"
)

// random is a function that takes a slice of strings and returns a random string from the slice.
func random(el []string) string {
	rand.Seed(time.Now().UnixNano())
	return el[rand.Intn(len(el))]
}

// GenerateName generates a random name by combining random words from different categories.
// It returns the generated name as a string and any error encountered during the process.
func GenerateName() string {
	name := random(GetNouns()) + " " + random(GetAdjectives())
	name = strings.ReplaceAll(random(GetNouns())+" "+random(GetAdjectives()), " ", "_")
	name = strings.ToLower(name)
	return name
}
