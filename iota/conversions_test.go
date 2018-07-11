package iota

import (
	"testing"

	"encoding/hex"

	"github.com/iotaledger/giota"
)

type tryteConversion struct {
	b []byte
	s string
	t giota.Trytes
}

type hashAddressConversion struct {
	hash    string
	address string
}

type chunkMessageConversion struct {
	notTrytes string
	trytes    string
}

var (
	caseOneTrytes, _   = giota.ToTrytes("IC")
	caseTwoTrytes, _   = giota.ToTrytes("HDWCXCGDEAXCGDEAPCEAHDTCGDHD")
	caseThreeTrytes, _ = giota.ToTrytes("QBCD9DPCBDVCEAXCGDEAHDWCTCEAQCTCGDHDEA9DPCBDVCFA")
	stringConvCases    = []tryteConversion{
		{b: []byte("Z"), s: "Z", t: caseOneTrytes},
		{b: []byte("this is a test"), s: "this is a test", t: caseTwoTrytes},
		{b: []byte("Golang is the best lang!"), s: "Golang is the best lang!",
			t: caseThreeTrytes},
		{b: []byte(""), s: "", t: ""},
	}
)

var hashAddressConvCases = []hashAddressConversion{
	{hash: "5804c3157e3de4e4a8b1f2417d8c61454e368883ec05e32f234386690e7c9696",
		address: "GCD9FGU9RDGBLHLHFFOFZHKBQDEEPCOBXB9BAEWDTHE9KHTAHAMBZDXCN9PDOEOE99999999999999999"},
	{hash: "080779a63f5822c2606bfdd2801b5c4429918efcecffbaa34c2daadd51bc5748",
		address: "H9G9MDDFIBGCGAEGOCZCJIUGTD9AKCNBNAJEGEIITHLIXFAFVBRAHFEH9CZFFCRB99999999999999999"},
	{hash: "d0199d3bd44c9301299de4d9d7054adb9c7fa11ac175cdee302794130b081681",
		address: "SGY9VEEBWGVBLEA9NAVELHAHZGE9TBCHUESDZEZ9DGIDPGVHUALAMES9K9H9V9UD99999999999999999"},
	{hash: "e512f80fa0e0c2872e0e29e621c40cf1693e112e020a708a619e7b87d421bf9c",
		address: "MHR9EIO9YEHHEG9ESAN9NANHFAGGL9YHXCHBQ9SAB9J9DDCEPCWEOD9EWGFABGUE99999999999999999"},
	{hash: "cca31d69bcddfdd0ecd53d98c3daeca17ed61e04bf456ebd56b9ddbaf660091a",
		address: "OGAFBAXCZFEHJISGTHXGGBQEFGBHTHZERDYGCAD9BGOBBD9GECWFEHXFCIOCI9Z999999999999999999"},
}

var chunkMessageConvCases = []chunkMessageConversion{
	{notTrytes: "5804c3157e3de4e4a8b1f2417d8c61454e368883ec05e32f234386690e7c9696",
		trytes: "GCD9FGU9RDGBLHLHFFOFZHKBQDEEPCOBXB9BAEWDTHE9KHTAHAMBZDXCN9PDOEOE"},
	{notTrytes: "080779a63f5822c2606bfdd2801b5c4429918efcecffbaa34c2daadd51bc5748",
		trytes: "H9G9MDDFIBGCGAEGOCZCJIUGTD9AKCNBNAJEGEIITHLIXFAFVBRAHFEH9CZFFCRB"},
	{notTrytes: "d0199d3bd44c9301299de4d9d7054adb9c7fa11ac175cdee302794130b081681",
		trytes: "SGY9VEEBWGVBLEA9NAVELHAHZGE9TBCHUESDZEZ9DGIDPGVHUALAMES9K9H9V9UD"},
	{notTrytes: "e512f80fa0e0c2872e0e29e621c40cf1693e112e020a708a619e7b87d421bf9c",
		trytes: "MHR9EIO9YEHHEG9ESAN9NANHFAGGL9YHXCHBQ9SAB9J9DDCEPCWEOD9EWGFABGUE"},
	{notTrytes: "cca31d69bcddfdd0ecd53d98c3daeca17ed61e04bf456ebd56b9ddbaf660091a",
		trytes: "OGAFBAXCZFEHJISGTHXGGBQEFGBHTHZERDYGCAD9BGOBBD9GECWFEHXFCIOCI9Z9"},
}

func Test_BytesToTrytes(t *testing.T) {
	for _, tc := range stringConvCases {
		result := BytesToTrytes([]byte(tc.b))
		if result != tc.t {
			t.Fatalf("BytesToTrytes(%q) should be %#v but returned %s",
				tc.b, tc.t, result)
		}
	}
}

func Test_TrytesToBytes(t *testing.T) {
	for _, tc := range stringConvCases {
		if string(TrytesToBytes(tc.t)) != string(tc.b) {
			t.Fatalf("TrytesToBytes(%q) should be %#v but returned %s",
				tc.t, tc.b, TrytesToBytes(tc.t))
		}
	}
}

func Test_TrytesToASCIITrimmed(t *testing.T) {
	for _, tc := range stringConvCases {
		result, _ := TrytesToASCIITrimmed(string(tc.t))
		if result != string(tc.s) {
			t.Fatalf("TrytesToASCIITrimmed(%q) should be %#v but returned %s",
				tc.t, tc.s, result)
		}
	}
}

func Test_ASCIIToTrytes(t *testing.T) {
	for _, tc := range stringConvCases {
		result, _ := ASCIIToTrytes(tc.s)
		if result != string(tc.t) {
			t.Fatalf("ASCIIToTrytes(%q) should be %#v but returned %s",
				tc.s, tc.t, result)
		}
	}
}

func Test_MakeAddress(t *testing.T) {
	for _, tc := range hashAddressConvCases {
		result := MakeAddress(tc.hash)
		if result != string(tc.address) {
			t.Fatalf("MakeAddress(%q) should be %#v but returned %s",
				tc.hash, tc.address, result)
		}
	}
}

func Test_ChunkMessageToTrytesWithStopper(t *testing.T) {
	for _, tc := range chunkMessageConvCases {
		bytes, err := hex.DecodeString(tc.notTrytes)
		if err != nil {
			t.Fatalf("Error encountered in test of ChunkMessageToTrytesWithStopper(): %s", err.Error())
		}
		tryteConvertedResult, err := ChunkMessageToTrytesWithStopper(string(bytes))
		if err != nil {
			t.Fatalf("ChunkMessageToTrytesWithStopper() with a non-tryte value "+
				"should not cause an error but returned %s", err.Error())
		}

		tryteResult, err := ChunkMessageToTrytesWithStopper(tc.trytes)
		if err != nil {
			t.Fatalf("ChunkMessageToTrytesWithStopper() with a tryte value "+
				"should not cause an error but returned %s", err.Error())
		}

		if string(tryteConvertedResult) != string(tc.trytes)+"A" {
			t.Fatalf("ChunkMessageToTrytesWithStopper() with a non-tryte value "+
				"should be %s but got %s", string(tc.trytes)+"A", string(tryteConvertedResult))
		}
		if string(tryteResult) != string(tc.trytes) {
			t.Fatalf("ChunkMessageToTrytesWithStopper() with a tryte value "+
				"should be %s but got %s", string(tc.trytes), string(tryteResult))
		}
	}
}

func Test_RunesToTrytes(t *testing.T) {

}
