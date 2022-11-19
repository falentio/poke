package types

func init() {
	Register(FightingData)
}

var FightingData TypeData = TypeData{
	Name: "fighting",
	Bit:  Fighting,

	DoubleDamageTaken: None | Fairy | Flying | Psychic,
	HalfDamageTaken:   None | Bug | Dark | Rock,
	ZeroDamageTaken:   None,
}
