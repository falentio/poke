package types

func init() {
	Register(RockData)
}

var RockData TypeData = TypeData{
	Name: "rock",
	Bit:  Rock,

	DoubleDamageTaken: None | Fighting | Grass | Ground | Steel | Water,
	HalfDamageTaken:   None | Fire | Flying | Normal | Poison,
	ZeroDamageTaken:   None,
}
