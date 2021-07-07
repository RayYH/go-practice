package advanced_types

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeFormat(t *testing.T) {
	dt := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	assert.Equal(t, fmt.Sprintf(dt.Format("2006-01-02 15:04:05")), "2009-11-10 23:00:00")
	assert.Equal(t, fmt.Sprintf(dt.Format(time.ANSIC)), "Tue Nov 10 23:00:00 2009")
	assert.Equal(t, fmt.Sprintf(dt.Format(time.RFC822)), "10 Nov 09 23:00 UTC")

	layout := "2006-01-02 15:04:05"
	assert.NotNil(t, time.Now().Format(layout))
}

func TestTimestamp(t *testing.T) {
	dt := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	assert.Equal(t, dt.Unix(), int64(1257894000))
	assert.Equal(t, dt.UnixNano(), int64(1257894000000000000))
}
