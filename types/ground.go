package types

func init() {
	Register(GroundData)
}

var GroundData TypeData = TypeData{
	Type: Ground,

	DoubleDamageTaken: []Type{
		Water,
		Grass,
		Ice,
	},
	HalfDamageTaken: []Type{
		Poison,
		Rock,
	},
	ZeroDamageTaken: []Type{
		Electric,
	},
}
