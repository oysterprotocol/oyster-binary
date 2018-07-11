package oysterErrors

import (
	"testing"

	"github.com/pkg/errors"
)

var errArray = []error{
	errors.New("first error"),
	errors.New("second error"),
	errors.New("third error"),
}

func Test_ReturnFirstError(t *testing.T) {
	err := ReturnFirstError(errArray)
	if err.Error() != "first error" {
		t.Fatalf("ReturnFirstError() result should be %s but returned %s",
			"first error", err.Error())
	}
}

func Test_CollectErrors(t *testing.T) {

	var errString = "Error 1: first error\n" +
		"Error 2: second error\n" +
		"Error 3: third error\n"

	err := CollectErrors(errArray)
	if err.Error() != errString {
		t.Fatalf("CollectErrors() result should be %s but returned %s",
			errString, err.Error())
	}
}
