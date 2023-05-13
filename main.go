package main

import (
	bot "capec/bot"
	"capec/script"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	ip "capec/processer"
	wo "capec/window-operator"
	"time"
)

func automationScriptWorker(pauseCh chan bool, continueCh chan bool, engine *script.Engine) {
	paused := false
	for {
		select {
		case <-pauseCh:
			paused = true
			fmt.Println("Bot went into image processing mode")
		case <-continueCh:
			paused = false
			fmt.Println("Bot went into script running mod")
		default:
			if !paused {
				next, _ := engine.ExecuteNextCommand()
				if !next {
					time.Sleep(10 * time.Second) // Only in development
					engine.Reset()
				}
			}
		}
	}
}

func windowManipulation(
	pauseCh chan bool,
	continueCh chan bool,
	bot bot.Bot,
	initialWindowProc wo.WinProcess,
) {
	currentWindowProc := wo.WinProcess{}
	hip := ip.HTTPImageProcesse{}
	for {
		time.Sleep(1 * time.Second)
		currentWindowProc.GetProcesses()
		fmt.Println("initialWindowProc", initialWindowProc.TotalLen())
		newWindows := initialWindowProc.GetNew(&currentWindowProc)
		if len(newWindows) == 0 {
			fmt.Println("NoWindows found")
			continue
		}
		fmt.Println("currentWindowProc", currentWindowProc.TotalLen())
		for _, winInfo := range newWindows {
			wo.SetForegroundWindow(winInfo.Hwnd)
			wo.SetActiveWindow(winInfo.Hwnd)
			wo.BringWindowToTop(winInfo.Hwnd)
			fmt.Println(winInfo.Name, winInfo.Rect)
			time.Sleep(2 * time.Second)
			bot.TakeScreenshot("..\\test1.png", &winInfo.Rect)
			response := hip.ProcessImage("..\\test1.png")
			fmt.Println(response)
			pauseCh <- true
			for _, box := range response {
				fmt.Println(response)
				bot.MoveByRelativeBox(&box, &winInfo.Rect)
				bot.ClickNTimes(1)
			}
			initialWindowProc.AddProcessedWindow(winInfo)
			continueCh <- true
		}
	}
}

func main() {
	fmt.Println("Exiting...")
	bot := bot.CreateBot()
	engine, _ := script.InitEngine("script.bot", &bot)
	pauseCh := make(chan bool)
	continueCh := make(chan bool)
	initialWindowProc := wo.WinProcess{}
	initialWindowProc.GetProcesses()
	time.Sleep(5 * time.Second)
	go automationScriptWorker(pauseCh, continueCh, engine)
	go windowManipulation(pauseCh, continueCh, &bot, initialWindowProc)

	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	<-exit

	fmt.Println("Exiting...")

}
