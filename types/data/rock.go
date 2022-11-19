package types

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(RockData)
}

var RockData types.TypeData = types.TypeData{
	Bit: types.Rock,

	DoubleDamageTaken: types.Fighting | types.Grass | types.Ground | types.Steel | types.Water,
	HalfDamageTaken:   types.Fire | types.Flying | types.Normal | types.Poison,
	ZeroDamageTaken:   types.None,
}
