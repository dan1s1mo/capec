package myroutines

import (
	"capec/bot"
	"fmt"

	ip "capec/processer"
	wo "capec/window-operator"
	"time"
)

type WindowWorker struct {
	initialWindowProc wo.WinProcess
	bot               bot.Bot
	pauseCh           chan bool
	continueCh        chan bool
}

func (ww *WindowWorker) Init(
	pauseCh chan bool,
	continueCh chan bool,
	bot bot.Bot,
) error {
	ww.initialWindowProc = wo.WinProcess{}
	ww.continueCh = continueCh
	ww.pauseCh = pauseCh
	ww.bot = bot
	return ww.initialWindowProc.GetProcesses()
}

func (ww *WindowWorker) Run() {
	currentWindowProc := wo.WinProcess{}
	hip := ip.HTTPImageProcesse{}
	for {
		time.Sleep(1 * time.Second)
		currentWindowProc.GetProcesses()
		newWindows := ww.initialWindowProc.GetNew(&currentWindowProc)
		if len(newWindows) == 0 {
			continue
		}
		fmt.Println("currentWindowProc", currentWindowProc.TotalLen())
		for _, winInfo := range newWindows {
			wo.SetForegroundWindow(winInfo.Hwnd)
			wo.SetActiveWindow(winInfo.Hwnd)
			wo.BringWindowToTop(winInfo.Hwnd)
			fmt.Println(winInfo.Name, winInfo.Rect)
			time.Sleep(2 * time.Second)
			ww.bot.TakeScreenshot("..\\test1.png", &winInfo.Rect)
			response := hip.ProcessImage("..\\test1.png")
			fmt.Println(response)
			ww.pauseCh <- true
			for _, box := range response {
				fmt.Println(response)
				ww.bot.MoveByRelativeBox(&box, &winInfo.Rect)
				ww.bot.ClickNTimes(1)
			}
			ww.initialWindowProc.AddProcessedWindow(winInfo)
			ww.continueCh <- true
		}
	}
}
