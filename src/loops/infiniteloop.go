package loops

import (
	"fmt"
	"time"
)

func testInfiniteLoop() {
	for {
		fmt.Println("Tick")
		time.Sleep(1 * time.Second)
	}
}
