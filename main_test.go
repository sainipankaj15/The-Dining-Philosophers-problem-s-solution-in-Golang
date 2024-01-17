package main

import (
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	eatingTime = 0 * time.Second
	thinkingTime = 0 * time.Second
	sleepTime = 0 * time.Second

	for i := 0; i < 100000; i++ {
		main()
		if len(whoFinishedFirst) != 3 {
			t.Error("wrong number of entries in slice")
		}

		whoFinishedFirst = []string{}
	}
}
