package types_test

import (
	"testing"

	"github.com/falentio/poke/types"
	_ "github.com/falentio/poke/types/data"
)

func TestTypesTestTypes(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		Name       string
		Receiver   types.TypeBit
		Dealer     types.TypeBit
		Multiplier float64
		Err        error
	}{
		{"zero", types.Steel, types.Poison, 0, nil},
		{"half", types.Steel, types.Bug, 0.5, nil},
		{"double", types.Steel, types.Fire, 2, nil},
		{"quad", types.Steel | types.Grass, types.Fire, 2, nil},
		{"invalid dealer", types.Steel, types.Fighting | types.Fire, -1, types.ErrInvalidDealerType},
		{"no type Receiver", types.None, types.Fighting, -1, types.ErrZeroType},
		{"no type Dealer", types.Fighting, types.None, -1, types.ErrZeroType},
	} {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			m, err := types.GetMultiplier(tc.Receiver, tc.Dealer)
			if err != tc.Err {
				t.Errorf("unexpected errror: %v", err)
			}
			if m != tc.Multiplier {
				t.Errorf("unexpected multiplier returned: %f", m)
			}
		})
	}
}
