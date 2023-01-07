package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(PoisonData)
}

var PoisonData types.TypeData = types.TypeData{
	Bit: types.Poison,

	DoubleDamageTaken: types.TypeNone | types.TypeGround | types.TypePsychic,
	HalfDamageTaken:   types.TypeNone | types.TypeBug | types.TypeFairy | types.TypeFighting | types.TypeGrass | types.TypePoison,
	ZeroDamageTaken:   types.TypeNone,
}
