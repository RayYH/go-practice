package basic

import (
    "fmt"
    "math"
)

// 一个安全的从 int 转为 uint 的 func
func Uint8FromInt(n int) (uint8, error) {
    if 0 <= n && n <= math.MaxUint8 {
        return uint8(n), nil
    }

    return 0, fmt.Errorf("%d is out of the uint8 range", n)
}
