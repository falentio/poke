package types

func init() {
	Register(NormalData)
}

var NormalData TypeData = TypeData{
	Name: "normal",
	Bit:  Normal,

	DoubleDamageTaken: None | Fighting,
	HalfDamageTaken:   None,
	ZeroDamageTaken:   None | Ghost,
}
