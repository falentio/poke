package move

type Target uint8

// Reference: "https://github.com/smogon/pokenmon-showdown/tree/master/sim/dex-moves.ts#L21-L24"
const (
	TargetAdjacentAlly Target = iota
	TargetAdjacentAllyOrSelf
	TargetAdjacentFoe
	TargetAll
	TargetAllAdjacent
	TargetAllAdjacentFoes
	TargetAllies
	TargetAllySide
	TargetAllyTeam
	TargetAny
	TargetFoeSide
	TargetNormal
	TargetRandomNormal
	TargetScripted
	TargetSelf
)
