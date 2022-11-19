package types

func init() {
	Register(GrassData)
}

var GrassData TypeData = TypeData{
	Bit: Grass,

	DoubleDamageTaken: Bug | Fire | Flying | Ice | Poison,
	HalfDamageTaken:   Electric | Grass | Ground | Water,
	ZeroDamageTaken:   None,
}
