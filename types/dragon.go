package types

func init() {
	Register(DragonData)
}

var DragonData TypeData = TypeData{
	Type: Dragon,

	DoubleDamageTaken: []Type{
		Ice,
		Dragon,
		Fairy,
	},
	HalfDamageTaken: []Type{
		Fire,
		Water,
		Grass,
		Electric,
	},
	ZeroDamageTaken: []Type{},
}
