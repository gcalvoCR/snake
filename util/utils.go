package util

import (
	"fmt"
	"log"
	"time"

	"github.com/go-vgo/robotgo"
)


func MyFunc() {

	duration, err := time.ParseDuration("30s") //Can support ns, us, ms, s, m, h
	if err != nil {
		log.Fatalln("Your time duration could not be understand")
	}
	start := time.Now()
	t := start
	for t.Sub(start) < duration {
		robotgo.Scroll(10, 10)
		robotgo.Scroll(20, -10)
		t = time.Now()
		fmt.Println(t.Sub(start))
	}

}