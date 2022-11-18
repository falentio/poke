package types

func init() {
	Register(SteelData)
}

var SteelData TypeData = TypeData{
	Type: Steel,

	DoubleDamageTaken: []Type{
		Fighting,
		Ground,
		Fire,
	},
	HalfDamageTaken: []Type{
		Normal,
		Flying,
		Rock,
		Bug,
		Steel,
		Grass,
		Psychic,
		Ice,
		Dragon,
		Fairy,
	},
	ZeroDamageTaken: []Type{
		Poison,
	},
}
