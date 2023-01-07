package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(GroundData)
}

var GroundData types.TypeData = types.TypeData{
	Bit: types.Ground,

	DoubleDamageTaken: types.TypeNone | types.TypeGrass | types.TypeIce | types.TypeWater,
	HalfDamageTaken:   types.TypeNone | types.TypePoison | types.TypeRock,
	ZeroDamageTaken:   types.TypeNone | types.TypeElectric,
}
