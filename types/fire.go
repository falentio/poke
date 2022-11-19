package types

func init() {
	Register(FireData)
}

var FireData TypeData = TypeData{
	Bit: Fire,

	DoubleDamageTaken: Ground | Rock | Water,
	HalfDamageTaken:   Bug | Fairy | Fire | Grass | Ice | Steel,
	ZeroDamageTaken:   None,
}
