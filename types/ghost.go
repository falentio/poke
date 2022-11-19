package types

func init() {
	Register(GhostData)
}

var GhostData TypeData = TypeData{
	Bit: Ghost,

	DoubleDamageTaken: None | Dark | Ghost,
	HalfDamageTaken:   None | Bug | Poison,
	ZeroDamageTaken:   None | Fighting | Normal,
}
