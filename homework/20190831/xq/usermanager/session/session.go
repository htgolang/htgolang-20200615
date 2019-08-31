package session

import (

	"net/http"
	"github.com/satori/go.uuid"
	"sync"
	"time"
)

type Session struct {
	contents    map[string]interface{}
	expiredTime time.Time
	lock    sync.RWMutex
}

func NewSession(expiredTime time.Time) *Session {
	return &Session{
		contents:    make(map[string]interface{}),
        expiredTime: expiredTime, 
        }

}

func (s *Session) Set(key string, value interface{})  {
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

func (s *Session) Delete(key string){

	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.contents, key)
}


type Manager struct {
	sessions map[string]*Session
	cookieName string //sid
	timeout int
	lock sync.Mutex
}

func NewManager(cookieName string, timeout int) *Manager {
 	manager := &Manager{
 		sessions: make(map[string]*Session),
		cookieName: cookieName, //sid
		timeout: timeout,

	}
 	go manager.SessionGC()
 	return manager
}

func (m *Manager) SessionStart(responseWriter http.ResponseWriter, request *http.Request) *Session {
	var session *Session
	var sid string
	cookie, err := request.Cookie(m.cookieName)

	m.lock.Lock()
	defer m.lock.Unlock()

	if err != nil || cookie.Value == ""{

		sid = uuid.Must(uuid.NewV4()).String()
		session = NewSession(time.Now().Add(time.Duration(m.timeout) * time.Second))

		m.sessions[sid] = session

	} else if _,ok := m.sessions[cookie.Value]; !ok {  // 伪造
		sid = uuid.Must(uuid.NewV4()).String()
		//session = NewSession()
		session = NewSession(time.Now().Add(time.Duration(m.timeout) * time.Second))

		m.sessions[sid] = session
	}else {
		return m.sessions[cookie.Value]

	}
	cookie = &http.Cookie{
		Name: m.cookieName,
		Value: sid,
		Path: "/",
		MaxAge: m.timeout,
		HttpOnly:true,
	}
	http.SetCookie(responseWriter, cookie)
	return session
}


func (m *Manager) SessionDestory(responseWriter http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie(m.cookieName)
	if err == nil || cookie.Value != ""{
	    m.lock.Lock()
	    defer m.lock.Unlock()

		delete(m.sessions, cookie.Value)
	}

	cookie = &http.Cookie{
		Name: m.cookieName,
		Value: "logout",
		Path: "/",
		MaxAge: -1,
		HttpOnly:true,
	}
	http.SetCookie(responseWriter, cookie)

}

func (m *Manager) SessionGC() {
	for now := range time.Tick(2 * time.Second) {
		m.lock.Lock()

		deleted := []string{}
		for sid, session := range m.sessions {
			if session.expiredTime.Before(now){
				deleted = append(deleted, sid)
			}
		}
		for _, sid := range deleted {
			delete(m.sessions, sid)
		}
		m.lock.Unlock()
	}
}

var DefaultManager *Manager

func init(){
	//DefaultManager = &Manager{
	//	sessions: make(map[string]*Session),
	//	cookieName:"sid",
	//	timeout:3600,
	//}

	DefaultManager = NewManager("sid", 7200)
}
