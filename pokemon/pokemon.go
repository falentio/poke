package pokemon

type Pokemon struct {
	BaseStats Stat
	IV        Stat
	EV        Stat
}

type Stat struct {
	HP             uint8
	Attack         uint8
	Deffense       uint8
	SpecialAttack  uint8
	SpecialDefense uint8
	Speed          uint8
}
