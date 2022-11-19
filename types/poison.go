package types

func init() {
	Register(PoisonData)
}

var PoisonData TypeData = TypeData{
	Bit: Poison,

	DoubleDamageTaken: Ground | Psychic,
	HalfDamageTaken:   Bug | Fairy | Fighting | Grass | Poison,
	ZeroDamageTaken:   None,
}
