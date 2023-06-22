// Issue 268
// Inefficient string comparison with strings.ToLower

package testdata

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "str1"
	s2 := "str2"
	if strings.ToLower(s1) == strings.ToLower(s2) {
		fmt.Print("Matched")
	}
}

//<<<<<162, 204>>>>>
