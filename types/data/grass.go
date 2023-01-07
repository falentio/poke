package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(GrassData)
}

var GrassData types.TypeData = types.TypeData{
	Bit: types.Grass,

	DoubleDamageTaken: types.TypeNone | types.TypeBug | types.TypeFire | types.TypeFlying | types.TypeIce | types.TypePoison,
	HalfDamageTaken:   types.TypeNone | types.TypeElectric | types.TypeGrass | types.TypeGround | types.TypeWater,
	ZeroDamageTaken:   types.TypeNone,
}
