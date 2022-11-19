package types

func init() {
	Register(GroundData)
}

var GroundData TypeData = TypeData{
	Name: "ground",
	Bit:  Ground,

	DoubleDamageTaken: None | Grass | Ice | Water,
	HalfDamageTaken:   None | Poison | Rock,
	ZeroDamageTaken:   None | Electric,
}
