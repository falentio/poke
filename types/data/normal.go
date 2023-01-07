package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(NormalData)
}

var NormalData types.TypeData = types.TypeData{
	Bit: types.Normal,

	DoubleDamageTaken: types.TypeNone | types.TypeFighting,
	HalfDamageTaken:   types.TypeNone,
	ZeroDamageTaken:   types.TypeNone | types.TypeGhost,
}
