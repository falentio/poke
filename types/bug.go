package types

func init() {
	Register(BugData)
}

var BugData TypeData = TypeData{
	Bit: Bug,

	DoubleDamageTaken: Fire | Flying | Rock,
	HalfDamageTaken:   Fighting | Grass | Ground,
	ZeroDamageTaken:   None,
}
