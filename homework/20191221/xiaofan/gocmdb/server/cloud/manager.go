package cloud

type Manager struct {
	Plugins map[string]ICloud
}

// 初始化
func NewManager() *Manager {
	return &Manager{Plugins: make(map[string]ICloud)}
}

// 注册云插件
func (m *Manager) Register(c ICloud) {
	m.Plugins[c.Type()] = c
}

// 获取对应插件
func (m *Manager) Cloud(typ string) (ICloud, bool) {
	cloud, ok := m.Plugins[typ]
	return cloud, ok
}

var DefaultManager = NewManager()
