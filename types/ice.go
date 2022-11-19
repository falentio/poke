package types

func init() {
	Register(IceData)
}

var IceData TypeData = TypeData{
	Bit: Ice,

	DoubleDamageTaken: Fighting | Fire | Rock | Steel,
	HalfDamageTaken:   Ice,
	ZeroDamageTaken:   None,
}
