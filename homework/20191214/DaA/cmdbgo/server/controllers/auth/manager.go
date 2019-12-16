package auth

import (
	"github.com/astaxie/beego/context"
	"github.com/xxdu521/cmdbgo/server/models"
)

type AuthPlugin interface {
	Name() string
	Is(*context.Context) bool
	IsLogin(*LoginRequiredController) *models.User
	GoToLoginPage(*LoginRequiredController)
	Login(*AuthController) bool
	Logout(*AuthController)
}

type Manager struct {
	plugins map[string]AuthPlugin
}

func NewManager() *Manager{
	return &Manager{
		plugins: map[string]AuthPlugin{},
	}
}
//插件注册方法，引入插件的实例，调用插件实例的Name方法，把插件名字，写入到plugins的映射里。
func (m *Manager) Register(p AuthPlugin) {
	m.plugins[p.Name()] = p
}

//判断用户访问的方式，通过IS方法(is方法是核心，其实就是通过用户访问的header去确认用户是session/token访问)，通过Ctx(beego的context.Context指针实例)里的header特征，去判断用户的访问方式。

//如果我们不用插件的方式处理，也可以按照IS的思想，去自行判断处理。
func (m *Manager) GetPlugin(c *context.Context) AuthPlugin {
	for _,plugin := range m.plugins {
		if plugin.Is(c) {
			return plugin
		}
	}
	return nil
}
//从loginrequired控制器产生的调用，是否登录方法，通过调用插件的Islogin方法，返回一个User实例，来确定是否登录
func (m *Manager) IsLogin(c *LoginRequiredController) *models.User{
	if plugin := m.GetPlugin(c.Ctx); plugin != nil {
		return plugin.IsLogin(c)
	}
	return nil
}
//从loginrequired控制器产生的调用，跳转到登录页方法，调用插件的GoToLoginPage方法，
func (m *Manager) GoToLoginPage(c *LoginRequiredController) {
	if plugin := m.GetPlugin(c.Ctx); plugin != nil {
		plugin.GoToLoginPage(c)
	}
}
//从auth控制器产生的调用，登录方法，
func (m *Manager) Login(c *AuthController) bool {
	if plugin := m.GetPlugin(c.Ctx); plugin != nil {
		return plugin.Login(c)
	}
	return false
}
//从auth控制器产生的调用，退出方法，
func (m *Manager) Logout(c *AuthController) {
	if plugin := m.GetPlugin(c.Ctx); plugin != nil {
		plugin.Logout(c)
	}
}

var DefaultManager = NewManager()