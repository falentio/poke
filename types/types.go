package types

type (
	Type uint8
)

type TypeData struct {
	Type Type

	DoubleDamageTaken []Type
	HalfDamageTaken   []Type
	ZeroDamageTaken   []Type
}

var typesData []TypeData

func Register(t TypeData) {
	typesData = append(make([]TypeData, t.Type), typesData...)
	typesData[t.Type] = t
}

func Get(t Type) TypeData {
	return TypeData{}
}
