package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(ElectricData)
}

var ElectricData types.TypeData = types.TypeData{
	Bit: types.Electric,

	DoubleDamageTaken: types.TypeNone | types.TypeGround,
	HalfDamageTaken:   types.TypeNone | types.TypeElectric | types.TypeFlying | types.TypeSteel,
	ZeroDamageTaken:   types.TypeNone,
}
