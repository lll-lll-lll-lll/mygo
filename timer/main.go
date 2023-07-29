// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker()

}

// ticker
func ticker() {
	t := time.NewTicker(2 * time.Second)
	defer t.Stop()
	done := make(chan bool)
	go func() {
		fmt.Println("goroutine 1 start")
		defer fmt.Println("goroutine 1 end")
		time.Sleep(time.Second * 4)
		t.Stop()
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("done")
			return
		// NewTickerの引数に渡した時間ごとに、t.Cは現在時刻を表し、tに代入、t.Cケースが実行される
		case t := <-t.C:
			fmt.Println("Current time: ", t)
		}
	}
}
