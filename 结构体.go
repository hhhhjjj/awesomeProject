package main

import "fmt"

//感觉这个就是class来实现面向对象，没有继承的概念，但是可以结构嵌入
//这里就定义了一个结构体
type User struct{
	name string
	age int
	address string
}

func main() {
	//user := User{name: "test", age:10}
	////就这样就可以使用了
	//user.address = "london"
	//fmt.Println(user)
	person:= struct {//匿名结构
		name string
		age int
	}{name:"匿名",age:1}
	//这样子就是个匿名结构体
	fmt.Println("person:",person)
}