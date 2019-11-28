package auth

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"

	//"fmt"
	//"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/xlotz/gocmdb/server/models"
	"github.com/xlotz/gocmdb/server/forms"
	"net/http"
)

type Session struct {

}

func (s *Session) Name() string{
	return "session"
}

func (s *Session) Is(c *context.Context) bool {

	return c.Input.Header("Authentication") == ""
	//return true

}

func (s *Session) IsLogin(c *LoginRequiredController) *models.User {

	if session := c.GetSession("user"); session != nil {
		if id, ok := session.(int); ok {

			return models.DefaultUserManager.GetById(id)

		}
	}

	return nil
}

func (s *Session) GoToLoginPage(c *LoginRequiredController) {
	//c.Redirect("/auth/login", http.StatusFound)
	c.Redirect(beego.URLFor(beego.AppConfig.String("login")), http.StatusFound)
}

func (s *Session) Login(c *AuthController) bool{

	form := &forms.LoginForm{}
	valid := &validation.Validation{}

	if c.Ctx.Input.IsPost() {

		if err := c.ParseForm(form); err != nil {
			valid.SetError("error", err.Error())
		}else {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
			}else if ok {
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

//func (s *Session) Index(c *AuthController){
//
//}


func (s *Session) Logout(c *AuthController) {

	c.DestroySession()
	c.Redirect(beego.URLFor(beego.AppConfig.String("login")), http.StatusFound)

}




//====================== Token 部分

type Token struct {

}


func (t *Token) Name() string{
	return "token"
}

func (t *Token) Is(c *context.Context) bool {

	return strings.ToLower(strings.TrimSpace(c.Input.Header("Authentication"))) == "token"
	//return true

}

func (t *Token) IsLogin(c *LoginRequiredController) *models.User {
	accessKey := strings.TrimSpace(c.Ctx.Input.Header("AccessKey"))
	secrectKey := strings.TrimSpace(c.Ctx.Input.Header("SecrectKey"))

	if token := models.DefaultTokenManager.GetByKey(accessKey, secrectKey); token != nil && token.User.DeletedTime == nil {

		return token.User

	}

	return nil
}

func (t *Token) GoToLoginPage(c *LoginRequiredController) {

	c.Data["json"] = map[string]interface{}{
		"code":403,
		"text": "Token 错误",
		"result": nil,
	}
	c.ServeJSON()

}

func (t *Token) Login(c *AuthController) bool{

	c.Data["json"] = map[string]interface{}{
		"code": 200,
		"text": "请使用Token请求API",
		"result": nil,
	}



	c.ServeJSON()
	return false
}



func (t *Token) Logout(c *AuthController) {

	//c.DestroySession()
	//c.Redirect("/auth/login", http.StatusFound)

	c.Data["json"] = map[string]interface{}{
		"code": 200,
		"text": "退出登录成功",
		"result": nil,
	}

	c.ServeJSON()

}


func init(){
	DefaultManger.Reginster(new(Session))
	DefaultManger.Reginster(new(Token))
}
