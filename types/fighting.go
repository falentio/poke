package types

func init() {
	Register(FightingData)
}

var FightingData TypeData = TypeData{
	Bit: Fighting,

	DoubleDamageTaken: None | Fairy | Flying | Psychic,
	HalfDamageTaken:   None | Bug | Dark | Rock,
	ZeroDamageTaken:   None,
}
