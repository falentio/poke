package types_data

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register(FairyData)
}

var FairyData types.TypeData = types.TypeData{
	Bit: types.Fairy,

	DoubleDamageTaken: types.TypeNone | types.TypePoison | types.TypeSteel,
	HalfDamageTaken:   types.TypeNone | types.TypeBug | types.TypeDark | types.TypeFighting,
	ZeroDamageTaken:   types.TypeNone | types.TypeDragon,
}
