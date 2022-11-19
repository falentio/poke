package types

func init() {
	Register(WaterData)
}

var WaterData TypeData = TypeData{
	Bit: Water,

	DoubleDamageTaken: None | Electric | Grass,
	HalfDamageTaken:   None | Fire | Ice | Steel | Water,
	ZeroDamageTaken:   None,
}
