package models

import (
	Base "Cacher/pkg/base"
	"container/list"
	"fmt"
)

// The simple last in first out queue like Eviction policy.
type LifoPolicy struct {
	Base.EvictionPolicy
	evictionStack *list.List
}

// Gets a key if it exists inside the LIFO cache.
func (lifoCache *LifoPolicy) Get(key interface{}) interface{} {
	lifoCache.Mutex.Lock()
	defer lifoCache.Mutex.Unlock()

	if element, exists := lifoCache.Cache[key]; exists {
		return element.Value
	}

	return nil
}

// Puts a new key inside the LIFO cache.
func (lifoCache *LifoPolicy) Put(key, value interface{}) {
	lifoCache.Mutex.Lock()
	defer lifoCache.Mutex.Unlock()

	if _, exists := lifoCache.Cache[key]; exists {
		lifoCache.Cache[key].Value = value
	} else {
		// Cache full, hence delete the oldest element
		// (Will be in the last of the queue).
		if len(lifoCache.Cache) >= lifoCache.MaxSize {
			oldest := lifoCache.evictionStack.Front()
			if oldest != nil {
				delete(lifoCache.Cache, oldest.Value)
				lifoCache.evictionStack.Remove(oldest)
			}
		}

		lifoCache.evictionStack.PushFront(key)
		var elementCopy *list.Element = &list.Element{
			Value: value,
		}

		lifoCache.Cache[key] = elementCopy
	}
}

// Deletes a key if it exists inside the LIFO cache.
func (lifoCache *LifoPolicy) Delete(key interface{}) {
	lifoCache.Mutex.Lock()
	defer lifoCache.Mutex.Unlock()

	if element, exists := lifoCache.Cache[key]; exists {
		delete(lifoCache.Cache, key)
		lifoCache.evictionStack.Remove(element)

		fmt.Println("Successfully deleted entry with key: ", key)
	}
}

// Returns a new instance of a LifoPolicy
func LifoEvictionPolicy(maxSize int) *LifoPolicy {
	return &LifoPolicy{
		evictionStack: list.New(),
		EvictionPolicy: Base.EvictionPolicy{
			MaxSize: maxSize,
			Cache:   make(map[interface{}]*list.Element),
		},
	}
}
