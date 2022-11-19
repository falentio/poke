package types

func init() {
	Register(ElectricData)
}

var ElectricData TypeData = TypeData{
	Bit: Electric,

	DoubleDamageTaken: None | Ground,
	HalfDamageTaken:   None | Electric | Flying | Steel,
	ZeroDamageTaken:   None,
}
