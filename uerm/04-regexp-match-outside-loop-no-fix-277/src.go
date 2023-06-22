// Issue 277
// Using regexp.MatchReader outside a loop is ok

package testdata

import (
	"fmt"
	"regexp"
)

func main() {
	matched, err := regexp.MatchReader("Hello", nil)
	fmt.Println(matched, err)
}

//<<<<<179, 218>>>>>
