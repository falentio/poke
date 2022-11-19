package types

import (
	"errors"
	"math"
	"math/bits"
	"sort"
)

var (
	ErrInvalidDealerType = errors.New("types: invalid dealer type")
	ErrZeroType          = errors.New("types: None received")
)

type (
	TypeBit uint32
)

type TypeData struct {
	Name string
	Bit  TypeBit

	DoubleDamageTaken TypeBit
	HalfDamageTaken   TypeBit
	ZeroDamageTaken   TypeBit
}

func GetMultiplier(receiver, dealer TypeBit) (float64, error) {
	if receiver == None || dealer == None {
		return -1, ErrZeroType
	}
	if bits.OnesCount16(uint16(dealer)) != 1 {
		return -1, ErrInvalidDealerType
	}

	var multiplier float64 = 1
	td := Get(receiver)
	if td.ZeroDamageTaken&dealer != 0 {
		return 0, nil
	}
	multiplier *= math.Pow(2, float64(bits.OnesCount32(uint32(td.DoubleDamageTaken&dealer))))
	multiplier /= math.Pow(2, float64(bits.OnesCount32(uint32(td.HalfDamageTaken&dealer))))
	return multiplier, nil
}

var typesData []TypeData

func Register(t TypeData) {
	typesData = append(typesData, t)
	sort.Slice(typesData, func(i, j int) bool {
		return typesData[j].Bit > typesData[i].Bit
	})
}

func Get(tb TypeBit) *TypeData {
	td := &TypeData{Bit: tb}
	for _, t := range typesData {
		if t.Bit&tb == 0 {
			continue
		}
		td.DoubleDamageTaken |= t.DoubleDamageTaken
		td.HalfDamageTaken |= t.HalfDamageTaken
		td.ZeroDamageTaken |= t.ZeroDamageTaken
	}
	td.DoubleDamageTaken = td.DoubleDamageTaken ^ (td.DoubleDamageTaken & td.HalfDamageTaken)
	td.HalfDamageTaken = td.HalfDamageTaken ^ (td.HalfDamageTaken & td.DoubleDamageTaken)
	return td
}

func GetFromString(typeNames ...string) *TypeData {
	tb := None
	for _, td := range typesData {
		for _, name := range typeNames {
			if td.Name == name {
				tb |= td.Bit
			}
		}
	}

	return Get(tb)
}
