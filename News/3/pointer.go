package main

import (
	"fmt"
	"sync"
	"time"
)
type pointer struct {
	sync.Mutex                       // <-- Added a mutex
	counters map[string]int
}
// START OMIT
func (p *pointer) inc(name string) {
	p.Lock()                         // <-- Added locking of the mutex
	defer p.Unlock()
	p.counters[name]++
}
// END OMIT
func main() {
	p := pointer{counters: map[string]int{"a": 0, "b": 0}}
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			p.inc(name)
		}
	}
	go doIncrement("a", 100000)
	go doIncrement("a", 100000)
	// Wait a bit for the goroutines to finish
	time.Sleep(300 * time.Millisecond)
	fmt.Println(p.counters)
}