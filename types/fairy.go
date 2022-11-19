package types

func init() {
	Register(FairyData)
}

var FairyData TypeData = TypeData{
	Bit: Fairy,

	DoubleDamageTaken: None | Poison | Steel,
	HalfDamageTaken:   None | Bug | Dark | Fighting,
	ZeroDamageTaken:   None | Dragon,
}
