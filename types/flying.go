package types

func init() {
	Register(FlyingData)
}

var FlyingData TypeData = TypeData{
	Name: "flying",
	Bit:  Flying,

	DoubleDamageTaken: None | Electric | Ice | Rock,
	HalfDamageTaken:   None | Bug | Fighting | Grass,
	ZeroDamageTaken:   None | Ground,
}
