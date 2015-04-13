package core

import (
	"github.com/op/go-logging"
	"os"
)

var log = logging.MustGetLogger("chat")

var format = logging.MustStringFormatter(
	"%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}",
)

type CLog struct {
}

func (this *CLog) Debug(msg string) {
	log.Debug(msg)
}

func (this *CLog) Info(msg string) {
	log.Info(msg)
}

func (this *CLog) Warning(msg string) {
	log.Warning(msg)
}

func (this *CLog) Notice(msg string) {
	log.Notice(msg)
}

func (this *CLog) Error(msg string) {
	log.Error(msg)
}

func NewLog() (clog *CLog) {
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	backend2Formatter := logging.NewBackendFormatter(backend2, format)

	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.ERROR, "")

	logging.SetBackend(backend1Leveled, backend2Formatter)
	return &CLog{}
}
