/*
@Time : 2019-03-25 13:47 
@Author : zhangjun
@File : tempconv
@Description:
@Run:
*/
package main

import (
	"fmt"
	"github.com/exerciseGo/ch1/tempconv"
	"os"
	"strconv"
)

func main(){
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "tempconv : %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		lb := tempconv.Lb(t)
		kg := tempconv.Kg(t)

		fmt.Printf("%s = %s, %s = %s, %s = %s, %s = %s", f, tempconv.FToC(f), c, tempconv.CToF(c), lb, tempconv.LbToKg(lb), kg, tempconv.KgToLb(kg))

	}
}
