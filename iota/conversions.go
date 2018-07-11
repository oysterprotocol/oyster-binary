package iota

import (
	"encoding/hex"
	"errors"
	"strings"

	"fmt"

	"github.com/iotaledger/giota"
)

var (
	/*TrytesAlphabet is all the characters that can be in a tryte string*/
	TrytesAlphabet = []rune("9ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

/*StopperTryte is the tryte character we will add to the end of our tryte strings*/
const StopperTryte = "A"

/*ASCIIToTrytes converts an ascii string to trytes*/
func ASCIIToTrytes(asciiString string) (string, error) {
	var b strings.Builder

	for _, character := range asciiString {
		var charCode = character

		// If not recognizable ASCII character, return null
		if charCode > 255 {

			fmt.Println(charCode)
			fmt.Println(character)

			err := errors.New("asciiString is not ASCII char in ASCIIToTrytes method")
			if err != nil {
				// TODO something with err
				fmt.Println(err)
			}
			return "", err
		}

		var firstValue = charCode % 27
		var secondValue = (charCode - firstValue) / 27
		var trytesValue = string(TrytesAlphabet[firstValue]) + string(TrytesAlphabet[secondValue])
		_, err := b.WriteString(trytesValue)
		if err != nil {
			// TODO something with err
			fmt.Println(err)
		}
	}

	return b.String(), nil
}

/*TrytesToASCIITrimmed will remove extra 9s from the end of a tryte string and will
then convert it to an ascii string*/
func TrytesToASCIITrimmed(inputTrytes string) (string, error) {
	notNineIndex := strings.LastIndexFunc(inputTrytes, func(rune rune) bool {
		return string(rune) != "9"
	})
	trimmedString := inputTrytes[0 : notNineIndex+1]

	if len(trimmedString)%2 != 0 {
		trimmedString += "9"
	}

	return TrytesToASCII(trimmedString)
}

/*TrytesToASCII converts trytes to an ascii string*/
func TrytesToASCII(inputTrytes string) (string, error) {
	// If input length is odd, return an error
	if len(inputTrytes)%2 != 0 {
		err := errors.New("method TrytesToASCII needs input with an even number of characters")
		if err != nil {
			// TODO something with err
			fmt.Println(err)
		}
		return "", err
	}

	var b strings.Builder
	for i := 0; i < len(inputTrytes); i += 2 {
		// get a trytes pair
		trytes := string(inputTrytes[i]) + string(inputTrytes[i+1])

		firstValue := strings.Index(string(TrytesAlphabet), (string(trytes[0])))
		secondValue := strings.Index(string(TrytesAlphabet), (string(trytes[1])))

		decimalValue := firstValue + secondValue*27
		character := string(decimalValue)
		_, err := b.WriteString(character)
		if err != nil {
			// TODO something with err
			fmt.Println(err)
		}
	}

	return b.String(), nil
}

/*TrytesToBytes converts trytes to a byte array*/
//TrytesToBytes and BytesToTrytes written by Chris Warner, thanks!
func TrytesToBytes(t giota.Trytes) []byte {
	var output []byte
	trytesString := string(t)
	for i := 0; i < len(trytesString); i += 2 {
		v1 := strings.IndexRune(string(TrytesAlphabet), rune(trytesString[i]))
		v2 := strings.IndexRune(string(TrytesAlphabet), rune(trytesString[i+1]))
		decimal := v1 + v2*27
		c := byte(decimal)
		output = append(output, c)
	}
	return output
}

/*BytesToTrytes accepts a byte array and converts to trytes*/
func BytesToTrytes(b []byte) giota.Trytes {
	var output string
	for _, c := range b {
		v1 := c % 27
		v2 := (c - v1) / 27
		output += string(TrytesAlphabet[v1]) + string(TrytesAlphabet[v2])
	}
	trytes, err := giota.ToTrytes(output)
	if err != nil {
		// TODO something with err
		fmt.Println(err)
	}
	return trytes
}

/*ChunkMessageToTrytesWithStopper will check if something is already trytes.  If it
is, we just return it.  If not, we convert it to trytes, add the stopper tryte to the
end of it, then return it*/
func ChunkMessageToTrytesWithStopper(messageString string) (giota.Trytes, error) {
	// messageString will be either a binary string or will already be in trytes
	trytes, err := giota.ToTrytes(messageString)
	if err == nil {
		// not capturing here since this isn't a "real" error
		return trytes, nil
	}

	// TODO look at this again.  Old implementation used RunesToTrytes

	//trytes, err = giota.ToTrytes(RunesToTrytes([]rune(messageString)) + StopperTryte)
	//trytesStart, err := ASCIIToTrytes(messageString)
	trytesStart := BytesToTrytes([]byte(messageString))
	if err != nil {
		// TODO something with err
		fmt.Println(err)
	}
	trytes, err = giota.ToTrytes(string(trytesStart) + StopperTryte)
	if err != nil {
		// TODO something with err
		fmt.Println(err)
	}
	return trytes, err
}

/*RunesToTrytes takes a rune array and converts to a tryte string*/
func RunesToTrytes(r []rune) string {
	var output string
	for _, c := range r {
		v1 := c % 27
		v2 := (c - v1) / 27
		output += string(TrytesAlphabet[v1]) + string(TrytesAlphabet[v2])
	}
	return output
}

/*MakeAddress accepts a string (should be a hex string) and converts it to an
iota address*/
func MakeAddress(hashString string) string {
	bytes, err := hex.DecodeString(hashString)
	if err != nil {
		// TODO something with err
		return ""
	}

	result := string(BytesToTrytes(bytes))

	if len(result) > 81 {
		return result[0:81]
	} else if len(result) < 81 {
		return padWith9s(result, 81)
	}
	return result
}

/*PadWith9s accepts a string to pad and desired length and pads that string with
9s up to the desired length*/
func padWith9s(stringToPad string, desiredLength int) string {
	padCountInt := desiredLength - len(stringToPad)
	var retStr = stringToPad + strings.Repeat("9", padCountInt)
	return retStr[0:desiredLength]
}
