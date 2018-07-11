package oysterDM

import (
	"testing"
)

var sha256Hash = "5dcf66247e14fef1934ab777db11a2625f46ad1f3c2a7c93cc215423fd79224e"
var expectedObfuscatedHash = "1c726bfc54a97d5a5144eeb301051a0bc7f57015eb66d574ebcab682ea7c77f4246ded0c9d08a064a43e9069cbad6543"
var expectedNextHash = "56f2d94d7f979bf36b811ee42bb5e6eeb7cd0e0b72da96b77b33cf50f8e01093"

var expectedHashes = []string{
	"5dcf66247e14fef1934ab777db11a2625f46ad1f3c2a7c93cc215423fd79224e",
	"56f2d94d7f979bf36b811ee42bb5e6eeb7cd0e0b72da96b77b33cf50f8e01093",
	"8d7541c45f90c29d8ccbaec6e1e1c086581ec0edd218392925ae559db34bebd4",
	"920efa02eaa9da7f4230476a435f766da1387c7414df76cfac2294578197aa08",
}

var expectedAddresses = []string{
	"AAFDZCIICCGFQDIC9CNBVHQFA9E9Z9K9JGBIDDU9SHUCXGHDSHMGTFVDRHPDKDAIIAADUHL9VEH9YESCB",
	"EGLESCY99ILCVDBFUAODLHFGZ9EEWCZCJITAQDGGJFZCLANELGQHFI9BDBXBKHJFZCCEGICE9EIHLHT99",
	"K9HCNFXAGEB9PFBDDFO9NFOCIHV9TCOEYEFDCBLAGFMBYFR9PGNBKIX9SELDZAY9EIIHXBHAJ9EG9EVCM",
	"XGA9PDPFUFTAKGM9HHSAQENCEGPGRDIFR9AAL9CCLDU9UAIGGILA9BCBCCHAQ9NAGDBGYGX9RGU9TFSDR",
}

func Test_GenerateDataMapHashes(t *testing.T) {
	dataMapHashes, err := GenerateDataMapHashes(sha256Hash, len(expectedHashes))
	if err != nil {
		t.Fatalf("GenerateDataMapHashes() should not have resulted in an error but received %s",
			err.Error())
	}

	if len(dataMapHashes) != len(expectedHashes) {
		t.Fatalf("GenerateDataMapHashes() should produce %d entries but produced %d",
			len(expectedHashes), len(dataMapHashes))
	}

	for i, dataMapHash := range dataMapHashes {
		if dataMapHash != expectedHashes[i] {
			t.Fatalf("GenerateDataMapHashes() should produce %s but returned %s",
				expectedHashes[i], dataMapHash)
		}
	}
}

func Test_GenerateDataMapAddresses(t *testing.T) {

	dataMapAddresses, err := GenerateDataMapAddresses(sha256Hash, len(expectedAddresses))
	if err != nil {
		t.Fatalf("GenerateDataMapAddresses() should not have resulted in an error but received %s",
			err.Error())
	}

	if len(dataMapAddresses) != len(expectedAddresses) {
		t.Fatalf("GenerateDataMapAddresses() should produce %d entries but produced %d",
			len(expectedAddresses), len(dataMapAddresses))
	}

	for i, dataMapAddress := range dataMapAddresses {
		if dataMapAddress != expectedAddresses[i] {
			t.Fatalf("GenerateDataMapAddresses() should produce %s but returned %s",
				expectedAddresses[i], dataMapAddress)
		}
	}
}

func Test_GenerateObfuscatedHashAndNextHash(t *testing.T) {

	obfuscatedHash, nextHash, err := GenerateObfuscatedHashAndNextHash(sha256Hash)

	if obfuscatedHash != expectedObfuscatedHash {
		t.Fatalf("GenerateObfuscatedHashAndNextHash() for obfuscated hash should be %s but returned %s",
			expectedObfuscatedHash, obfuscatedHash)
	}
	if nextHash != expectedNextHash {
		t.Fatalf("GenerateObfuscatedHashAndNextHash() for next hash should be %s but returned %s",
			expectedNextHash, nextHash)
	}
	if err != nil {
		t.Fatalf("GenerateObfuscatedHashAndNextHash() should not have resulted in an error but received %s",
			err.Error())
	}
}
