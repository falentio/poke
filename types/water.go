package types

func init() {
	Register(WaterData)
}

var WaterData TypeData = TypeData{
	Bit: Water,

	DoubleDamageTaken: Electric | Grass,
	HalfDamageTaken:   Fire | Ice | Steel | Water,
	ZeroDamageTaken:   None,
}
