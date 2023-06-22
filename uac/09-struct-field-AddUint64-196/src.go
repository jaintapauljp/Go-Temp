// Issue 196
// Direct assignment to atomic value should be avoided.
// Fix will be generated.

package main

import (
	"fmt"
	"sync/atomic"
)

type T struct {
	x uint64
}

func main() {
	t := T{x: 20}
	t.x = atomic.AddUint64(&t.x, 30)
	fmt.Println(t)
}

//<<<<<203, 235>>>>>
