// Issue 143
// Setting tls.Config.MinVersion to tls.VersionSSL30
// Generate warning

package main

import . "crypto/tls"

func main() {
	config := &Config{}
	config.MinVersion = VersionSSL30
}

//<<<<<160, 192>>>>>
