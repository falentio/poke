package types

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(GrassData)
}

var GrassData types.TypeData = types.TypeData{
	Bit: types.Grass,

	DoubleDamageTaken: types.Bug | types.Fire | types.Flying | types.Ice | types.Poison,
	HalfDamageTaken:   types.Electric | types.Grass | types.Ground | types.Water,
	ZeroDamageTaken:   types.None,
}
