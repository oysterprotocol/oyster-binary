package oysterErrors

import (
	"bytes"
	"strconv"

	"github.com/pkg/errors"
)

/*ReturnFirstError accepts an array of errors and returns the first that is not nil*/
func ReturnFirstError(arrayOfErrs []error) error {
	var err error
	for _, errInArray := range arrayOfErrs {
		if errInArray != nil {
			err = errInArray
			break
		}
	}
	return err
}

/*CollectErrors returns all the errors*/
func CollectErrors(arrayOfErrs []error) error {
	var buffer bytes.Buffer
	var errString string
	for i, errInArray := range arrayOfErrs {
		if errInArray != nil {
			buffer.WriteString("Error ")
			buffer.WriteString(strconv.Itoa(i + 1))
			buffer.WriteString(": ")
			buffer.WriteString(errInArray.Error())
			buffer.WriteString("\n")
		}
	}
	errString = buffer.String()
	if errString == "" {
		return nil
	}
	return errors.New(errString)
}
