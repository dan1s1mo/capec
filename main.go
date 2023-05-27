package main

import (
	"capec/bot"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	mr "capec/myroutines"
)

func main() {
	bot := bot.CreateBot()

	pauseCh := make(chan bool)
	continueCh := make(chan bool)

	ww := mr.WindowWorker{}
	err := ww.Init(pauseCh, continueCh, &bot)
	if err != nil {
		fmt.Println(err)
		return
	}
	aw := mr.AutomationScriptWorker{}
	err = aw.Init("script.bot", pauseCh, continueCh, &bot)
	if err != nil {
		fmt.Println(err)
		return
	}
	fw := mr.FileWorker{}
	err = fw.Init("3001", "files.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	go aw.Run()
	go ww.Run()
	go fw.Run()

	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	<-exit

	fmt.Println("Exiting...")

}
