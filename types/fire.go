package types

func init() {
	Register(FireData)
}

var FireData TypeData = TypeData{
	Bit: Fire,

	DoubleDamageTaken: None | Ground | Rock | Water,
	HalfDamageTaken:   None | Bug | Fairy | Fire | Grass | Ice | Steel,
	ZeroDamageTaken:   None,
}
