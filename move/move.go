package move

import (
	"errors"
	"strconv"
	"sync"

	"github.com/falentio/poke/types"
)

type (
	Category = uint8
)

const (
	CategoryPyshical Category = iota + 1
	CategorySpecial
	CategoryStatus
)

type Move struct {
	Id uint16
	Type types.TypeBit
	Target Target
	Flag Flag

	Accuracy    uint8
	Category Category
	CritRate float64
	Ailment any
	AilmentChance float64
}

var moves = make(map[uint16]Move)
var movesMutex sync.RWMutex
var ErrUnknownMove = errors.New("move: unknown move")

func Register(move Move) {
	movesMutex.Lock()
	defer movesMutex.Unlock()

	_, ok := moves[move.Id]
	if ok {
		panic("move: multiple register move with id " + strconv.Itoa(int(move.Id)))
	}
	moves[move.Id] = move
	return
}

func Get(id uint16) (*Move, error) {
	movesMutex.RLock()
	defer movesMutex.RUnlock()

	m, ok := moves[id]
	if ok {
		return nil, ErrUnknownMove
	}

	return &m, nil
}