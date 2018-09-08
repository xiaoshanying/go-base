package main

import (
	"fmt"
)

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	// [2 3 4 5]
	fmt.Println("arr[2:6] = ", arr[2:6])
	//[0 1 2 3 4 5]
	fmt.Println("arr[:6] = ", arr[:6])
	//[2 3 4 5 6 7]
	fmt.Println("arr[2:] = ", arr[2:])
	//输出数组的全部元素
	fmt.Println("arr[:] = ", arr[:])

	//改变切片,会改变底层的数组,切片可以向后扩展,不可以向前
	s1 := arr[2:6]
	fmt.Println("s1 is ", s1)
	//slice是可以扩展的,
	s2 := s1[3:5]
	fmt.Println("s2 is ", s2)

	s3 := append(s2, 10)
	s4 := append(s3, 20)
	fmt.Println("s3 is ", s3)
	fmt.Println("s4 is ", s4)
}
