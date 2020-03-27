package main
//在goland里面不需要用那些命令，直接配置好之后右上角点运行就可以了
//至于使用bee run导致的报错懒得去研究了
//其实这个和django差不多，也是那种MVC的
//go语言的执行流程：先routers里面执行init再到main的run里面
import (
	_ "awesomeProject/beego_proj/routers"
	//这个下划线表示是引入routers包并且执行init方法
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

