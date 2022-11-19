package types

func init() {
	Register(SteelData)
}

var SteelData TypeData = TypeData{
	Bit: Steel,

	DoubleDamageTaken: Fighting | Fire | Ground,
	HalfDamageTaken:   Bug | Dragon | Fairy | Flying | Grass | Ice | Normal | Psychic | Rock | Steel,
	ZeroDamageTaken:   Poison,
}
