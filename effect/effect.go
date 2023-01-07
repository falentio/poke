package effect

import "github.com/falentio/poke/field"

type Effect interface{}

type Base struct{}

type BattleEvent struct{
	Field *field.Field
}

func (Base) BeforeEnterBattle(ev BattleEvent) {}
func (Base) BeforeLeaveBattle(ev BattleEvent) {}
func (Base) BeforeTurnStart() {}
func (Base) BeforeTurnEnd() {}
func (Base) BeforeSelfFainted() {}