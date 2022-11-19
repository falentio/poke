package types

func init() {
	Register(NormalData)
}

var NormalData TypeData = TypeData{
	Bit: Normal,

	DoubleDamageTaken: None | Fighting,
	HalfDamageTaken:   None,
	ZeroDamageTaken:   None | Ghost,
}
