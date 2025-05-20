package MutexMap

import (
	"fmt"
	"sync"
)

type MutexMap struct {
	storage map[string]float64
	mu sync.RWMutex
}

func NewStorage(initStorage map[string]float64) *MutexMap {
	if initStorage != nil {
		return &MutexMap{
			storage : initStorage,
		}
	}
	return &MutexMap{
		storage: make(map[string]float64),
	}
}

func (m *MutexMap) GetValue(key string) float64 {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.storage[key]
}

func (m *MutexMap) IncreaseValue(key string, value float64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.storage[key] += value
}

func (m *MutexMap) GetKeys() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	keys := make([]string, 0, len(m.storage))
	for k := range m.storage {
		keys = append(keys, k)
	}
	return keys
}

func (m *MutexMap) Print() {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.storage {
		fmt.Printf("%s:%v\n", k, v)
	}
}
