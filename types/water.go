package types

func init() {
	Register(WaterData)
}

var WaterData TypeData = TypeData{
	Type: Water,

	DoubleDamageTaken: []Type{
		Grass,
		Electric,
	},
	HalfDamageTaken: []Type{
		Steel,
		Fire,
		Water,
		Ice,
	},
	ZeroDamageTaken: []Type{},
}
