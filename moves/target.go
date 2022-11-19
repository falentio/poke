package moves

type MoveTarget uint8

// Reference: "https://github.com/smogon/pokenmon-showdown/tree/master/sim/dex-moves.ts#L21-L24"
const (
	AdjacentAlly = iota
	AdjacentAllyOrSelf 
	AdjacentFoe 
	All 
	AllAdjacent 
	AllAdjacentFoes
	Allies 
	AllySide 
	AllyTeam 
	Any 
	FoeSide 
	Normal 
	RandomNormal 
	Scripted 
	Self
)