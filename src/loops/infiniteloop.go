package loops

import (
	"fmt"
	"time"
)

func TestInfiniteLoop() {
	for {
		fmt.Println("Tick")
		time.Sleep(1 * time.Nanosecond)
		break
	}
}
