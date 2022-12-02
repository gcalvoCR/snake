package util

import (
	"fmt"
	"log"
	"time"

	"github.com/go-vgo/robotgo"
)


func MyFunc(dur string) {

	duration, err := time.ParseDuration(dur) //Can support ns, us, ms, s, m, h
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

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}