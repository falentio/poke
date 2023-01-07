package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(FireData)
}

var FireData types.TypeData = types.TypeData{
	Bit: types.Fire,

	DoubleDamageTaken: types.TypeNone | types.TypeGround | types.TypeRock | types.TypeWater,
	HalfDamageTaken:   types.TypeNone | types.TypeBug | types.TypeFairy | types.TypeFire | types.TypeGrass | types.TypeIce | types.TypeSteel,
	ZeroDamageTaken:   types.TypeNone,
}
