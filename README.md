# kswitch

Small killswitch library.

# Example usage

```go
package main

import (
	"fmt"
	"math/rand"

	ks "github.com/lateralusd/kswitch"
)

func checkSomething() bool {
	num := rand.Intn(10)
	if num == 5 {
		return true
	}
	return false
}

func main() {
	k := ks.NewKillSwitch(1)
	k.Run(checkSomething, func() {
		fmt.Println("Function returned true. Sending mail to the news company")
	})
}
```