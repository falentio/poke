package effect

import (
	"fmt"
	"strconv"
	"sync"
)

var moveEffect = make(map[uint16]Effect)
var moveEffectMutex sync.RWMutex

func RegisterMoveEffect(id uint16, e Effect) {
	moveEffectMutex.Lock()
	defer moveEffectMutex.Unlock()

	_, ok := moveEffect[id]
	if ok {
		panic("effect: multiple register move with id " + strconv.Itoa(int(id)))
	}
	moveEffect[id] = e
	return
}

func GetMoveEffect(id uint16) (Effect, error) {
	moveEffectMutex.RLock()
	defer moveEffectMutex.RUnlock()

	e, ok := moveEffect[id]
	if !ok {
		return nil, fmt.Errorf("effect: can not find effect for move with id %q", id)
	}

	return e, nil
}

func GetOrNopMoveEffect(id uint16) Effect {
	e, err := GetMoveEffect(id)
	if err != nil {
		return Base{}
	}
	return e
}