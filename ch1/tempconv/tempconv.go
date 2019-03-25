/*
@Time : 2019-03-25 13:50 
@Author : zhangjun
@File : tempconv
@Description:华氏度与摄氏度转换，千克和磅的转换
@Run:go run cf.go 23
*/
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

type Kg float32
type Lb float32

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius	  = 0
	BoilingC Celsius 	  = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g.C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g.F", f)
}

func (lb Lb) String() string {
	return fmt.Sprintf("%g lb", lb)
}

func (kg Kg) String() string {
	return fmt.Sprintf("%g kg", kg)
}