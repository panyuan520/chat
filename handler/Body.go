package handler

import (
	"bytes"
	"fmt"
	"github.com/bitly/go-simplejson"
	"time"
)

var mongo *Mongodb

type Body struct {
}

func (this *Body) Map(content string) (*simplejson.Json, error) {
	b := []byte(content)
	b = bytes.Trim(b, "\x00")
	return simplejson.NewJson([]byte(b))
}

func (this *Body) Login(j *simplejson.Json) {
	fmt.Println("login")

}

func (this *Body) Single(j *simplejson.Json) {
	uid, _ := j.Get("from").String()
	opt_uid, _ := j.Get("to").String()
	msg, _ := j.Get("msg").String()
	datetime, _ := j.Get("date").String()
	cid, _ := j.Get("cid").String()
	user := mongo.GetUser(uid)
	singleMsg := &SingleMsg{nil, uid, user.Avatar_img, user.Name, opt_uid, msg, 1, cid, time.Now()}
	mongo.Save(singleMsg)

}

func (this *Body) Heartbeat(j *simplejson.Json) {
	fmt.Println("Heartbeat")
}

func (this *Body) Group(j *simplejson.Json) {
	fmt.Println("Group")
}

func (this *Body) Logout(j *simplejson.Json) {
	fmt.Println("Logout")
}

func (this *Body) Confirm(j *simplejson.Json) {
	fmt.Println("Confirm")
}

func NewBody() *Body {
	mongo := NewMongodb()
	return &Body{}
}
