package types

func init() {
	Register(GhostData)
}

var GhostData TypeData = TypeData{
	Type: Ghost,

	DoubleDamageTaken: []Type{
		Ghost,
		Dark,
	},
	HalfDamageTaken: []Type{
		Poison,
		Bug,
	},
	ZeroDamageTaken: []Type{
		Normal,
		Fighting,
	},
}
