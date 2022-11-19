package types

func init() {
	Register(FairyData)
}

var FairyData TypeData = TypeData{
	Bit: Fairy,

	DoubleDamageTaken: Poison | Steel,
	HalfDamageTaken:   Bug | Dark | Fighting,
	ZeroDamageTaken:   Dragon,
}
