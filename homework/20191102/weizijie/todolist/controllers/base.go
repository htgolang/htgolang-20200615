package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
	"todolist/models"
)

type BaseController struct {
	beego.Controller
	User *models.User
}

func (this *BaseController) ValidateSession() {
	if session := this.GetSession("user"); session != nil {
		if id, ok := session.(int); ok {
			user := models.User{Id: id}
			if err := models.GetUser(&user); err == nil {
				this.User = &user
				this.Data["user"] = &user
			} else {
				fmt.Printf("Can not get user for id %s, err: %s\n", id, err.Error())
			}
		} else {
			fmt.Printf("Invaid session is %v\n", session)
		}
	} else {
		fmt.Println("Not Login")
	}
}

type LoginRequiredController struct {
	BaseController
}

func (this *LoginRequiredController) Prepare() {
	this.ValidateSession()
	if this.User == nil {
		this.Redirect("/auth/login", http.StatusFound)
	} else {
		beego.ReadFromRequest(&this.Controller)
	}
}
