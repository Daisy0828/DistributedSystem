package main
import (
    "fmt"
    "sync" // TODO: Use sync with WaitGroup and Mutex!
)

type SafeValue struct{
    v int
    m sync.Mutex
}

var x = SafeValue{v:0}
var wg sync.WaitGroup

func increment() {
    x.m.Lock()
	x.v = x.v + 5
	x.m.Unlock()
    wg.Done()
}

func main() {
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go increment()
    }
    wg.Wait()
    fmt.Println("final x", x.v)
}