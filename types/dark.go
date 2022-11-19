package types

func init() {
	Register(DarkData)
}

var DarkData TypeData = TypeData{
	Bit: Dark,

	DoubleDamageTaken: Bug | Fairy | Fighting,
	HalfDamageTaken:   Dark | Ghost,
	ZeroDamageTaken:   Psychic,
}
