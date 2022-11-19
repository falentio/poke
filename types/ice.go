package types

func init() {
	Register(IceData)
}

var IceData TypeData = TypeData{
	Name: "ice",
	Bit:  Ice,

	DoubleDamageTaken: None | Fighting | Fire | Rock | Steel,
	HalfDamageTaken:   None | Ice,
	ZeroDamageTaken:   None,
}
