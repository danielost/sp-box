package pbox

import (
	"fmt"
	"math"
	"testing"

	"github.com/danielost/sp-box/src/sbox"
)

func TestSBox(t *testing.T) {
	for i := 0; i <= math.MaxUint8; i++ {
		t.Run(fmt.Sprintf("S-box: %d", i), func(t *testing.T) {
			sI := sbox.Lookup(uint8(i))
			rsI := sbox.ReverseLookup(sI)
			if uint8(i) != rsI {
				t.Errorf("S-box: got = %v, want %v", rsI, i)
			}
		})
	}
}
