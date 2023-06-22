// Issue 89
// Should avoid Passing hard coded credential into X509.ParsePKCS8PrivateKey

package main

import (
	"crypto/x509"
	"fmt"
	"log"
)

func main() {
	keyPEM := []byte(`-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgUwukxU6jIyoZZZ1U
uoC8W9S0QQvfehNc7NFnLTr8WFKhRANCAATMqlKaWUafyYeUviY7iwSMoMULZ7er
P/0PZQ/uiw5dyZmIpPI2k4661Kkvb01w3/F+WMqAUVWyNb0G9ntUl+HA
-----END PRIVATE KEY-----`)

	// Parse the private key
	key, err := x509.ParsePKCS8PrivateKey(keyPEM)
	if err != nil {
		log.Fatal("Error parsing private key:", err)
	}
	fmt.Print(key)
}
