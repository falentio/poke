package types

import (
	"testing"
)

func TestTypesTestTypes(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		Name       string
		Receiver   TypeBit
		Dealer     TypeBit
		Multiplier float64
		Err        error
	}{
		{"zero", Steel, Poison, 0, nil},
		{"half", Steel, Bug, 0.5, nil},
		{"double", Steel, Fire, 2, nil},
		{"quad", Steel | Grass, Fire, 2, nil},
		{"invalid dealer", Steel, Fighting | Fire, -1, ErrInvalidDealerType},
		{"no type Receiver", None, Fighting, -1, ErrZeroType},
		{"no type Dealer", Fighting, None, -1, ErrZeroType},
	} {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			m, err := GetMultiplier(tc.Receiver, tc.Dealer)
			if err != tc.Err {
				t.Errorf("unexpected errror: %v", err)
			}
			if m != tc.Multiplier {
				t.Errorf("unexpected multiplier returned: %f", m)
			}
		})
	}
}
