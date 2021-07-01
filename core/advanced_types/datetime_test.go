package advanced_types

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 当前时间可以使用 time.Now() 获取，或者使用 t.Day()、t.Minute() 等来获取时间的一部分

func TestTimeFormat(t *testing.T) {
	// Go 中时间格式化不是采用 Java 中的 yyyy-MM-dd HH:mm:ss，而是采用一种叫做 layout 的布局格式
	dt := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	// Format 可以根据一个格式化字符串来将一个时间 t 转换为相应格式的字符串
	// 你可以使用一些预定义的格式，如：time.ANSIC 或 time.RFC822
	assert.Equal(t, fmt.Sprintf(dt.Format("2006-01-02 15:04:05")), "2009-11-10 23:00:00")
	assert.Equal(t, fmt.Sprintf(dt.Format(time.ANSIC)), "Tue Nov 10 23:00:00 2009")
	assert.Equal(t, fmt.Sprintf(dt.Format(time.RFC822)), "10 Nov 09 23:00 UTC")

	layout := "2006-01-02 15:04:05"
	assert.NotNil(t, time.Now().Format(layout))
}

func TestTimestamp(t *testing.T) {
	// Go 中时间戳的类型是 int64
	dt := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	assert.Equal(t, dt.Unix(), int64(1257894000))
	assert.Equal(t, dt.UnixNano(), int64(1257894000000000000))
}
