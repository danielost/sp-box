package pbox

import (
	"fmt"
	"math"
	"testing"

	"github.com/danielost/sp-box/src/pbox"
)

func TestPBox(t *testing.T) {
	for i := 0; i <= math.MaxUint8; i++ {
		t.Run(fmt.Sprintf("P-box: %d", i), func(t *testing.T) {
			pI, reverseI := pbox.Permute(uint8(i))
			rpI := reverseI(pI)
			if uint8(i) != rpI {
				t.Errorf("P-box: got = %v, want %v", rpI, i)
			}
		})
	}
}
