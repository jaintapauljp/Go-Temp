// Issue 143
// Insecure cipher suites in CipherSuites key.
// Generate warning.

package main

import (
	"crypto/tls"
	"fmt"
)

func main() {
	a, config := "kfjdkf", &tls.Config{
		CipherSuites: []uint16{
			tls.TLS_RSA_WITH_RC4_128_SHA,
		},
	}
	fmt.Print(a, config)
}

//<<<<<168, 246>>>>>
