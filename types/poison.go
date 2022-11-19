package types

func init() {
	Register(PoisonData)
}

var PoisonData TypeData = TypeData{
	Name: "poison",
	Bit:  Poison,

	DoubleDamageTaken: None | Ground | Psychic,
	HalfDamageTaken:   None | Bug | Fairy | Fighting | Grass | Poison,
	ZeroDamageTaken:   None,
}
