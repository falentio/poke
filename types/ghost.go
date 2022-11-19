package types

func init() {
	Register(GhostData)
}

var GhostData TypeData = TypeData{
	Bit: Ghost,

	DoubleDamageTaken: Dark | Ghost,
	HalfDamageTaken:   Bug | Poison,
	ZeroDamageTaken:   Fighting | Normal,
}
