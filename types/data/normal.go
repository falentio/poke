package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(NormalData)
}

var NormalData types.TypeData = types.TypeData{
	Bit: types.Normal,

	DoubleDamageTaken: types.Fighting,
	HalfDamageTaken:   types.None,
	ZeroDamageTaken:   types.Ghost,
}
