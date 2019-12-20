package auth

import (
	"net/http"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/validation"

	"github.com/imsilence/gocmdb/server/forms"
	"github.com/imsilence/gocmdb/server/models"
)

type Session struct {
}

func (s *Session) Name() string {
	return "session"
}

func (s *Session) Is(c *context.Context) bool {
	return c.Input.Header("Authentication") == ""
}

func (s *Session) IsLogin(c *LoginRequiredController) *models.User {
	if session := c.GetSession("user"); session != nil {
		if uid, ok := session.(int); ok {
			return models.DefaultUserManager.GetById(uid)
		}
	}
	return nil
}

func (s *Session) GoToLoginPage(c *LoginRequiredController) {
	if c.Ctx.Input.IsAjax() {
		// ajax请求
		c.Data["json"] = map[string]interface{}{
			"code":   401,
			"text":   "请进行登录",
			"result": nil,
		}
		c.ServeJSON()
	} else {
		// http请求
		c.Redirect(beego.URLFor(beego.AppConfig.String("login")), http.StatusFound)
	}
}

func (s *Session) Login(c *AuthController) bool {
	if session := c.GetSession("user"); session != nil {
		if _, ok := session.(int); ok {
			c.Redirect(beego.URLFor(beego.AppConfig.String("home")), http.StatusFound)
			return false
		}
	}

	form := &forms.LoginForm{}
	valid := &validation.Validation{}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err != nil {
			valid.SetError("error", err.Error())
		} else {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
			} else if ok {
				c.SetSession("user", form.User.Id)
				c.Redirect(beego.URLFor(beego.AppConfig.String("home")), http.StatusFound)
				return true
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
	c.Redirect(beego.URLFor(beego.AppConfig.String("login")), http.StatusFound)
}

type Token struct {
}

func (t *Token) Name() string {
	return "token"
}

func (t *Token) Is(c *context.Context) bool {
	return strings.ToLower(strings.TrimSpace(c.Input.Header("Authentication"))) == "token"
}

func (t *Token) IsLogin(c *LoginRequiredController) *models.User {
	// 关闭XSRF认证
	c.EnableXSRF = false

	accessKey := strings.TrimSpace(c.Ctx.Input.Header("AccessKey"))
	secrectKey := strings.TrimSpace(c.Ctx.Input.Header("SecrectKey"))
	if token := models.DefaultTokenManager.GetByKey(accessKey, secrectKey); token != nil && token.User.DeletedTime == nil {
		return token.User
	}
	return nil
}

func (t *Token) GoToLoginPage(c *LoginRequiredController) {
	c.Data["json"] = map[string]interface{}{
		"code":   403,
		"text":   "请使用正确Token发起请求",
		"result": nil,
	}
	c.ServeJSON()
}

func (t *Token) Login(c *AuthController) bool {
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "请使用Token请求API",
		"result": nil,
	}
	c.ServeJSON()
	return false
}

func (t *Token) Logout(c *AuthController) {
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "退出登陆成功",
		"result": nil,
	}
	c.ServeJSON()
}

func init() {
	DefaultManger.Register(new(Session))
	DefaultManger.Register(new(Token))
}
