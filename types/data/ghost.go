package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(GhostData)
}

var GhostData types.TypeData = types.TypeData{
	Bit: types.Ghost,

	DoubleDamageTaken: types.Dark | types.Ghost,
	HalfDamageTaken:   types.Bug | types.Poison,
	ZeroDamageTaken:   types.Fighting | types.Normal,
}
