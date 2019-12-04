package auth

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/validation"
	"github.com/xxdu521/cmdbgo/server/forms"
	"github.com/xxdu521/cmdbgo/server/models"
	"net/http"
	"strings"
)


type Session struct {}

func (s *Session) Name() string{
	return "session"
}

func (s *Session) Is(c *context.Context) bool {
	//根据每种登录方法的特征，去判断当前的登录方法是什么？
	return c.Input.Header("Authentication") == ""
}

func (s *Session) IsLogin(c *LoginRequiredController) *models.User {
	if session := c.GetSession("user"); session != nil {
		if uid,ok := session.(int); ok {
			return models.DefaultUserManager.GetById(uid)
			}
		}
	return nil
}

func (s *Session) GoToLoginPage(c *LoginRequiredController) {
	if c.Ctx.Input.IsAjax() {
		c.Data["json"] = map[string]interface{}{
			"code": 401,
			"text": "请登录",
			"result": nil,
		}
		c.ServeJSON()
		//c.Redirect(beego.URLFor(beego.AppConfig.String("login")),http.StatusFound)
	} else {
		//beego.ReadFromRequest(&c.Controller)
		c.Redirect(beego.URLFor(beego.AppConfig.String("login")), http.StatusFound)
	}
}

func (s *Session) Login(c *AuthController) bool {
	form := &forms.LoginForm{}
	valid := &validation.Validation{}
	if c.Ctx.Input.IsPost() {
		if c.Ctx.Input.IsPost() {
			if err := c.ParseForm(form); err != nil {
				valid.SetError("error",err.Error())
			} else {
				if ok,err := valid.Valid(form); err != nil {
					valid.SetError("error", err.Error())
				} else if ok {
					c.SetSession("user",form.User.Id)
					c.Redirect(beego.URLFor(beego.AppConfig.String("home")),http.StatusFound)
					return true
				}
			}
		}
	}
	c.TplName = "auth/login.html"
	c.Data["form"] = form
	c.Data["valid"] = valid
	return false
}

func (s *Session) Logout(c *AuthController) {
	c.DestroySession()
	c.Redirect(beego.URLFor(beego.AppConfig.String("login")),http.StatusFound)
}


type Token struct {}

func (s *Token) Name() string{
	return "token"
}

func (s *Token) Is(c *context.Context) bool {
	//根据每种登录方法的特征，去判断当前的登录方法是什么？
	return strings.ToLower(strings.TrimSpace(c.Input.Header("Authentication"))) == "token"
}

func (s *Token) IsLogin(c *LoginRequiredController) *models.User {
	accessKey := strings.TrimSpace(c.Ctx.Input.Header("AccessKey"))
	secrectKey := strings.TrimSpace(c.Ctx.Input.Header("SecrectKey"))
	if token := models.DefaultTokenManager.GetByKey(accessKey,secrectKey);token != nil && token.User.DeletedTime == nil {
		return token.User
	}
	return nil
}

func (s *Token) GoToLoginPage(c *LoginRequiredController) {
	c.Data["json"] = map[string]interface{}{
		"code":403,
		"text":"请使用正确的Token进行认证",
		"result":nil,
	}
	c.ServeJSON()
}

func (s *Token) Login(c *AuthController) bool {
	c.Data["json"] = map[string]interface{}{
		"code":200,
		"text":"请使用正确的Token进行认证",
		"result":nil,
	}
	c.ServeJSON()
	return false
}

func (s *Token) Logout(c *AuthController) {
	c.Data["json"] = map[string]interface{}{
		"code":200,
		"text":"退出成功",
		"result":nil,
	}
	c.ServeJSON()
}


func init() {
	DefaultManager.Register(new(Session))
	DefaultManager.Register(new(Token))
}