package moves

import (
	"github.com/falentio/poke/types"
)

type (
	DamageClass = uint8
)

const (
	PyhsicalDamage = iota
	SpecialDamage
)

type Move struct {
	Type types.TypeBit

	Accuracy    uint8
	DamageClass DamageClass
}
