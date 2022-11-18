package types

func init() {
	Register(GrassData)
}

var GrassData TypeData = TypeData{
	Type: Grass,

	DoubleDamageTaken: []Type{
		Flying,
		Poison,
		Bug,
		Fire,
		Ice,
	},
	HalfDamageTaken: []Type{
		Ground,
		Water,
		Grass,
		Electric,
	},
	ZeroDamageTaken: []Type{},
}
