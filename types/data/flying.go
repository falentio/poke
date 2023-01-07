package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(FlyingData)
}

var FlyingData types.TypeData = types.TypeData{
	Bit: types.Flying,

	DoubleDamageTaken: types.TypeNone | types.TypeElectric | types.TypeIce | types.TypeRock,
	HalfDamageTaken:   types.TypeNone | types.TypeBug | types.TypeFighting | types.TypeGrass,
	ZeroDamageTaken:   types.TypeNone | types.TypeGround,
}
