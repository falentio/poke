package types

func init() {
	Register(PsychicData)
}

var PsychicData TypeData = TypeData{
	Bit: Psychic,

	DoubleDamageTaken: Bug | Dark | Ghost,
	HalfDamageTaken:   Fighting | Psychic,
	ZeroDamageTaken:   None,
}
