package types

func init() {
	Register(ElectricData)
}

var ElectricData TypeData = TypeData{
	Type: Electric,

	DoubleDamageTaken: []Type{
		Ground,
	},
	HalfDamageTaken: []Type{
		Flying,
		Steel,
		Electric,
	},
	ZeroDamageTaken: []Type{},
}
