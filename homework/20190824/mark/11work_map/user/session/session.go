package session

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
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
	cookieName string // sid
	timeout    int    // 3600s
	lock       sync.Mutex
}

func NewManager(cookieName string, timeout int) *Manager {
	manager := &Manager{
		sessions:   make(map[string]*Session),
		cookieName: "sid",
		timeout:    timeout,
	}
	go manager.SessionGC()
	return manager
}
func (m *Manager) SessionStart(w http.ResponseWriter, r *http.Request) *Session {
	var session *Session
	var sid string
	cookie, err := r.Cookie(m.cookieName)

	m.lock.Lock()
	defer m.lock.Unlock() //  延迟释放锁

	// 检查并生成session的过程
	if err != nil || cookie.Value == "" {
		sid = uuid.NewV4().String()
		session = NewSession(time.Now().Add(time.Duration(m.timeout) * time.Second))
		m.sessions[sid] = session
	} else if _, ok := m.sessions[cookie.Value]; !ok { // 伪造的sessions，或者服务重启
		sid = uuid.NewV4().String()
		session = NewSession(time.Now().Add(time.Duration(m.timeout) * time.Second))
		m.sessions[sid] = session
	} else {
		// 正常情况
		return m.sessions[cookie.Value]
	}
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
		m.lock.Lock() //  加锁
		delete(m.sessions, cookie.Value)
		m.lock.Unlock() //  释放锁
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

var DefaultManager *Manager

func (m *Manager) SessionGC() {
	// 每隔多长时间执行一次代码
	for noew := range time.Tick(2 * time.Second) {
		m.lock.Lock() // 加锁
		deleted := []string{}
		for sid, sessions := range m.sessions {
			if sessions.expiredTime.Before(noew) {
				deleted = append(deleted, sid)
			}
		}
		//fmt.Println("deleted")
		for _, sid := range deleted {
			fmt.Println(sid)
			delete(m.sessions, sid)
		}
		m.lock.Unlock()
	}
}
func init() {
	DefaultManager = NewManager("sid", 3600) // session 时间
}
