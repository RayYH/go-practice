package errors

import (
	"fmt"
)

// MyError custom error struct
type MyError struct {
	when string
	what string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %s, %s", e.when, e.what)
}
