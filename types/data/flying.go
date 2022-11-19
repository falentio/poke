package types

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(FlyingData)
}

var FlyingData types.TypeData = types.TypeData{
	Bit: types.Flying,

	DoubleDamageTaken: types.Electric | types.Ice | types.Rock,
	HalfDamageTaken:   types.Bug | types.Fighting | types.Grass,
	ZeroDamageTaken:   types.Ground,
}
