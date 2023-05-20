package main

import (
	"capec/bot"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	mr "capec/myroutines"
	"time"
)

func main() {
	bot := bot.CreateBot()

	pauseCh := make(chan bool)
	continueCh := make(chan bool)

	ww := mr.WindowWorker{}
	err := ww.Init(pauseCh, continueCh, &bot)

	aw := mr.AutomationScriptWorker{}
	err = aw.Init("script.bot", pauseCh, continueCh, &bot)

	fw := mr.FileWorker{}
	fw.Init("3001")

	if err != nil {
		fmt.Println(err)
		return
	}

	time.Sleep(5 * time.Second)

	go aw.Run()
	go ww.Run()
	go fw.Run()

	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	<-exit

	fmt.Println("Exiting...")

}
