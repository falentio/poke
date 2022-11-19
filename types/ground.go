package types

func init() {
	Register(GroundData)
}

var GroundData TypeData = TypeData{
	Bit: Ground,

	DoubleDamageTaken: Grass | Ice | Water,
	HalfDamageTaken:   Poison | Rock,
	ZeroDamageTaken:   Electric,
}
