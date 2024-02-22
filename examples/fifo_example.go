package examples

import (
	interfaces "Cacher/pkg/base"
	"Cacher/pkg/models"
	"fmt"
	"sync"
	"time"
)

/*
A simple example showing fifo implementation
and the thread safety feature via a mutex lock.
*/
func FifoExample() {
	fmt.Println("\n-----Fifo Example-----")

	var fifoCache interfaces.Cache = models.FifoEvictionPolicy(4)

	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// Self invoking anonymous function
		// that inserts {i, i} as a key value
		// pair in the FIFO cache
		go func(index int) {
			defer wg.Done()
			key := fmt.Sprintf("%d", index)

			fifoCache.Put(key, index)
			time.Sleep(time.Millisecond * 100)

			result := fifoCache.Get(key)
			fmt.Printf("Goroutine %d: Key=%s, Value=%v\n\n", index, key, result)
		}(i)
	}

	wg.Wait()

	fmt.Println(fifoCache)

	// Not a guaranteed replacement of an existing value
	// as the order of routine execution is random,
	// hence 4 may have been the first value
	// inserted, thus may have been removed.
	println("\nExisting value will now be replaced.\n")
	fifoCache.Put("4", 9)
	fmt.Println(fifoCache)

	println("\nOldest value will now be replaced.\n")
	fifoCache.Put("7", 5)
	fmt.Println(fifoCache)

	println("\nKey value will now be deleted.\n")
	fifoCache.Delete("7")
	fmt.Println(fifoCache)
}
