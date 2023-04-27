package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(10).Seconds().Do(func() {
		fmt.Println("to the terminal....")
	})

	s.StartBlocking()

}
