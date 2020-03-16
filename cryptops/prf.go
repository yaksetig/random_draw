package cryptops

import "golang.org/x/crypto/sha3"

// PRF receives the data to be hashes and produces a digest
func PRF(data []byte)([]byte){
	h := sha3.NewLegacyKeccak256()
	h.Write(data)
	return h.Sum(nil)
}
