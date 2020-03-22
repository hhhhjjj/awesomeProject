package routers

import (
	"awesomeProject/beego_proj/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/hello", &controllers.HelloControllers{})
    //这个要对齐的是里面的func，不是外面这个go文件
}
