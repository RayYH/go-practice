package main

import (
	"fmt"
	"strings"
)

func main() {
	// 第三个参数表示最多替换的数目，负数表示不受限制 (等价于 `ReplaceAll`)，如果 `old` 参数 (第二个参数) 为空，则代表替换每一个空格
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      // oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) // moo moo moo
	fmt.Println(strings.Replace("something", "", "A", -1))            // AsAoAmAeAtAhAiAnAgA
}
