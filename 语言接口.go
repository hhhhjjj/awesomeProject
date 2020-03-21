package main
import (
	"fmt"
)

//这个是定义了接口
type Phone interface {
	call()
}

//这个是结构体
type NokiaPhone struct {
}

//接口实现方法
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

//结构体
type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone
	//我怎么感觉这个像是父类子类这些
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}