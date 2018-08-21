package create

//使用类型0值进行初始化
func CreateWithZero() interface{} {
	var array [5]int
	return array
}

//使用数组字面量声明,用具体值进行初始化
func CreateWithValue() interface{} {
	array := [5]int{10, 20, 30, 40, 50}
	return array
}

//自动计算数组的长度,用...代表数组长度,数组长度由初始化值的数量决定
func CreateWithAutoLength() (interface{}, int) {
	array := [...]int{1, 2, 3, 4, 5, 6, 7}
	return array, len(array)
}

//声明数组并制定特定元素的值
func CreateWithSpecialValue() interface{} {
	array := [5]int{1: 10, 2: 20}
	return array
}
