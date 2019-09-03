package session

import (
	uuid "github.com/satori/go.uuid"
	"net/http"
	"sync"
	"time"
)

type Session struct {
	contents    map[string]interface{}
	expiredTime time.Time
	lock        sync.RWMutex
}

func NewSession(expiredTime time.Time) *Session {
	return &Session{
		contents:    make(map[string]interface{}),
		expiredTime: expiredTime,
	}
}

func (s *Session) Set(key string, value interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.contents[key] = value
}

func (s *Session) Get(key string) (interface{}, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	value, ok := s.contents[key]
	return value, ok
}

func (s *Session) Delete(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.contents, key)
}

type Manager struct {
	sessions   map[string]*Session
	cookieName string //sid
	timeout    int
	lock       sync.Mutex
}

func NewManager(cookieName string, timeout int) *Manager {
	manager := &Manager{
		sessions:   make(map[string]*Session),
		cookieName: cookieName,
		timeout:    timeout,
	}
	go manager.SessionGC()
	return manager
}

func (m *Manager) SessionStart(w http.ResponseWriter, r *http.Request) *Session {
	var session *Session
	var sid string

	// 查看当前连接是否存在sid这个cookie
	cookie, err := r.Cookie(m.cookieName) // m.cookieName == "sid"

	m.lock.Lock()
	defer m.lock.Unlock()

	// 如果没有cookie
	if err != nil || cookie.Value == "" {
		sid = uuid.NewV4().String()

		// 生成一个已经有过期时间的session
		session = NewSession(time.Now().Add(time.Duration(m.timeout) * time.Second))

		// 给Manager添加一个session
		m.sessions[sid] = session
	} else if _, ok := m.sessions[cookie.Value]; !ok { // 如果当前连接的sid的值不存在
		sid := uuid.NewV4()
		session = &Session{}
		m.sessions[sid.String()] = session
	} else {
		// 如果cookie存在并且没有问题，则直接将其session返回
		return m.sessions[cookie.Value]
	}

	// 为当前连接设置cookie
	cookie = &http.Cookie{
		Name:     m.cookieName,
		Value:    sid,
		Path:     "/",
		MaxAge:   m.timeout,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	return session
}

func (m *Manager) SessionDestory(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(m.cookieName)

	if err == nil && cookie.Value != "" {
		m.lock.Lock()
		delete(m.sessions, cookie.Value)
		m.lock.Unlock()
	}

	cookie = &http.Cookie{
		Name:     m.cookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
}

func (m *Manager) SessionGC() {
	for now := range time.Tick(2 * time.Second) {
		m.lock.Lock()

		deleted := []string{}
		for sid, session := range m.sessions {
			if session.expiredTime.Before(now) {
				deleted = append(deleted, sid)
			}
		}

		for _, sid := range deleted {
			delete(m.sessions, sid)
		}

		m.lock.Unlock()
	}
}

var DefaultSessionManager *Manager

func init() {
	DefaultSessionManager = NewManager("sid", 3600)
}
