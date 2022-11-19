package types

func init() {
	Register(GrassData)
}

var GrassData TypeData = TypeData{
	Bit: Grass,

	DoubleDamageTaken: None | Bug | Fire | Flying | Ice | Poison,
	HalfDamageTaken:   None | Electric | Grass | Ground | Water,
	ZeroDamageTaken:   None,
}
