package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(IceData)
}

var IceData types.TypeData = types.TypeData{
	Bit: types.Ice,

	DoubleDamageTaken: types.TypeNone | types.TypeFighting | types.TypeFire | types.TypeRock | types.TypeSteel,
	HalfDamageTaken:   types.TypeNone | types.TypeIce,
	ZeroDamageTaken:   types.TypeNone,
}
