package battle

import (
	"github.com/falentio/poke/field"
	"github.com/falentio/poke/move"
)

type State struct {
	Turns []*Turn
}

type Turn struct{
	Item any
	Move *move.Move
}

type Battle struct {
	Field *field.Field
	State *State
}

func (b *Battle) Init() error {
	return nil
}

func (b *Battle) Start() error {
	return nil
}

func (b *Battle) PlayTurn(side uint8, turn *Turn) error {
	return nil
}

func (b *Battle) NextTurn() error {
	return nil
}
