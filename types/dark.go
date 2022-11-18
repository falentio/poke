package types

func init() {
	Register(DarkData)
}

var DarkData TypeData = TypeData{
	Type: Dark,

	DoubleDamageTaken: []Type{
		Fighting,
		Bug,
		Fairy,
	},
	HalfDamageTaken: []Type{
		Ghost,
		Dark,
	},
	ZeroDamageTaken: []Type{
		Psychic,
	},
}
