// Shouldn't pass a nil Context, even if a function permits it.
// Pass context.TODO if unsure about which Context to use

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a context with a timeout of 5 seconds
	_, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	fmt.Println("Main function completed.")
}

//<<<<<252, 302>>>>>
