package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(BugData)
}

var BugData types.TypeData = types.TypeData{
	Bit: types.Bug,

	DoubleDamageTaken: types.TypeNone | types.TypeFire | types.TypeFlying | types.TypeRock,
	HalfDamageTaken:   types.TypeNone | types.TypeFighting | types.TypeGrass | types.TypeGround,
	ZeroDamageTaken:   types.TypeNone,
}
