package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(DragonData)
}

var DragonData types.TypeData = types.TypeData{
	Bit: types.Dragon,

	DoubleDamageTaken: types.Dragon | types.Fairy | types.Ice,
	HalfDamageTaken:   types.Electric | types.Fire | types.Grass | types.Water,
	ZeroDamageTaken:   types.None,
}
