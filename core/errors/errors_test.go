package errors

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyError(t *testing.T) {
	run := func() error {
		return &MyError{
			"now",
			"something happened",
		}
	}
	err := run()
	assert.NotNil(t, err)
	assert.Equal(t, "at now, something happened", fmt.Sprint(err))
}
