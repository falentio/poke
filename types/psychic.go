package types

func init() {
	Register(PsychicData)
}

var PsychicData TypeData = TypeData{
	Name: "psychic",
	Bit:  Psychic,

	DoubleDamageTaken: None | Bug | Dark | Ghost,
	HalfDamageTaken:   None | Fighting | Psychic,
	ZeroDamageTaken:   None,
}
