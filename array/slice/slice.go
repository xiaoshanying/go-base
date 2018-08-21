package slice

import "fmt"

//创建一个切片,长度和容量都是5
func CreateSliceDefault() {
	slice := make([]string, 5)
	//slice is: [    ]
	fmt.Println("slice is:", slice)
}

//创建一个长度为5,容量为6的切片
func CreateSliceWithParam() {
	slice := make([]int, 5, 6)

	fmt.Println("slice is:", slice)
}

func CreateNilSlice() {
	var slice []int
	fmt.Println("slice is:", slice)
}

//创建一个空切片或者 slice := []int{}
func CreateNullSlice() {
	slice := make([]int, 0)
	fmt.Println("slice is:", slice)
}
