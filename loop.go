package go_util

import (
	"log"
	"os"
	"time"
)

func FileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}


func SafeDelay(delay time.Duration, f func()) {
	safeFunc := BuildSafeFunc(f)
	time.AfterFunc(delay, func() {
		safeFunc()
	})
}

func SafeTimer(interval time.Duration, f func(), immediately bool) {

	safeFunc := BuildSafeFunc(f)

	timer := func() {
		if immediately {
			safeFunc()
		}
		ticker := time.NewTicker(interval)
		for {
			<-ticker.C
			safeFunc()
		}
	}

	go timer()

}

func SafeGo(f func()) {
	go BuildSafeFunc(f)()
}

func BuildSafeFunc(f func()) func() {
	safeFunc := func() {
		defer func() {
			if err := recover(); nil != err {
				log.Printf("%s",err)
			}
		}()
		f()
	}
	return safeFunc
}

func SafeLoop(f func()) {
	realF := BuildSafeFunc(f)
	for {
		realF()
	}
}