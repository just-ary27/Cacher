package examples

import (
	interfaces "Cacher/pkg/base"
	"Cacher/pkg/models"
	"fmt"
)

/*
A simple example showing lru implementation.
*/
func LruExample() {
	fmt.Println("\n-----Lru Example-----")

	var lruCache interfaces.Cache = models.LruEvictionPolicy(5)

	lruCache.Put("1", "Cacher")

	var cachedData map[int]string = map[int]string{
		1: "justary27",
	}

	lruCache.Put("2", cachedData)
	lruCache.Put("3", 3)
	lruCache.Put(3, "xyz")

	fmt.Println(lruCache)

	println("\nExisting value will now be replaced.\n")
	lruCache.Put(3, 9)
	fmt.Println(lruCache)

	println("\nLeast recent value will now be replaced.\n")
	lruCache.Put("7", 5)
	fmt.Println(lruCache)

	println("\nKey value will now be deleted.\n")
	lruCache.Delete(3)
	fmt.Println(lruCache)
}
