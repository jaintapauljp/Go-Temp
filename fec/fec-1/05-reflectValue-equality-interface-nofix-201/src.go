// Issue 201
// Comparing reflect.Value directly is almost certainly not correct,
// as it compares the reflect package's internal representation, not the underlying value.

package testdata

import (
	"fmt"
	"reflect"
)

func main() {
	v1 := reflect.ValueOf(10)
	v2 := reflect.ValueOf(30)
	if v1.Interface() == v2.Interface() {
		fmt.Print("dkfj")
	}
}

//<<<<<294, 326>>>>>
