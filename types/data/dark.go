package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(DarkData)
}

var DarkData types.TypeData = types.TypeData{
	Bit: types.Dark,

	DoubleDamageTaken: types.TypeNone | types.TypeBug | types.TypeFairy | types.TypeFighting,
	HalfDamageTaken:   types.TypeNone | types.TypeDark | types.TypeGhost,
	ZeroDamageTaken:   types.TypeNone | types.TypePsychic,
}
