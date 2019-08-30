package session

import (
	"net/http"
	"sync"
	"time"

	"github.com/satori/go.uuid"
)

var DefaultSessionManager *SessionManager

type Session struct {
	id          string
	datas       map[string]interface{}
	expiredTime time.Time
}

func NewSession(id string, expiredTime time.Time) *Session {
	return &Session{
		id:          id,
		datas:       make(map[string]interface{}),
		expiredTime: expiredTime,
	}
}

func (s *Session) Get(key string) (value interface{}, ok bool) {
	value, ok = s.datas[key]
	return
}

func (s *Session) Set(key string, value interface{}) {
	s.datas[key] = value
}

func (s *Session) Delete(key string) {
	delete(s.datas, key)
}

type SessionManager struct {
	cookieName string
	sessions   map[string]*Session
	timeout    time.Duration
	mutex      sync.RWMutex
}

func NewSessionManager() *SessionManager {
	mgr := &SessionManager{
		cookieName: "sid",
		sessions:   make(map[string]*Session),
		timeout:    time.Duration(time.Second * 10),
	}
	go mgr.gc()
	return mgr
}

func (p *SessionManager) StartSession(w http.ResponseWriter, r *http.Request) *Session {
	cookie, err := r.Cookie(p.cookieName)
	if err == nil && cookie.Value != "" {
		if session, ok := p.Load(cookie.Value); ok {
			return session
		}
	}

	id := p.Gid()
	expired := time.Now().Add(p.timeout)
	session := NewSession(id, expired)
	p.mutex.Lock()
	p.sessions[id] = session
	p.mutex.Unlock()

	cookie = &http.Cookie{
		Name:     p.cookieName,
		Value:    id,
		Path:     "/",
		Expires:  expired,
		MaxAge:   int(p.timeout / time.Second),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	return session

}

func (p *SessionManager) DesotrySession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(p.cookieName)
	if err == nil && cookie.Value != "" {
		if session, ok := p.Load(cookie.Value); ok {
			p.Destory(session.id)
		}
	}

	cookie = &http.Cookie{
		Name:     p.cookieName,
		Value:    "",
		Path:     "/",
		Expires:  time.Now(),
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
}

func (p *SessionManager) Gid() string {
	return uuid.NewV4().String()
}

func (p *SessionManager) Load(id string) (session *Session, ok bool) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	session, ok = p.sessions[id]
	return
}

func (p *SessionManager) Destory(id string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	delete(p.sessions, id)
}

func (p *SessionManager) gc() {
	for now := range time.Tick(time.Second * 3) {
		expired := []string{}
		for id, session := range p.sessions {
			if session.expiredTime.Before(now) {
				expired = append(expired, id)
			}
		}
		p.mutex.Lock()
		for _, id := range expired {
			delete(p.sessions, id)
		}
		p.mutex.Unlock()
	}
}

func init() {
	DefaultSessionManager = NewSessionManager()
}
