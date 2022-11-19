package pokemon

import (
	"github.com/falentio/poke/types"
)

type Pokemon struct {
	BaseStats Stat
	IV        Stat
	EV        Stat

	Type types.TypeBit
}

type Stat struct {
	HP             uint8
	Attack         uint8
	Deffense       uint8
	SpecialAttack  uint8
	SpecialDefense uint8
	Speed          uint8
}
