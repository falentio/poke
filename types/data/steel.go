package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(SteelData)
}

var SteelData types.TypeData = types.TypeData{
	Bit: types.Steel,

	DoubleDamageTaken: types.TypeNone | types.TypeFighting | types.TypeFire | types.TypeGround,
	HalfDamageTaken:   types.TypeNone | types.TypeBug | types.TypeDragon | types.TypeFairy | types.TypeFlying | types.TypeGrass | types.TypeIce | types.TypeNormal | types.TypePsychic | types.TypeRock | types.TypeSteel,
	ZeroDamageTaken:   types.TypeNone | types.TypePoison,
}
