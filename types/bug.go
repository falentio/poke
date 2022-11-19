package types

func init() {
	Register(BugData)
}

var BugData TypeData = TypeData{
	Name: "bug",
	Bit:  Bug,

	DoubleDamageTaken: None | Fire | Flying | Rock,
	HalfDamageTaken:   None | Fighting | Grass | Ground,
	ZeroDamageTaken:   None,
}
