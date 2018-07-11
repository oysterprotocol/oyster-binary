package oysterCrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"hash"

	"github.com/oysterprotocol/oyster-binary/errors"
)

/*Encrypt accepts a key, unencrypted secret, and nonce, and should return the encrypted
result as a byte array*/
func Encrypt(key string, unencryptedSecret string, nonce string) ([]byte, error) {
	nonceInBytes, secretInBytes, gcm, err := prepareNonceSecretAndGCM(key, unencryptedSecret, nonce)
	data := gcm.Seal(nil, nonceInBytes, secretInBytes, nil)
	return data, err
}

/*Decrypt accepts a key, encrypted secret, and nonce, and should return the decrypted
result as a byte array*/
func Decrypt(key string, encryptedSecret string, nonce string) ([]byte, error) {
	nonceInBytes, secretInBytes, gcm, prepErr := prepareNonceSecretAndGCM(key, encryptedSecret, nonce)
	data, openErr := gcm.Open(nil, nonceInBytes, secretInBytes, nil)
	err := oysterErrors.ReturnFirstError([]error{prepErr, openErr})
	return data, err
}

func prepareNonceSecretAndGCM(key string, secret string, nonce string) ([]byte, []byte, cipher.AEAD, error) {
	keyInBytes, keyDecodeErr := hex.DecodeString(key)
	secretInBytes, secretDecodeErr := hex.DecodeString(secret)
	block, createCipherErr := aes.NewCipher(keyInBytes)
	gcm, newGCMErr := cipher.NewGCM(block)
	nonceInBytes, nonceToBytesErr := hex.DecodeString(nonce[0 : 2*gcm.NonceSize()])
	err := oysterErrors.ReturnFirstError([]error{keyDecodeErr, secretDecodeErr, createCipherErr, newGCMErr, nonceToBytesErr})
	return nonceInBytes, secretInBytes, gcm, err
}

/*HashString receives a string and hashes it according to the hashing algorithm passed in, and should return a string*/
func HashString(str string, shaAlg hash.Hash) (string, error) {
	_, err := shaAlg.Write([]byte(str))
	hashString := hex.EncodeToString(shaAlg.Sum(nil))
	return hashString, err
}

/*HashBytesFromHex receives a hex string, converts it to bytes and hashes it according
to the hashing algorithm passed in, and should return a string*/
func HashBytesFromHex(hexStr string, shaAlg hash.Hash) (string, error) {
	input, decodeErr := hex.DecodeString(hexStr)
	_, writeErr := shaAlg.Write(input)
	hashString := hex.EncodeToString(shaAlg.Sum(nil))
	err := oysterErrors.ReturnFirstError([]error{decodeErr, writeErr})
	return hashString, err
}
