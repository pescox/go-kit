package task

import (
	"testing"
	"time"

	"github.com/pescox/go-kit/log"
)

func TestDoInterval(t *testing.T) {
	DoInterval(*newIntervalOption(time.Second, 5), func() {
		log.InfoF("start1")
		time.Sleep(time.Second * 10)
		log.InfoF("end1")
	})

	DoInterval(*newIntervalOption(time.Second*4, 5), func() {
		log.InfoF("start2")
		time.Sleep(time.Second)
		log.InfoF("end2")
	})
}

func TestDoTimeout(t *testing.T) {
	DoTimeout(time.Second*2, func() {
		log.Info("start1")
		log.Info("end1")
	})

	DoTimeout(time.Second*2, func() {
		log.Info("start2")
		time.Sleep(time.Second * 5)
		log.Info("end2")
	})
}

func TestDoAfter(t *testing.T) {
	log.Info("test start")
	DoAfter(time.Second*2, func() {
		log.Info("do start")
	})
}
