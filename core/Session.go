package core

import (
	"errors"
	"net"
)

var clog *CLog

type Session struct {
	net.Conn
	callback       func(session *Session)
	Id             int64
	sign           chan bool
	readErrorChan  chan bool
	readZeroChan   chan bool
	writeErrorChan chan bool
}

func (this *Session) listener() {
	defer this.Cclose()
	for {
		select {
		case <-this.sign:
			return
		case <-this.readErrorChan:
			return
		case <-this.readZeroChan:
			return
		case <-this.writeErrorChan:
			return
		default:
			this.callback(this)
		}
	}
}

func (this *Session) Cread(hlen int, blen int) (string, error) {
	head := make([]byte, hlen)
	tlen, err := this.Read(head)
	if err != nil {
		this.readErrorChan <- true
		return "", errors.New("read error")
	}
	if tlen != hlen {
		this.readZeroChan <- true
		return "", errors.New("read num != hlen error")
	}
	length, err := bti(head)
	if err != nil {
		this.readZeroChan <- true
		return "", errors.New("read head error")
	}
	size := make([]byte, hlen)
	body := make([]byte, blen)
	total := 0
	for total < length {
		rest := length - total
		if blen > rest {
			size = make([]byte, rest)
		}
		if l, e := this.Read(size); l > 0 && e == nil {
			total += l
			body = append(body, size...)
		} else {
			this.readZeroChan <- true
			clog.Error(e.Error())
			return "", errors.New("read zero error")
		}
	}
	return string(body), nil
}

func (this *Session) Cwrite(message string) {
	body, err := encode(message)
	if err != nil {
		clog.Error(err.Error())
		this.writeErrorChan <- true
	} else {
		if _, err := this.Write(body); err != nil {
			clog.Error(err.Error())
			this.writeErrorChan <- true
		}
	}
}

func (this *Session) Cclose() {
	this.sign <- true
	this.Close()
}

func NewSession(listener net.Conn, callback func(session *Session)) (session *Session) {
	clog = NewLog()
	return &Session{listener, callback, NewSessionId(), make(chan bool, 1), make(chan bool, 1), make(chan bool, 1), make(chan bool, 1)}
}
