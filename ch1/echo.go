/*
@Time : 2019-03-24 08:16 
@Author : zhangjun
@File : echo.go
*/
package main

import (
"fmt"
"os"
	"strings"
)
//go run echo.go test_dup.txt

func main()  {
	var s, sep string
	for i :=1; i<len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "  "
	}
	fmt.Println(s, "\n")
	EchoTwo()
	EchoThree()
	EchoOsCommandName()
}

func EchoTwo() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println("EchoTwo ---> " + s, "\n")
}

func EchoThree() {
	fmt.Println("EchoThree ---> " +strings.Join(os.Args[1:], ""), "\n")
}

func EchoOsCommandName(){
	fmt.Println("EchoOsCommandName:"+os.Args[0], "\n")
}
