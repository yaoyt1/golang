package memory

import (
	"errors"
	"session/errorCode"
	"sync"
)

//MemorySession session
type MemorySession struct {
	data    map[string]interface{}
	rwMutex sync.RWMutex
}

func (ms *MemorySession) Set(key string, value interface{}) (err error) {
	ms.rwMutex.Lock()
	defer ms.rwMutex.Unlock()
	ms.data[key] = value
	return
}

func (ms *MemorySession) Get(key string) (value interface{}, err error) {
	ms.rwMutex.RLock()
	defer ms.rwMutex.RUnlock()

	value, ok := ms.data[key]
	if ok {
		err = errors.New(errorCode.MemorySessionNotExists)
	}
	return
}
func (ms *MemorySession) Del(key string) (err error) {
	ms.rwMutex.Lock()
	defer ms.rwMutex.Unlock()
	delete(ms.data, key)
	return
}
func (ms *MemorySession) Save() (err error) {
	return
}
