package types

func init() {
	Register(PsychicData)
}

var PsychicData TypeData = TypeData{
	Type: Psychic,

	DoubleDamageTaken: []Type{
		Bug,
		Ghost,
		Dark,
	},
	HalfDamageTaken: []Type{
		Fighting,
		Psychic,
	},
	ZeroDamageTaken: []Type{},
}
