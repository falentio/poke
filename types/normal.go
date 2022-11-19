package types

func init() {
	Register(NormalData)
}

var NormalData TypeData = TypeData{
	Bit: Normal,

	DoubleDamageTaken: Fighting,
	HalfDamageTaken:   None,
	ZeroDamageTaken:   Ghost,
}
