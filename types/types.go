package types

import (
	"errors"
	"math"
	"math/bits"
	"sort"
)

var (
	ErrInvalidDealerType   = errors.New("types: invalid dealer type")
	ErrInvalidReceiverType = errors.New("types: invalid receiver type")
	ErrZeroType            = errors.New("types: None received")
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
	if bits.OnesCount32(uint32(dealer)) != 1 {
		return -1, ErrInvalidDealerType
	}
	if bits.OnesCount32(uint32(receiver)) > 2 {
		return -1, ErrInvalidReceiverType
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
	td.DoubleDamageTaken = (td.DoubleDamageTaken|td.ZeroDamageTaken)^td.ZeroDamageTaken
	td.HalfDamageTaken = (td.HalfDamageTaken|td.ZeroDamageTaken)^td.ZeroDamageTaken
	s := td.DoubleDamageTaken&td.HalfDamageTaken
	td.DoubleDamageTaken = (td.DoubleDamageTaken|s)^s
	td.HalfDamageTaken = (td.HalfDamageTaken|s)^s
	return td
}

func (tb TypeBit) GetData() *TypeData {
	return Get(tb)
}

func (tb TypeBit) Single() bool {
	return bits.OnesCount(uint(tb)) == 0
}
