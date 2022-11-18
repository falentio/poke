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
	Name string
	Type types.Type

	Accuracy    uint8
	DamageClass DamageClass
}
