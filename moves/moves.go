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

	Accuracy    uint8
	DamageClass DamageClass
}
