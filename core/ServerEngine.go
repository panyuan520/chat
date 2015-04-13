package core

import (
	"fmt"
	"net"
	"os"
)

type ServerEngine struct {
	Manger   *SessionManger
	Callback func(session *Session)
}

func (this *ServerEngine) Start() {
	clog.Info("serve start....")
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP("0.0.0.0"), 32000, ""})
	if err != nil {
		clog.Error(err.Error())
		os.Exit(1)
	}
	defer listener.Close()
	for {
		c, error := listener.Accept()
		if error != nil {
			clog.Error(fmt.Sprintf("error Accept:", err.Error()))
			break
		}
		session := NewSession(c, this.Callback)
		this.Manger.Smap[session.Id] = session
		go session.listener()
	}
}
