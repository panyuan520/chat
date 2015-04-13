package handler

import (
	"github.com/gosexy/redis"
	"os"
)

var client *redis.Client

type Cache struct {
}

func (this *Cache) Cpush(key, body string) {
	client.Command(nil, "rpush", key, body)
}

func (this *Cache) Clen(pointer string) int32 {
	var i int32
	err := client.Command(&i, "llen", pointer)
	if err != nil {
		return 0
	}
	return i
}

func (this *Cache) Chget(pointer, key string) string {
	var i string
	err := client.Command(&i, "hget", pointer, key)
	if err != nil {
		return ""
	}
	return i
}

func (this *Cache) Chset(pointer, key, value string) {
	client.Command(nil, "hset", pointer, key, value)
}

func (this *Cache) Chlen(pointer string) int32 {
	var i int32
	err := client.Command(&i, "hlen", pointer)
	if err != nil {
		return 0
	}
	return i
}

func (this *Cache) Chexists(pointer, key string) bool {
	var i bool
	err := client.Command(&i, "hexists", pointer, key)
	if err != nil {
		return false
	}
	return i
}

func (this *Cache) Chkeys(pointer, key string) []string {
	var i []string
	err := client.Command(&i, "hkeys", pointer, key)
	if err != nil {
		return []string{}
	}
	return i
}

func (this *Cache) Chdel(pointer, key string) {
	client.Command(nil, "hdel", pointer, key)
}

func NewCache() *Cache {
	client = redis.New()
	err := client.Connect("localhost", 2701)
	if err != nil {
		os.Exit(1)
	}
	return &Cache{}
}
