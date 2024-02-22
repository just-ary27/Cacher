package examples

import (
	interfaces "Cacher/pkg/base"
	"Cacher/pkg/models"
	"fmt"
)

/*
A simple example showing lifo implementation.
*/
func LifoExample() {
	fmt.Println("\n-----Lifo Example-----")

	var lifoCache interfaces.Cache = models.LifoEvictionPolicy(4)

	lifoCache.Put("1", 1)
	lifoCache.Put("2", 2)
	lifoCache.Put("3", 3)
	lifoCache.Put("4", 4)

	fmt.Println(lifoCache)

	println("\nExisting value will now be replaced.\n")
	lifoCache.Put("4", 9)
	fmt.Println(lifoCache)

	println("\nLatest value will now be replaced.\n")
	lifoCache.Put("7", 5)
	fmt.Println(lifoCache)

	println("\nKey value will now be deleted.\n")
	lifoCache.Delete("7")
	fmt.Println(lifoCache)
}
