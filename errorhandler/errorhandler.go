package errorhandler

import (
	"fmt"
)

func FormatErrors(errors []error) string {
	formattedErrors := ""
	for idx, err := range errors {
		formattedErrors += fmt.Sprintf("%d. %v\n", idx, err)
	}
	return formattedErrors
}

func Return(errors []error) {
	fmt.Printf("Program execution failed\nErrors:\n-------\n%v", FormatErrors(errors))
}
