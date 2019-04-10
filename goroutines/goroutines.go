package main

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("test %s", s)
	}
}

func main() {
	log.Warnf("sff//%s", "test")
	go say("world")
	say("hello")
}
