package types

func init() {
	Register(FlyingData)
}

var FlyingData TypeData = TypeData{
	Bit: Flying,

	DoubleDamageTaken: Electric | Ice | Rock,
	HalfDamageTaken:   Bug | Fighting | Grass,
	ZeroDamageTaken:   Ground,
}
