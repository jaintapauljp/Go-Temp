// Issue 199
// comparison of func type param with nil is ok

package main

func F(p func()) bool {
	return p == nil
}
func main() {
	F(nil)
}

//<<<<<103, 111>>>>>
