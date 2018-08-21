package use

import "fmt"

//访问制定索引位置的值
func SimpleAccessSpecialIndex() interface{} {
	array := [5]int{10, 20, 30, 40, 50}
	return array[3]
}

//访问指针数组的元素,数组中存的地址,默认为nil
/*
array := use.AccessPtrArray()

	for _, v := range array {
		if v != nil {
			fmt.Println("v is：", *v)
		}
	}
	直接打印v是地址
*/
func AccessPtrArray() [5]*int {
	array := [5]*int{0: new(int), 1: new(int)}
	*array[0] = 10
	*array[1] = 20
	return array
}

//把同样类型的一个数组赋值给另外一个数组,只有类型和长度一样,才可以赋值
func ArrayCopy() {
	var array1 [5]string

	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}

	array1 = array2

	fmt.Println("array is:", array1)
}

//赋值数组指针
func ArrayPtrCopy() {
	var array1 [3]*string

	array2 := [3]*string{new(string), new(string), new(string)}

	*array2[0] = "Red"
	*array2[1] = "Blue"
	*array2[2] = "Green"

	array1 = array2

	//赋值之后.两个数组指向同一组字符串,每个元素的地址也是一样的
	for _, v := range array1 {
		fmt.Println("v is:", *v, ",addr is:", v)
	}

	for _, v1 := range array2 {
		fmt.Println("v1 is:", *v1, "addr is:", v1)
	}
}

//多维数组,初始化所有
func MoreDirectionArray() {
	array := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println("array is:", array)
}

//多维数组,声明并初始化外层数组索引为1和3的元素
func MoreDirectionInitOpen() {
	array := [3][2]int{0: {10, 11}, 2: {20, 21}}
	fmt.Println("array is:", array)
}

//多维数组,初始化内层和外层的单个元素
func MoreDirectionInit() {
	array := [3][2]int{1: {0: 10}}
	fmt.Println("array is:", array)
}
