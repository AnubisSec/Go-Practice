package main

import (
	"os"
	"fmt"
)

func main() {
	// range stores, in order, the index and the value of a specific range
	// Here i == os.Args[0:] and arg == os.Args[:1]
	// So the result of "./echo3 Hello World" will be: 0: ./echo3 \n 1: Hello \n 2: World
	for i, arg := range os.Args {
		fmt.Printf("%d: %s\n", i, arg)

}

}
