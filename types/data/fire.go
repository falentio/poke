package types

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(FireData)
}

var FireData types.TypeData = types.TypeData{
	Bit: types.Fire,

	DoubleDamageTaken: types.Ground | types.Rock | types.Water,
	HalfDamageTaken:   types.Bug | types.Fairy | types.Fire | types.Grass | types.Ice | types.Steel,
	ZeroDamageTaken:   types.None,
}
