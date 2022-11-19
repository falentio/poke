package types

func init() {
	Register(DragonData)
}

var DragonData TypeData = TypeData{
	Bit: Dragon,

	DoubleDamageTaken: None | Dragon | Fairy | Ice,
	HalfDamageTaken:   None | Electric | Fire | Grass | Water,
	ZeroDamageTaken:   None,
}
