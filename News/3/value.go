package main

import (
	"fmt"
	"sync"
	"time"
)
type value struct {
	sync.Mutex                       // <-- Added a mutex
	counters map[string]int
}
// START OMIT
func (v value) inc(name string) {
	v.Lock()                         // <-- Added locking of the mutex
	defer v.Unlock()
	v.counters[name]++
}
// END OMIT
func main() {
	v := value{counters: map[string]int{"a": 0, "b": 0}}
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			v.inc(name)
		}
	}
	go doIncrement("a", 100000)
	go doIncrement("a", 100000)
	// Wait a bit for the goroutines to finish
	time.Sleep(300 * time.Millisecond)
	fmt.Println(v.counters)
}