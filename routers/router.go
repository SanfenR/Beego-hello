package routers

import (
	"Beego-hello/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
}
