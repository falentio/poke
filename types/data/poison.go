package types

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(PoisonData)
}

var PoisonData types.TypeData = types.TypeData{
	Bit: types.Poison,

	DoubleDamageTaken: types.Ground | types.Psychic,
	HalfDamageTaken:   types.Bug | types.Fairy | types.Fighting | types.Grass | types.Poison,
	ZeroDamageTaken:   types.None,
}
