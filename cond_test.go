package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var loker = sync.Mutex{}
var cond = sync.NewCond(&loker)
var group = sync.WaitGroup{}

func WaitCondition(value int){
	defer group.Done()
	group.Add(1)
	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done ", value)

	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	//go func() {
	//	time.Sleep(1 * time.Second)
	//	cond.Broadcast()
	//}()

	group.Wait()
}
