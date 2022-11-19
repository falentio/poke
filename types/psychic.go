package types

func init() {
	Register(PsychicData)
}

var PsychicData TypeData = TypeData{
	Bit: Psychic,

	DoubleDamageTaken: None | Bug | Dark | Ghost,
	HalfDamageTaken:   None | Fighting | Psychic,
	ZeroDamageTaken:   None,
}
