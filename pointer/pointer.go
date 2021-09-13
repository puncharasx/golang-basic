package pointer

import "fmt"

func Pointer() {
	var x int
	x = 10
	var y *int
	y = &x
	fmt.Println(&x)
	fmt.Println(y)
}
