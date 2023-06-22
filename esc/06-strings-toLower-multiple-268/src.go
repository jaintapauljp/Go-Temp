// Issue 268
// Inefficient string comparison with strings.ToLower

package testdata

import (
	"fmt"
	"strings"
)

func isEqual(s1, s2 string) bool {
	return strings.ToLower(s1) == strings.ToLower(s2) && len(s1+s2) < 10
}
func main() {
	equal := isEqual("str1", "str2")
	if equal {
		fmt.Print("Matched")
	}
}

//<<<<<159, 201>>>>>
