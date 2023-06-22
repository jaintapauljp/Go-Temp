// Passing pointer arg into gob.Unmarshal
// No fix

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Person struct {
}

func main() {
	var buffer bytes.Buffer
	var person Person
	decoder := gob.NewDecoder(&buffer)
	err := decoder.Decode(&person)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
}

//<<<<<237, 260>>>>>
