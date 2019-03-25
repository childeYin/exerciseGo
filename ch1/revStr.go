/*
@Time : 2019-03-25 20:07 
@Author : zhangjun
@File : revStr
@Description:翻转数组内容
@Run: go run revStr.go
*/
package main

import "fmt"

func reverse(s []int) {
	for i, j :=0, len(s)-1; i<j; i,j= i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main(){
	a := [...]int{0,1,2,3,4,5,6,7}
	reverse(a[:])
	fmt.Println(a)
}

