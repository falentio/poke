package types

func init() {
	Register(RockData)
}

var RockData TypeData = TypeData{
	Bit: Rock,

	DoubleDamageTaken: Fighting | Grass | Ground | Steel | Water,
	HalfDamageTaken:   Fire | Flying | Normal | Poison,
	ZeroDamageTaken:   None,
}
