package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(PsychicData)
}

var PsychicData types.TypeData = types.TypeData{
	Bit: types.Psychic,

	DoubleDamageTaken: types.Bug | types.Dark | types.Ghost,
	HalfDamageTaken:   types.Fighting | types.Psychic,
	ZeroDamageTaken:   types.None,
}
