package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)

	for {
		go func() {
			a := 1

			runtime.SetFinalizer(&a, func(obj *int) {
				fmt.Printf("finalize a, %v\n", obj)
			})
			runtime.GC()
			os.Exit(1)
		}()
	}
}
