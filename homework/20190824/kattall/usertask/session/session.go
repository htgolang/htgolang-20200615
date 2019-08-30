package session

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
	"sync"
	"time"
)

type Session struct {
	contents map[string]interface{}
	expiredTime time.Time
	lock sync.RWMutex
}

func NewSession(expiredTime time.Time) *Session {
	return &Session{
		contents: make(map[string]interface{}),
		expiredTime: expiredTime,
	}
}

func (s *Session) GET(key string) (interface{}, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	value, ok := s.contents[key]
	return value, ok
}

func (s *Session) SET(key string, value interface{}) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	s.contents[key] = value
}

func (s *Session) DELETE(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.contents, key)
}

type Manager struct {
	sessions   map[string]*Session
	cookieName string // sid
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
	cookie, err := r.Cookie(m.cookieName)

	m.lock.Lock()
	defer m.lock.Unlock()

	if err != nil || cookie.Value == "" {
		session = NewSession(time.Now().Add(time.Duration(m.timeout) * time.Second))
		sid = uuid.NewV4().String()
		m.sessions[sid] = session
	} else if _, ok := m.sessions[cookie.Value]; !ok {
		session = NewSession(time.Now().Add(time.Duration(m.timeout) * time.Second))
		sid = uuid.NewV4().String()
		m.sessions[sid] = session
	} else {
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

func (m *Manager) SessionGC(){
	for now := range time.Tick(2 * time.Second) {
		m.lock.Lock()
		deleted := []string{}
		for sid, session := range m.sessions {
			if session.expiredTime.Before(now) {
				deleted = append(deleted, sid)
			}
		}
		for _, sid := range deleted {
			fmt.Println("Delete: ", sid)
			delete(m.sessions, sid)
		}
		m.lock.Unlock()
	}
}

var DefaultManager *Manager

func init() {
	DefaultManager = NewManager("sid", 3600)
}
