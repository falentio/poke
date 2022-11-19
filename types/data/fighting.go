package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(FightingData)
}

var FightingData types.TypeData = types.TypeData{
	Bit: types.Fighting,

	DoubleDamageTaken: types.Fairy | types.Flying | types.Psychic,
	HalfDamageTaken:   types.Bug | types.Dark | types.Rock,
	ZeroDamageTaken:   types.None,
}
