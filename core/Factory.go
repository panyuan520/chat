package core

import (
	"sync"
)

var slock sync.RWMutex

var id int64

func NewSessionId() int64 {
	defer slock.Unlock()
	slock.Lock()
	id++
	return id

}
