package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// Base64 is 26 lowercase + 26 uppercase + 10 digits +  ['+', '/']
	// URL safe Base64 is same but last two char as ['-', '_']

	data := "abc123!?$*&()'-=@~"
	// Go supports both standard and URL-compatible base64.
	// Here’s how to encode using the standard encoder.
	// The encoder requires a []byte so we convert our string to that type.
	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encoded) // YWJjMTIzIT8kKiYoKSctPUB+
	url_encoded := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(url_encoded) // YWJjMTIzIT8kKiYoKSctPUB-

	byte_decoded, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Println(string(byte_decoded))
	// abc123!?$*&()'-=@~

	byte_url_decoded, _ := base64.URLEncoding.DecodeString(url_encoded)
	fmt.Println(string(byte_url_decoded))
	// abc123!?$*&()'-=@~
}
