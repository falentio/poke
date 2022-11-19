package types

func init() {
	Register(DragonData)
}

var DragonData TypeData = TypeData{
	Bit: Dragon,

	DoubleDamageTaken: Dragon | Fairy | Ice,
	HalfDamageTaken:   Electric | Fire | Grass | Water,
	ZeroDamageTaken:   None,
}
