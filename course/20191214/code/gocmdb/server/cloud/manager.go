package cloud

type Manager struct {
	Plugins map[string]ICloud
}

func NewManager() *Manager {
	return &Manager{
		Plugins: make(map[string]ICloud),
	}
}

func (m *Manager) Register(c ICloud) {
	m.Plugins[c.Type()] = c
}

func (m *Manager)Cloud(typ string) (ICloud, bool) {
	cloud, ok := m.Plugins[typ]
	return cloud, ok
}


var DefaultManager = NewManager()