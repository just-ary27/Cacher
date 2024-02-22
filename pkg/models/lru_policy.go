package models

import (
	Base "Cacher/pkg/base"
	"container/list"
	"fmt"
)

// The simple first in first out queue like Eviction policy.
type LruPolicy struct {
	Base.EvictionPolicy
	evictionList *list.List
}

func (lruCache *LruPolicy) Get(key interface{}) interface{} {
	lruCache.Mutex.Lock()
	defer lruCache.Mutex.Unlock()

	// Critical section
	if element, exists := lruCache.Cache[key]; exists {
		// Newly accessed, hence move to front.
		lruCache.evictionList.MoveToFront(element)
		return element.Value
	}

	return nil
}

func (lruCache *LruPolicy) Put(key, value interface{}) {
	lruCache.Mutex.Lock()
	defer lruCache.Mutex.Unlock()

	if _, exists := lruCache.Cache[key]; exists {
		lruCache.Cache[key].Value = value

		// Newly accessed, hence move to front.
		lruCache.evictionList.MoveToFront(lruCache.Cache[key])
	} else {
		// Cache full, hence delete the oldest element
		// (Will be in the last of the queue).
		if len(lruCache.Cache) >= lruCache.MaxSize {
			oldest := lruCache.evictionList.Back()
			if oldest != nil {
				delete(lruCache.Cache, oldest.Value)
				lruCache.evictionList.Remove(oldest)
			}
		}

		lruCache.evictionList.PushFront(key)
		var elementCopy *list.Element = &list.Element{
			Value: value,
		}

		lruCache.Cache[key] = elementCopy
	}
}

func (lruCache *LruPolicy) Delete(key interface{}) {
	lruCache.Mutex.Lock()
	defer lruCache.Mutex.Unlock()

	if element, exists := lruCache.Cache[key]; exists {
		delete(lruCache.Cache, key)
		lruCache.evictionList.Remove(element)

		fmt.Println("Successfully deleted entry with key:", key)
	}
}

// Returns a new instance of a LifoPolicy
func LruEvictionPolicy(maxSize int) *LruPolicy {
	return &LruPolicy{
		evictionList: list.New(),
		EvictionPolicy: Base.EvictionPolicy{
			MaxSize: maxSize,
			Cache:   make(map[interface{}]*list.Element),
		},
	}
}
