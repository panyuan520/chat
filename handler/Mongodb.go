package handler

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"time"
)

type UN_recive_msg struct {
	Msg         string
	Create_time time.Time
	Uid         string
	Sid         string
	Id          bson.ObjectId "_id,omitempty"
}

type Break_clients struct {
	Id          bson.ObjectId "_id,omitempty"
	Create_time time.Time
	Uid         string
}

type Message struct {
	Id              bson.ObjectId "_id,omitempty"
	Is_send         int32
	From_name       string
	From            string
	To              string
	Datetime        time.Time
	From_avatar_img string
	Msg             string
	Team_id         string
	Active          int32
}

type User struct {
	Id       bson.ObjectId "_id,omitempty"
	Username string

	First_name string
	Last_name  string
	Initials   string

	Email        string
	Password     string
	Is_staff     string
	Is_active    bool
	Is_superuser bool
	Last_login   time.Time
	Last_logout  time.Time

	Avatar_img string

	Screen_name  string
	Name         string
	Gender       string
	Age_range    string
	Game_age     int32
	Industry     string
	Industry_id  int32
	Rod_nums     string
	My_courses   []string
	Show_mylist  bool
	Device_token string
	Platform     string
	App_version  string
	Credit_score int32
	Weixin       string
	School       map[string]string
	Hobby        string
	Come_from    string
	Teams        []string
	Default_team string

	Birthday time.Time
	Address  string

	Signup_token    string
	Reset_pwd_token string
	Grade           int32
	Loc             []string
	Lock_1          []string
	Lock_2          []string
	Lock_3          []string
}

type SingleMsg struct {
	Id              bson.ObjectId "_id,omitempty"
	From            string
	From_avatar_img string
	From_name       string
	To              string
	Msg             string
	Active          int
	Cid             string
	Datetime        time.Time
}

var DB1 *mgo.Database
var DB2 *mgo.Database

type Mongodb struct {
}

func (this *Mongodb) GetUser(uid string) User {
	item := User{}
	DB1.C("user").Find(bson.M{"Id": bson.ObjectId(uid)}).One(&item)
	return item
}

func (this *Mongodb) Save(key string, data interface{}) *bson.ObjectId {
	data.Id = bson.NewObjectId()
	err := DB1.Insert(&data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return data.Id
}

func NewMongodb() *Mongodb {
	session1, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		os.Exit(1)
	}
	DB1 := session1.DB("hpgolf")
	DB2 := session1.DB("chat")
	return &Mongodb{}
}
