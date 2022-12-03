// SHA256 hashes are frequently used to compute short identities for
// binary or text blobs. For example, TLS/SSL certificates use SHA256
// to compute a certificate’s signature.
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	message := "SHA256 is an algorithm"
	hash_func := sha256.New()
	// hash function takes byte array as input
	hash_func.Write([]byte(message))

	// This gets the finalized hash result as a byte slice.
	// The argument to Sum can be used to append to an existing byte slice:
	// it usually isn’t needed
	byte_hash := hash_func.Sum(nil)
	fmt.Println(byte_hash)
	// [240 235 61 213 ....... 235 34 162 187 99 105]
	fmt.Printf("%x\n", byte_hash)
	// %x	base 16, with lower-case letters for a-f

}
