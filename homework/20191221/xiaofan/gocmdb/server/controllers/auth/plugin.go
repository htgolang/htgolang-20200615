package auth

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/validation"
	"github.com/dcosapp/gocmdb/server/forms"
	"github.com/dcosapp/gocmdb/server/models"
	"net/http"
	"strings"
)

/*
	plugin来负责管理不同的登录插件
*/

// session插件
type Session struct {
}

// session的Name
func (s *Session) Name() string {
	return "session"
}

// 判断当前的登录方式是不是自己这个插件, 返回true即是
func (s *Session) Is(c *context.Context) bool {
	return c.Input.Header("Authentication") == ""
}

// 通过session获取信息, 判断用户是否登录
func (s *Session) IsLogin(c *LoginRequireController) *models.User {
	if session := c.GetSession("user"); session != nil {
		if uid, ok := session.(int); ok {
			return models.DefaultUserManager.GetById(uid)
		}
	}
	return nil
}

// 跳转到登录页面
func (s *Session) GoToLoginPage(c *LoginRequireController) {
	if c.Ctx.Input.IsAjax() {
		c.Data["json"] = map[string]interface{}{
			"code":   401,
			"text":   "请进行登录",
			"result": nil,
		}
		c.ServeJSON()
	}
	c.Redirect(beego.URLFor(beego.AppConfig.String("login")), http.StatusFound)
}

// 登录
func (s *Session) Login(c *AuthController) bool {
	if session := c.GetSession("user"); session != nil {
		if _, ok := session.(int); ok {
			c.Redirect(beego.URLFor(beego.AppConfig.String("home")), http.StatusFound)
			return true
		}
	}

	form := &forms.LoginForm{}
	valid := &validation.Validation{}

	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err != nil {
			_ = valid.SetError("error", err.Error())
		} else {
			if ok, err := valid.Valid(form); err != nil {
				_ = valid.SetError("error", err.Error())
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

// 登出
func (s *Session) Logout(c *AuthController) {
	c.DestroySession()
	c.Redirect(beego.URLFor(beego.AppConfig.String("login")), http.StatusFound)
}

// token插件
type Token struct {
}

// 插件Name
func (s *Token) Name() string {
	return "token"
}

// 如果Header里面有Authentication，则认为是通过token登录
func (s *Token) Is(c *context.Context) bool {
	return strings.ToLower(strings.TrimSpace(c.Input.Header("Authentication"))) == "token"
}

// 判断token是否正确
func (s *Token) IsLogin(c *LoginRequireController) *models.User {
	c.EnableXSRF = false
	accessKey := strings.TrimSpace(c.Ctx.Input.Header("AccessKey"))
	secretKey := strings.TrimSpace(c.Ctx.Input.Header("SecretKey"))
	if token := models.DefaultTokenManager.GetByKey(accessKey, secretKey); token != nil && token.User.DeletedTime == nil {
		return token.User
	}
	return nil
}

// token认证失败不需要跳转页面，返回错误状态码即可
func (s *Token) GoToLoginPage(c *LoginRequireController) {
	c.Data["json"] = map[string]interface{}{
		"code":   403,
		"text":   "请使用正确Token发起请求",
		"result": nil,
	}
	c.ServeJSON()
}

// 登录
func (s *Token) Login(c *AuthController) bool {
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "请使用Token请求API",
		"result": nil,
	}
	c.ServeJSON()
	return false
}

// 登出
func (s *Token) Logout(c *AuthController) {
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "退出登录成功",
		"result": nil,
	}
	c.ServeJSON()
}

// 注册插件
func init() {
	DefaultManager.Register(new(Session))
	DefaultManager.Register(new(Token))
}
