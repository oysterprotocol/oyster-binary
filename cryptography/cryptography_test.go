package oysterCrypto

import (
	"encoding/hex"
	"testing"

	"crypto/sha256"
	"crypto/sha512"

	"golang.org/x/crypto/sha3"
)

type EncryptionTestStruct struct {
	key    string
	secret string
	nonce  string
	result string
}

var encryptionTestCases = []EncryptionTestStruct{
	{
		key:    "64dc1ce4655554f514a4ce83e08c1d08372fdf02bd8c9b6dbecfc74b783d39d1",
		secret: "0000000000000000000000000000000000000000000000000000000000000001",
		nonce:  "948791aa5dfd8f71405da35c637ad58cc9f5fec7424dba3e97630921e130c5b6",
		result: "592d93e3bb89f8835c9949460a2b0195e8ea915e724a35b3c713a6201ce94002ae94b5546647db1ffa94a3002f500897",
	},
	{
		key:    "99577b266e77d07e364d0b87bf1bcef44c78e3668dfdc3881969b375c09d4fcd",
		secret: "1004444400000006780000000000000000000000000012345000000765430001",
		nonce:  "23384a8eabc4a4ba091cfdbcb3dbacdc27000c03e318fd52accb8e2380f11320",
		result: "73fb51882b7fdd04d1f92146fed5b198e820ea08d00dd7bb65cde4f8a0b0e00cfedb93317ef05d7d149371b4b6b2c272",
	},
	{
		key:    "7fb4ca9cc0032bafc2ebd0fda018a41f5adfcf441123de22ab736a42207933f7",
		secret: "7777777774444444777777744444447777777444444777777744444777777744",
		nonce:  "0d412fa10c9027b7163302e38c96a5c0904b1b04cb55c66162296d0be2e3caa2",
		result: "8a859e9a265f28d36153c5d3849f5e1ec7574431fb1af68b0bc74d928772edcce1ae50fae6c4634bdcc876eef85679a9",
	},
}

type HashBytesFromHexStruct struct {
	input        string
	sha256Output string
	sha512Output string
	sha3Output   string
	sha384Output string
}

var hashHexTestCases = []HashBytesFromHexStruct{
	{
		input:        "D6C816D2AE89F2AB6B19CACFA932B03D328E01963880BCA29D8008EAB5F63199",
		sha256Output: "40c0dcf3029c9817537c25ffeb6ed2bb2441fcc62217dd43f8eed933c63cd434",
		sha512Output: "5524374c919e2fb7ca909bf8db6c3411f88be7fca6777a5afc70075f4a4a75ff4aefe351b7c9a37459778b079fed2c87e08d00235c46fd2ae00a22ab47d0be7a",
		sha384Output: "0f4d0104f0aa044768b0d4a113e8b0c7d17e3d997c84cfd002048ceda8940a0a22208008b2e7bfda3415d37e3179ace4",
		sha3Output:   "61ab1b77a6aa736a95b3cd56eaf0d724047e6e207f93727fe287d4a430d74667",
	},
	{
		input:        "C145341739B561697D280AB844C99D3372F2F7EC521096EC7AF9EC02E3EE8B49",
		sha256Output: "317b6622d1717ef0337b59494762aff2e6da880d50d70e4ec5fe2a0e14d55a5e",
		sha512Output: "af2b0be4fe379809788d7424c9459e5fd61c7b36cfb9b6f1646c3d1ceb8f47923a5f2124ed86cb54bf2cf5100429f6bcbdaabb173508da134a195ea7f792cbee",
		sha384Output: "ccdfd20c74086eec8bbf220526019b7acc62773118303038d2105ffdba52f990b05fe8dffabf916e34ec44f0681853dd",
		sha3Output:   "6f1e39a0c4f42e04cbcb6169dae8614948938389a6db49a96b358e51548425a5",
	},
	{
		input:        "E1B3D3E5541846ECE56B06F69C1728AC95DF84308245E7D2FD6327EBC3FB30ED",
		sha256Output: "2067c13debdb309673247701bccd68e2fef4f9bf9cfc9eb730b2449fd197aee1",
		sha512Output: "2ced527a9a0af4f848eb4685c5bd99d1ff38fc01cf8458a5b8764cfe73b646ca18204eea28cc95a9cca4c413d34753e6ed961dc5ca65ffdd3340cd97253ae4e9",
		sha384Output: "c8a276ef52c9e70ba17a780c60937ca6997ab95ad2959aa4e5af3551e48bb8d25a8281f054b9680a328fcf2f453131df",
		sha3Output:   "d5945384aa3fba8a9d57eec40d27269ffb773ff9e7ad8e3d8a60826338ef8689",
	},
}

func Test_Encrypt(t *testing.T) {
	for _, tc := range encryptionTestCases {
		result, _ := Encrypt(tc.key, tc.secret, tc.nonce)
		if hex.EncodeToString(result) != tc.result {
			t.Fatalf("Encrypt() result should be %s but returned %s",
				tc.result, hex.EncodeToString(result))
		}
	}
}

func Test_Decrypt(t *testing.T) {
	for _, tc := range encryptionTestCases {
		secret, _ := Decrypt(tc.key, tc.result, tc.nonce)
		if hex.EncodeToString(secret) != tc.secret {
			t.Fatalf("Decrypt() should be %s but returned %s",
				tc.secret, hex.EncodeToString(secret))
		}
	}
}

func Test_HashBytesFromHex(t *testing.T) {
	for _, tc := range hashHexTestCases {
		hash256, _ := HashBytesFromHex(tc.input, sha256.New())
		hash512, _ := HashBytesFromHex(tc.input, sha512.New())
		hash384, _ := HashBytesFromHex(tc.input, sha512.New384())
		hash3, _ := HashBytesFromHex(tc.input, sha3.New256())
		if hash256 != tc.sha256Output {
			t.Fatalf("HashBytesFromHex() with sha256 should be %s but returned %s",
				tc.sha256Output, hash256)
		}
		if hash512 != tc.sha512Output {
			t.Fatalf("HashBytesFromHex() with sha512 should be %s but returned %s",
				tc.sha512Output, hash512)
		}
		if hash384 != tc.sha384Output {
			t.Fatalf("HashBytesFromHex() with sha384 should be %s but returned %s",
				tc.sha384Output, hash384)
		}
		if hash3 != tc.sha3Output {
			t.Fatalf("HashBytesFromHex() with sha3 should be %s but returned %s",
				tc.sha3Output, hash3)
		}
	}
}
