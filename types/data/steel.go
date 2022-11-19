package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(SteelData)
}

var SteelData types.TypeData = types.TypeData{
	Bit: types.Steel,

	DoubleDamageTaken: types.Fighting | types.Fire | types.Ground,
	HalfDamageTaken:   types.Bug | types.Dragon | types.Fairy | types.Flying | types.Grass | types.Ice | types.Normal | types.Psychic | types.Rock | types.Steel,
	ZeroDamageTaken:   types.Poison,
}
