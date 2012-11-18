// Generate semi-strong random number seeds.
package seed

import (
	crypt "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
)

// Return a seed that is usable as a input to math/rand. Uses
// crypto/rand as a source. The seeds have true entropy in them, but
// are not necessarily cryptographically safe.
func Seed() int64 {
	var seed int64
	err := binary.Read(crypt.Reader, binary.BigEndian, &seed)
	if err != nil {
		panic(fmt.Sprintf("Reading entropy failed: %v", err))
	}
	return seed
}

// Create a randomly-seeded pseudo-random number source
// See math/rand.
func Source() rand.Source {
	return rand.NewSource(Seed())
}

// Create a randomly-seeded pseudo-random number generator.
// See math/rand.
func Rand() *rand.Rand {
	return rand.New(rand.NewSource(Seed()))
}
