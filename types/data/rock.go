package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(RockData)
}

var RockData types.TypeData = types.TypeData{
	Bit: types.Rock,

	DoubleDamageTaken: types.TypeNone | types.TypeFighting | types.TypeGrass | types.TypeGround | types.TypeSteel | types.TypeWater,
	HalfDamageTaken:   types.TypeNone | types.TypeFire | types.TypeFlying | types.TypeNormal | types.TypePoison,
	ZeroDamageTaken:   types.TypeNone,
}
