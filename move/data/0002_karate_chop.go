package move_data

import (
	"github.com/falentio/poke/move"
	"github.com/falentio/poke/types"
)

func init() {
	move.Register(KarateChopData)
}

var KarateChopData move.Move = move.Move{
	Id: 2,
	Type: types.TypeFighting,
}

