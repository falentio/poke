package types

func init() {
	Register(FlyingData)
}

var FlyingData TypeData = TypeData{
	Type: Flying,

	DoubleDamageTaken: []Type{
		Rock,
		Electric,
		Ice,
	},
	HalfDamageTaken: []Type{
		Fighting,
		Bug,
		Grass,
	},
	ZeroDamageTaken: []Type{
		Ground,
	},
}
