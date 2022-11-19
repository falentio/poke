package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(ElectricData)
}

var ElectricData types.TypeData = types.TypeData{
	Bit: types.Electric,

	DoubleDamageTaken: types.Ground,
	HalfDamageTaken:   types.Electric | types.Flying | types.Steel,
	ZeroDamageTaken:   types.None,
}
