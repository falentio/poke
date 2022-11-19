package types

func init() {
	Register(FairyData)
}

var FairyData TypeData = TypeData{
	Name: "fairy",
	Bit:  Fairy,

	DoubleDamageTaken: None | Poison | Steel,
	HalfDamageTaken:   None | Bug | Dark | Fighting,
	ZeroDamageTaken:   None | Dragon,
}
