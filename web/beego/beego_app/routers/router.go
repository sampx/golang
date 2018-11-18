package routers

import (
	"github.com/sampx/golang/web/beego/beego_app/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
