package types

func init() {
	Register(DarkData)
}

var DarkData TypeData = TypeData{
	Name: "dark",
	Bit:  Dark,

	DoubleDamageTaken: None | Bug | Fairy | Fighting,
	HalfDamageTaken:   None | Dark | Ghost,
	ZeroDamageTaken:   None | Psychic,
}
