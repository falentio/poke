package types

func init() {
	Register(FairyData)
}

var FairyData TypeData = TypeData{
	Type: Fairy,

	DoubleDamageTaken: []Type{
		Poison,
		Steel,
	},
	HalfDamageTaken: []Type{
		Fighting,
		Bug,
		Dark,
	},
	ZeroDamageTaken: []Type{
		Dragon,
	},
}
