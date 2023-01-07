package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(DragonData)
}

var DragonData types.TypeData = types.TypeData{
	Bit: types.Dragon,

	DoubleDamageTaken: types.TypeNone | types.TypeDragon | types.TypeFairy | types.TypeIce,
	HalfDamageTaken:   types.TypeNone | types.TypeElectric | types.TypeFire | types.TypeGrass | types.TypeWater,
	ZeroDamageTaken:   types.TypeNone,
}
