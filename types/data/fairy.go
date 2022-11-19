package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(FairyData)
}

var FairyData types.TypeData = types.TypeData{
	Bit: types.Fairy,

	DoubleDamageTaken: types.Poison | types.Steel,
	HalfDamageTaken:   types.Bug | types.Dark | types.Fighting,
	ZeroDamageTaken:   types.Dragon,
}
