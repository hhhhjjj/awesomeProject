package main
//在goland里面不需要用那些命令，直接配置好之后右上角点运行就可以了
//至于使用bee run导致的报错懒得去研究了
import (
	_ "awesomeProject/beego_proj/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

