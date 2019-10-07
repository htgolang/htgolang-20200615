package frontend

import (
	"github.com/astaxie/beego"
	"github.com/imsilence/issues/controllers/frontend"
)

func init() {
	beego.Router("/", &frontend.IndexController{})
}
