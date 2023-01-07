package move_data

import (
	"github.com/falentio/poke/move"
	"github.com/falentio/poke/types"
)

func init() {
	move.Register(DoubleSlapData)
}

var DoubleSlapData move.Move = move.Move{
	Id: 3,
	Type: types.TypeNormal,
}

