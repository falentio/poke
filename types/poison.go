package types

func init() {
	Register(PoisonData)
}

var PoisonData TypeData = TypeData{
	Type: Poison,

	DoubleDamageTaken: []Type{
		Ground,
		Psychic,
	},
	HalfDamageTaken: []Type{
		Fighting,
		Poison,
		Bug,
		Grass,
		Fairy,
	},
	ZeroDamageTaken: []Type{},
}
