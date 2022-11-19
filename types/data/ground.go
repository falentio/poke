package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(GroundData)
}

var GroundData types.TypeData = types.TypeData{
	Bit: types.Ground,

	DoubleDamageTaken: types.Grass | types.Ice | types.Water,
	HalfDamageTaken:   types.Poison | types.Rock,
	ZeroDamageTaken:   types.Electric,
}
