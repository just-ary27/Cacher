package models

import (
	Base "Cacher/pkg/base"
	"container/list"
	"fmt"
	"time"
)

// The simple first in first out queue like Eviction policy.
type FifoPolicy struct {
	Base.EvictionPolicy
	evictionQueue *list.List
}

// Gets a key if it exists inside the FIFO cache.
func (fifoCache *FifoPolicy) Get(key interface{}) interface{} {
	fifoCache.Mutex.Lock()

	// Defer here to prevent forgetting later to off the mutex.
	defer fifoCache.Mutex.Unlock()

	fmt.Printf("Acquired Mutex - Get: Key=%v\n", key)
	defer fmt.Printf("Released Mutex - Get: Key=%v\n\n", key)
	defer fmt.Println(time.Now())

	// Critical section
	if element, exists := fifoCache.Cache[key]; exists {
		return element.Value
	}

	return nil
}

// Puts a new key inside the FIFO cache.
func (fifoCache *FifoPolicy) Put(key, value interface{}) {
	fifoCache.Mutex.Lock()
	defer fifoCache.Mutex.Unlock()

	fmt.Printf("Acquired Mutex - Put: Key=%v, Value=%v\n", key, value)
	defer fmt.Printf("Released Mutex - Put: Key=%v, Value=%v\n\n", key, value)
	defer fmt.Println(time.Now())

	if _, exists := fifoCache.Cache[key]; exists {
		fifoCache.Cache[key].Value = value
	} else {
		// Cache full, hence delete the oldest element
		// (Will be in the last of the queue).
		if len(fifoCache.Cache) >= fifoCache.MaxSize {
			oldest := fifoCache.evictionQueue.Back()
			if oldest != nil {
				delete(fifoCache.Cache, oldest.Value)
				fifoCache.evictionQueue.Remove(oldest)
			}
		}

		fifoCache.evictionQueue.PushFront(key)
		var elementCopy *list.Element = &list.Element{
			Value: value,
		}

		fifoCache.Cache[key] = elementCopy
	}
}

// Deletes a key if it exists inside the FIFO cache.
func (fifoCache *FifoPolicy) Delete(key interface{}) {
	fifoCache.Mutex.Lock()
	defer fifoCache.Mutex.Unlock()

	fmt.Printf("Acquired Mutex - Delete: Key=%v\n", key)
	defer fmt.Printf("Released Mutex - Delete: Key=%v\n\n", key)
	defer fmt.Println(time.Now())

	if element, exists := fifoCache.Cache[key]; exists {
		delete(fifoCache.Cache, key)
		fifoCache.evictionQueue.Remove(element)

		fmt.Println("Successfully deleted entry with key: ", key)
	}
}

// Returns a new instance of a FifoPolicy
func FifoEvictionPolicy(maxSize int) *FifoPolicy {
	return &FifoPolicy{
		evictionQueue: list.New(),
		EvictionPolicy: Base.EvictionPolicy{
			MaxSize: maxSize,
			Cache:   make(map[interface{}]*list.Element),
		},
	}
}
