package utils

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
