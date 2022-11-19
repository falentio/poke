package types

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(IceData)
}

var IceData types.TypeData = types.TypeData{
	Bit: types.Ice,

	DoubleDamageTaken: types.Fighting | types.Fire | types.Rock | types.Steel,
	HalfDamageTaken:   types.Ice,
	ZeroDamageTaken:   types.None,
}
