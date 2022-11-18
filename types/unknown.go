package types

func init() {
	Register(UnknownData)
}

var UnknownData TypeData = TypeData{
	Type: Unknown,

	DoubleDamageTaken: []Type{},
	HalfDamageTaken:   []Type{},
	ZeroDamageTaken:   []Type{},
}
