package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(WaterData)
}

var WaterData types.TypeData = types.TypeData{
	Bit: types.Water,

	DoubleDamageTaken: types.TypeNone | types.TypeElectric | types.TypeGrass,
	HalfDamageTaken:   types.TypeNone | types.TypeFire | types.TypeIce | types.TypeSteel | types.TypeWater,
	ZeroDamageTaken:   types.TypeNone,
}
