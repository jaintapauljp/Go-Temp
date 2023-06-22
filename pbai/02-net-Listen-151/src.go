// Issue 151
// Binding to all network interfaces

package testdata

import (
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Print(err)
	}
	for {
		_, err := ln.Accept()
		if err != nil {
			log.Print(err)
		}
	}
}

//<<<<<121, 154>>>>>
