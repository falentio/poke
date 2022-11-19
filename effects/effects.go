package effects

import (
	"github.com/falentio/poke/pokemon"
	"github.com/falentio/poke/battle"
)

type EffectContext struct {
	Pokemon *pokemon.Pokemon
	Battle  *battle.Battle
}

type EffectBase struct{}

func (eb *EffectBase) OnBattleStart(ctx *EffectContext)      {}
func (eb *EffectBase) OnBattleBeforeTurn(ctx *EffectContext) {}
func (eb *EffectBase) OnBattleAfterTurn(ctx *EffectContext)  {}
func (eb *EffectBase) OnBattleEnd(ctx *EffectContext)        {}
