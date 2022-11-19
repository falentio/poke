package types

func init() {
	Register(FightingData)
}

var FightingData TypeData = TypeData{
	Bit: Fighting,

	DoubleDamageTaken: Fairy | Flying | Psychic,
	HalfDamageTaken:   Bug | Dark | Rock,
	ZeroDamageTaken:   None,
}
