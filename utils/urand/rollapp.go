package urand

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/rand"
)

// RollappID generates a unique rollapp ID, following the pattern: "name_1234-1"
func RollappID() string {
	name := make([]byte, 8)
	for i := range name {
		name[i] = byte(rand.Intn('z'-'a'+1) + 'a')
	}
	return fmt.Sprintf("%s_%d-1", string(name), rand.Int63())
}
