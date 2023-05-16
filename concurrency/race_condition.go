package concurrency

import (
	"fmt"
	"time"
)

func TestPrint() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func RaceCondition() {
	go TestPrint()
	go TestPrint()

	time.Sleep(time.Second * 5)
}
