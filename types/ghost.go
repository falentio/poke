package types

func init() {
	Register(GhostData)
}

var GhostData TypeData = TypeData{
	Name: "ghost",
	Bit:  Ghost,

	DoubleDamageTaken: None | Dark | Ghost,
	HalfDamageTaken:   None | Bug | Poison,
	ZeroDamageTaken:   None | Fighting | Normal,
}
