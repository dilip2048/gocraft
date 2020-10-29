package routine

import (
"fmt"
"sync"
)

func routine() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go printHelloWorld(&wg, i)
	}
	wg.Wait()
}

func printHelloWorld(wg *sync.WaitGroup, flag int) {
	defer wg.Done()
	fmt.Printf("%v: Hello World\n", flag)
}
