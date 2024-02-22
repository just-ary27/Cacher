package interfaces

import (
	"container/list"
	"fmt"
	"sync"
)

// The BaseEvictionPolicy struct,
// can be used to make your own custom
// cache eviction policies.
type EvictionPolicy struct {
	Cache map[interface{}]*list.Element

	// The maximum size of cache.
	MaxSize int

	// Mutex lock for thread safety.
	Mutex sync.Mutex
}

// Overriding the default string representation
func (b *EvictionPolicy) String() string {
	var result string

	for key, element := range b.Cache {
		result += fmt.Sprintf("%s: %v\n", key, element.Value)
	}

	return fmt.Sprintf("State of <EvictionPolicy@%v>:\n%v", &b, result)
}
