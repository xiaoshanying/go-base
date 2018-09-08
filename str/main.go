package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "My name is wmz"
	splitNormal(name)
	fmt.Println()
}

func splitNormal(str string) {
	fmt.Print("|")
	for _, str1 := range strings.Split(str, " ") {
		fmt.Printf("%s|", str1)
	}
}

func splitAfter(str string) {

}
