package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(DarkData)
}

var DarkData types.TypeData = types.TypeData{
	Bit: types.Dark,

	DoubleDamageTaken: types.Bug | types.Fairy | types.Fighting,
	HalfDamageTaken:   types.Dark | types.Ghost,
	ZeroDamageTaken:   types.Psychic,
}
