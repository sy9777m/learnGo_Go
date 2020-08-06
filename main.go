package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(word ...string) {
	fmt.Println(word)
}

func main() {
	fmt.Println(multiply(2, 3))
	totalLength, upperName := lenAndUpper("Siyun Min")
	fmt.Println(totalLength, upperName)
	totalLen, _ := lenAndUpper("Siyun Min")
	fmt.Println(totalLen)
	repeatMe("S", "I", "Y", "U", "M")
}
