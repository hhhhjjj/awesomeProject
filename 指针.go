package main

import "fmt"

//这个不涉及指针的运算，不能像c语言一样直接操作指针，也不会导致内存溢出
//其实也就是&取地址，*根据地址取值
func main(){
	n := 18
	fmt.Println(&n)
	p := &n
	fmt.Println(*p)
}

