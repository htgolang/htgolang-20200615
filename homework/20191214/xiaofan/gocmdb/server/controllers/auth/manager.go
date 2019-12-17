package auth

import (
	"github.com/astaxie/beego/context"
	"github.com/dcosapp/gocmdb/server/models"
)

/*
	manager.go主要负责使用插件来进行controller的一系列流程
*/

// 认证插件
type AuthPlugin interface {
	Name() string                                 // 插件名
	Is(*context.Context) bool                     // 返回true就代表认证方式是自己
	IsLogin(*LoginRequireController) *models.User // 判断是否登录，返回该用户或nil
	GoToLoginPage(*LoginRequireController)        // 跳转到登录页面
	Login(*AuthController) bool                   // 登录
	Logout(*AuthController)                       // 登出
}

// 负责管理插件的,其实叫PluginManager更合适
type Manager struct {
	plugins map[string]AuthPlugin
}

// 创建新的Manager
func NewManager() *Manager {
	return &Manager{
		plugins: map[string]AuthPlugin{},
	}
}

// 注册插件
func (m *Manager) Register(p AuthPlugin) error {
	m.plugins[p.Name()] = p
	return nil
}

// 获取插件类型
func (m *Manager) GetPlugin(c *context.Context) AuthPlugin {
	for _, plugin := range m.plugins {
		if plugin.Is(c) {
			return plugin
		}
	}
	return nil
}

// 判断是否已经登录
func (m *Manager) IsLogin(c *LoginRequireController) *models.User {
	if plugin := m.GetPlugin(c.Ctx); plugin != nil {
		return plugin.IsLogin(c)
	}
	return nil
}

// 跳转登录页面
func (m *Manager) GoToLoginPage(c *LoginRequireController) {
	if plugin := m.GetPlugin(c.Ctx); plugin != nil {
		plugin.GoToLoginPage(c)
	}
}

// 登录
func (m *Manager) Login(c *AuthController) bool {
	if plugin := m.GetPlugin(c.Ctx); plugin != nil {
		return plugin.Login(c)
	}
	return false
}

// 登出
func (m *Manager) Logout(c *AuthController) {
	if plugin := m.GetPlugin(c.Ctx); plugin != nil {
		plugin.Logout(c)
	}
}

var DefaultManager = NewManager()
