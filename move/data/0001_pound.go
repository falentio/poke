package move_data

import (
	"github.com/falentio/poke/move"
	"github.com/falentio/poke/types"
)

func init() {
	move.Register(PoundData)
}

var PoundData move.Move = move.Move{
	Id: 1,
	Type: types.TypeNormal,
}

