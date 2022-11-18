package types

func init() {
	Register(IceData)
}

var IceData TypeData = TypeData{
	Type: Ice,

	DoubleDamageTaken: []Type{
		Fighting,
		Rock,
		Steel,
		Fire,
	},
	HalfDamageTaken: []Type{
		Ice,
	},
	ZeroDamageTaken: []Type{},
}
