package main

import (
	. "chat/core"
	. "chat/handler"
	"fmt"
	"runtime"
)

var hand *Body
var sessionManger *SessionManger

func callback(session *Session) {
	if body, err := session.Cread(4, 1024); err != nil {
		fmt.Println(fmt.Sprintf("callback read error:%s", err.Error()))
	} else {
		if message, err := hand.Map(body); err != nil {
			fmt.Println(fmt.Sprintf("hand map:%s", err.Error()))
		} else {
			head_type, err := message.Get("type").Int()
			if err != nil {
				fmt.Println("head_type:", err.Error())
			} else {
				switch head_type {
				case 0:
					hand.Login(message)
				case 1:
					hand.Single(message)
				case 2:
					hand.Heartbeat(message)
				case 3:
					hand.Group(message)
				case 4:
					hand.Logout(message)
				case 5:
					hand.Confirm(message)
				}
			}
		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() + 1)
	//handler body
	hand = &Body{}

	sessionManger = &SessionManger{make(map[int64]*Session)}
	serverEngine := &ServerEngine{sessionManger, callback}
	serverEngine.Start()
}
