package core

import (
	"sync"
)

var lock sync.Mutex

type SessionManger struct {
	Smap map[int64]*Session
}

func (this *SessionManger) Add(session *Session) {
	defer lock.Unlock()
	lock.Lock()
	this.Smap[NewSessionId()] = session
}

func (this *SessionManger) Get(id int64) (session *Session) {
	defer lock.Unlock()
	lock.Lock()
	return this.Smap[id]
}

func (this *SessionManger) Remove(id int64) {
	defer lock.Unlock()
	lock.Lock()
	delete(this.Smap, id)
}
