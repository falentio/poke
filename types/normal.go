package types

func init() {
	Register(NormalData)
}

var NormalData TypeData = TypeData{
	Type: Normal,

	DoubleDamageTaken: []Type{
		Fighting,
	},
	HalfDamageTaken: []Type{},
	ZeroDamageTaken: []Type{
		Ghost,
	},
}
