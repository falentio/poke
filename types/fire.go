package types

func init() {
	Register(FireData)
}

var FireData TypeData = TypeData{
	Name: "fire",
	Bit:  Fire,

	DoubleDamageTaken: None | Ground | Rock | Water,
	HalfDamageTaken:   None | Bug | Fairy | Fire | Grass | Ice | Steel,
	ZeroDamageTaken:   None,
}
