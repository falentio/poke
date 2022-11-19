package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(WaterData)
}

var WaterData types.TypeData = types.TypeData{
	Bit: types.Water,

	DoubleDamageTaken: types.Electric | types.Grass,
	HalfDamageTaken:   types.Fire | types.Ice | types.Steel | types.Water,
	ZeroDamageTaken:   types.None,
}
