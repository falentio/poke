package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(PsychicData)
}

var PsychicData types.TypeData = types.TypeData{
	Bit: types.Psychic,

	DoubleDamageTaken: types.TypeNone | types.TypeBug | types.TypeDark | types.TypeGhost,
	HalfDamageTaken:   types.TypeNone | types.TypeFighting | types.TypePsychic,
	ZeroDamageTaken:   types.TypeNone,
}
