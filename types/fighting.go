package types

func init() {
	Register(FightingData)
}

var FightingData TypeData = TypeData{
	Type: Fighting,

	DoubleDamageTaken: []Type{
		Flying,
		Psychic,
		Fairy,
	},
	HalfDamageTaken: []Type{
		Rock,
		Bug,
		Dark,
	},
	ZeroDamageTaken: []Type{},
}
