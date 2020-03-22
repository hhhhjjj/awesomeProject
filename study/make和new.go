package main

import "fmt"

//这两就是创建分配类型内存
func main(){
	//要想声明变量直接var i int就行了，这时候默认值为0
	var i *int
	//这个是引用类型
	i = new(int)
	//对于引用类别的变量，不光要声明，还要分配内容空间
	//new返回的是指针，指向内存地址
	*i = 10
	fmt.Println(*i)
}
//make也差不多，但是只用于slice，map以及channel的初始化
//其实new都没人用，直接：=就行了