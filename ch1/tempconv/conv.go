/*
@Time : 2019-03-25 13:56 
@Author : zhangjun
@File : conv
@Description:
@Run:
*/
package tempconv


func FToC(f Fahrenheit) Celsius {
	return Celsius((f-32)*5/9)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5+32)
}


func LbToKg(lb Lb) Kg {
	return Kg(lb*0.45)
}

func KgToLb(kg Kg) Lb {
	return Lb(kg/0.45)
}