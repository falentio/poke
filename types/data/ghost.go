package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(GhostData)
}

var GhostData types.TypeData = types.TypeData{
	Bit: types.Ghost,

	DoubleDamageTaken: types.TypeNone | types.TypeDark | types.TypeGhost,
	HalfDamageTaken:   types.TypeNone | types.TypeBug | types.TypePoison,
	ZeroDamageTaken:   types.TypeNone | types.TypeFighting | types.TypeNormal,
}
