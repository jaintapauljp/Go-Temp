// Issue 149
// Secure implementation of HostKeyCallback
// No fix

package testdata

import (
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"
)

func insecureIgnoreHostKey() {
	_ = &ssh.ClientConfig{
		User: "username",
		Auth: []ssh.AuthMethod{nil},
		HostKeyCallback: ssh.HostKeyCallback(
			func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				if ssh.FingerprintSHA256(key) == ssh.FingerprintSHA256(key) {
					return nil
				}
				return fmt.Errorf("host key mismatch for %s", hostname)
			}),
	}
}

//<<<<<177, 524>>>>>
