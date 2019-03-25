

/*
@Time : 2019-03-25 16:31 
@Author : zhangjun
@File : test
@Description:
@Run:
*/
package main

import "fmt"

/*
 * Go 语言中的字符串是一个只读的字节数组切片
		type StringHeader struct {
			Data uintptr
			Len  int
		}
 * 字符串作为只读的类型，我们并不会直接向字符串直接追加元素改变其本身的内存空间，所有追加的操作都是通过拷贝来完成的。
 * 新的字符串其实就是一片新的内存空间，与原来的字符串没有任何关联。指针地址并不会改
 * gc的暂停时间小于1ms，指阻塞程序小于1ms
 * 字符串类型设置为不可变是由一些好处的。 我想到的有，可以更好的重用字符串。 还有多线程操作字符串不用加锁。--南哥
 */

func main(){
	s := "hello,world"
	t := s
	fmt.Printf("s===%s  t===%s\n", s, t)
	fmt.Printf("s===%p  t===%p\n", &s, &t)

	s += ";zhangjun"
	fmt.Printf("\n\n s===%s  t===%s\n", s, t)
	fmt.Printf("s===%p  t===%p\n", &s, &t)
}


