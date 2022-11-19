package types

func init() {
	Register(ElectricData)
}

var ElectricData TypeData = TypeData{
	Bit: Electric,

	DoubleDamageTaken: Ground,
	HalfDamageTaken:   Electric | Flying | Steel,
	ZeroDamageTaken:   None,
}
