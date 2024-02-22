package interfaces

// The Cache interface that implements
// basic methods on a cache.
type Cache interface {
	Get(key interface{}) interface{}
	Put(key, value interface{})
	Delete(key interface{})
}
