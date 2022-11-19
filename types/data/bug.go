package types

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(BugData)
}

var BugData types.TypeData = types.TypeData{
	Bit: types.Bug,

	DoubleDamageTaken: types.Fire | types.Flying | types.Rock,
	HalfDamageTaken:   types.Fighting | types.Grass | types.Ground,
	ZeroDamageTaken:   types.None,
}
