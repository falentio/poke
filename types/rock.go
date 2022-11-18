package types

func init() {
	Register(RockData)
}

var RockData TypeData = TypeData{
	Type: Rock,

	DoubleDamageTaken: []Type{
		Fighting,
		Ground,
		Steel,
		Water,
		Grass,
	},
	HalfDamageTaken: []Type{
		Normal,
		Flying,
		Poison,
		Fire,
	},
	ZeroDamageTaken: []Type{},
}
