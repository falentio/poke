package types

func init() {
	Register(SteelData)
}

var SteelData TypeData = TypeData{
	Bit: Steel,

	DoubleDamageTaken: None | Fighting | Fire | Ground,
	HalfDamageTaken:   None | Bug | Dragon | Fairy | Flying | Grass | Ice | Normal | Psychic | Rock | Steel,
	ZeroDamageTaken:   None | Poison,
}
