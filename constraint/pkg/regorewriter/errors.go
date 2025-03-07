package regorewriter

import (
	"fmt"
	"io"
	"strings"
)

// Errors is a list of error.
type Errors []error

// Error implements error.
func (errs Errors) Error() string {
	var s []string
	for _, err := range errs {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

// Format implements fmt.Formatter to make this play nice with handling stack traces produced from
// github.com/pkg/errors.
//nolint:errcheck // suppress errors rendering errors
func (errs Errors) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		_, _ = fmt.Fprintf(s, "errors (%d):\n", len(errs))
		for _, err := range errs {
			if formatter, ok := err.(fmt.Formatter); ok {
				formatter.Format(s, verb)
				_, _ = io.WriteString(s, "\n")
			} else {
				_, _ = fmt.Fprintf(s, "%v\n", err)
			}
		}

	case 's':
		_, _ = io.WriteString(s, errs.Error())

	case 'q':
		_, _ = fmt.Fprintf(s, "%q", errs.Error())
	}
}
