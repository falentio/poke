package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(FightingData)
}

var FightingData types.TypeData = types.TypeData{
	Bit: types.Fighting,

	DoubleDamageTaken: types.TypeNone | types.TypeFairy | types.TypeFlying | types.TypePsychic,
	HalfDamageTaken:   types.TypeNone | types.TypeBug | types.TypeDark | types.TypeRock,
	ZeroDamageTaken:   types.TypeNone,
}
