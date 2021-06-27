package advanced_types

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeFormat(t *testing.T) {
	// Go 中时间格式化不是采用 Java 中的 yyyy-MM-dd HH:mm:ss，而是采用一种叫做 layout 的布局格式
	dt := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	assert.Equal(t, fmt.Sprintf(dt.Format("2006-01-02 15:04:05")), "2009-11-10 23:00:00")

	layout := "2006-01-02 15:04:05"
	assert.NotNil(t, time.Now().Format(layout))
}

func TestTimestamp(t *testing.T) {
	// Go 中时间戳的类型是 int64
	dt := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	assert.Equal(t, dt.Unix(), int64(1257894000))
	assert.Equal(t, dt.UnixNano(), int64(1257894000000000000))
}
