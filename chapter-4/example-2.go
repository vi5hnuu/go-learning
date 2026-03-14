// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var arr1 [5]int
// 	fmt.Println("Array 1:", arr1)
// 	var arr2 = [10]int{0}
// 	fmt.Println("Array 2:", arr2)
// 	var arr3 = [...]int{1, 2, 3, 4, 5}
// 	fmt.Println("Array 3:", arr3)
// 	var arr4 = [10]int{1, 2, 3, 0, 5}
// 	fmt.Println("Array 4:", arr4)

// 	fmt.Println(arr1==arr2, arr1==arr3, arr1==arr4)
// }


///////////

package main

import (
	"fmt"
)

func main() {
	var arr1 [5]int
	fmt.Println("Array 1:", arr1)
	var arr2 = [5]int{0}
	fmt.Println("Array 2:", arr2)
	var arr3 = [...]int{1, 2, 3, 4, 5}
	fmt.Println("Array 3:", arr3)
	var arr4 = [5]int{1, 2, 3, 0, 5}
	fmt.Println("Array 4:", arr4)

	fmt.Println(arr1==arr2, arr1==arr3, arr1==arr4)
}