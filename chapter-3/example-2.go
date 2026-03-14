// package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	var list []int

// 	for i:=0;i<10000000;i++{
// 		list=append(list,i)
// 	}
// 	var m runtime.MemStats
// 	runtime.ReadMemStats(&m)
// 	fmt.Printf("Total Allocated (Heap) %v MiB\n",m.TotalAlloc/1024/1024)
// }

//////////////////
package main

import (
	"fmt"
	"runtime"
)

func main() {
	var list []uint8

	for i:=0;i<10000000;i++{
		list=append(list,uint8(i))
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Total Allocated (Heap) %v MiB\n",m.TotalAlloc/1024/1024)
}