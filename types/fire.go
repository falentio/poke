package types

func init() {
	Register(FireData)
}

var FireData TypeData = TypeData{
	Type: Fire,

	DoubleDamageTaken: []Type{
		Ground,
		Rock,
		Water,
	},
	HalfDamageTaken: []Type{
		Bug,
		Steel,
		Fire,
		Grass,
		Ice,
		Fairy,
	},
	ZeroDamageTaken: []Type{},
}
