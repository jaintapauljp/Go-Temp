// Self assignment in if statement. Warning
// should be over if statement.

package testdata

func isEven(a int) bool {
	if a = a; a%2 == 0 {
		return true
	}
	return false
}

//<<<<<125, 130>>>>>
