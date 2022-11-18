package types

func init() {
	Register(ShadowData)
}

var ShadowData TypeData = TypeData{
	Type: Shadow,

	DoubleDamageTaken: []Type{},
	HalfDamageTaken:   []Type{},
	ZeroDamageTaken:   []Type{},
}
