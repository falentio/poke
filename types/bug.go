package types

func init() {
	Register(BugData)
}

var BugData TypeData = TypeData{
	Type: Bug,

	DoubleDamageTaken: []Type{
		Flying,
		Rock,
		Fire,
	},
	HalfDamageTaken: []Type{
		Fighting,
		Ground,
		Grass,
	},
	ZeroDamageTaken: []Type{},
}
