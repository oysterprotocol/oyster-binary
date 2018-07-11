package oysterDM

import (
	"crypto/sha256"
	"crypto/sha512"

	"github.com/oysterprotocol/oyster-binary/cryptography"
	"github.com/oysterprotocol/oyster-binary/errors"
	"github.com/oysterprotocol/oyster-binary/iota"
)

/*GenerateDataMapAddresses accepts a hash and an integer for the number of chunks
and generates the datamap addresses*/
func GenerateDataMapAddresses(genesisHash string, numChunks int) ([]string, error) {
	var dataMapAddresses []string
	// TODO so something with error
	obfuscatedHash, nextHash, _ := GenerateObfuscatedHashAndNextHash(genesisHash)
	dataMapAddresses = append(dataMapAddresses, iota.MakeAddress(obfuscatedHash))

	currentHash := nextHash
	for i := 1; i < numChunks; i++ {
		//TODO do something with error
		obfuscatedHash, nextHash, err := GenerateObfuscatedHashAndNextHash(currentHash)
		if err != nil {
			dataMapAddresses = []string{}
			break
		}
		currentHash = nextHash
		dataMapAddresses = append(dataMapAddresses, iota.MakeAddress(obfuscatedHash))
	}
	return dataMapAddresses, nil
}

/*GenerateDataMapHashes accepts a hash and an integer for the number of chunks
and generates the datamap hashes*/
func GenerateDataMapHashes(genesisHash string, numChunks int) ([]string, error) {
	var dataMapHashes []string
	dataMapHashes = append(dataMapHashes, genesisHash)

	currentHash := genesisHash
	for i := 1; i < numChunks; i++ {
		//TODO do something with error
		_, nextHash, err := GenerateObfuscatedHashAndNextHash(currentHash)
		if err != nil {
			dataMapHashes = []string{}
			break
		}
		currentHash = nextHash
		dataMapHashes = append(dataMapHashes, nextHash)
	}
	return dataMapHashes, nil
}

/*GenerateObfuscatedHashAndNextHash accepts a hash and returns that hash obfuscated and the next
hash in the datamap*/
func GenerateObfuscatedHashAndNextHash(currentHash string) (string, string, error) {
	obfuscatedHash, sha384Err := oysterCrypto.HashBytesFromHex(currentHash, sha512.New384())
	nextHash, sha256Err := oysterCrypto.HashBytesFromHex(currentHash, sha256.New())
	err := oysterErrors.CollectErrors([]error{sha384Err, sha256Err})
	return obfuscatedHash, nextHash, err
}
