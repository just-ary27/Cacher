package main

import (
	"Cacher/examples"
	cacher "Cacher/pkg"
)

func main() {
	cacher.Cacher()

	// All cache implementation examples.

	examples.FifoExample()

	examples.LifoExample()

	examples.LruExample()
}
