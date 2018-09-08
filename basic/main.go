package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"math/cmplx"
	"os"
	"reflect"
	"runtime"
	"strconv"
)

//声明变量并且初始化
func varInitValue() {
	var a, b int = 3, 4
	var c, d = 5, 6
	var s1 string = "s1 with type"
	var s2 = "s2 no type"
	fmt.Println(a, b, c, d)
	fmt.Println(s1, s2)
}

//不声明类型,自行推断
func varTypeDeduction() {
	var a, b, c, s = 1, 2, true, "hehhe"
	fmt.Println(a, b, c, s)
}

//不用var声明
func varShorter() {
	a, b, c, s := 1, 2, true, "哈哈"
	fmt.Println(a, b, c, s)

}

//包内的var 不能用 :=声明变量
var ga = 1

//一下子多个
var (
	gb = 2
	gc = true
	gs = "哈哈哈"
)

//欧拉公式 e的iπ + 1 = 0
func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

//计算直角三角形面积
func calcTriangle(a, b int) (c int) {
	return int(math.Sqrt(float64(a*a + b*b)))
}

//定义常量
func consts() {
	const a, b = 3, 4
	fmt.Println(a, b)
}

//定义枚举
func enums() {
	const (
		cpp    = 0
		java   = 1
		py     = 2
		golang = 3
	)
	fmt.Println(cpp, java, py, golang)
}

//iota定义常量 iota开始是0 每次自增1
func enumsWithIota() {
	const (
		cpp = iota
		java
		py
		golang
	)
	fmt.Println(cpp, java, py, golang)
}

//iota定义b,kb,mb,gb,tb,pb..
func storeSizeWithIota() {
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

/*
  读文件, string(字节切片) 转字符串
  bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
*/
func readFile(filename string) {
	if contents, err := ioutil.ReadFile(filename); err == nil {
		fmt.Println(string(contents))
	} else {
		fmt.Println("can not open file:", err)
	}
}

//switch例子,switch会自动break
func switchFunc(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)
	}
	return result
}

//switch后面可以没有表达式
func switchWithoutOp(score int) string {
	switch {
	case score < 60:
		return "F"
	case score < 80:
		return "C"
	case score < 90:
		return "B"
	default:
		return "A"
	}
}

//数字转二进制 1 for 省略初始条件
func convert2Binary(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

//按行读文件
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//无限的for
func forWithoutEnd() {
	for {
		fmt.Println("hhah")
	}
}

//可以不指定返回值名称,也可以定义(建议)
func div(a, b int) (int, int) {
	return a / b, a % b
}

/*
   函数作为参数 fmt.Println(apply(pow, 3, 4))
   result := apply(
		func(a int, b int) int {
			return a + b
		}, 3, 4)
	fmt.Println("result is:", result)
*/
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("call func %s with args (%d,%d)", opName, a, b)
	fmt.Println()
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func main() {
	fmt.Println(ga, gb, gc, gs)
	result := apply(
		func(a int, b int) int {
			return a + b
		}, 3, 4)
	fmt.Println("result is:", result)
}
