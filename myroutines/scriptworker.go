package myroutines

import (
	"capec/bot"
	"capec/script"
	"fmt"

	"time"
)

type AutomationScriptWorker struct {
	pauseCh    chan bool
	continueCh chan bool
	engine     *script.Engine
}

func (aw *AutomationScriptWorker) Init(
	fileName string,
	pauseCh chan bool,
	continueCh chan bool,
	bot bot.Bot,
) error {
	aw.continueCh = continueCh
	aw.pauseCh = pauseCh
	engine, err := script.InitEngine("script.bot", bot)
	aw.engine = engine
	return err
}

func (aw *AutomationScriptWorker) Run() {
	paused := false
	for {
		select {
		case <-aw.pauseCh:
			paused = true
			fmt.Println("Bot went into image processing mode")
		case <-aw.continueCh:
			paused = false
			fmt.Println("Bot went into script running mod")
		default:
			if !paused {
				next, _ := aw.engine.ExecuteNextCommand()
				time.Sleep(1 * time.Second)
				if !next {
					fmt.Println("stoped")
					time.Sleep(1 * time.Second) // Only in development
					aw.engine.Reset()
				}
			}
		}
	}
}
