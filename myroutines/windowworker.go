package myroutines

import (
	"capec/bot"
	"fmt"
	"time"

	ip "capec/processer"
	wo "capec/window-operator"
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
		currentWindowProc.GetProcesses()
		newWindows := ww.initialWindowProc.GetNew(&currentWindowProc)
		if len(newWindows) == 0 {
			continue
		}
		for _, winInfo := range newWindows {
			wo.SetForegroundWindow(winInfo.Hwnd)
			wo.SetActiveWindow(winInfo.Hwnd)
			wo.BringWindowToTop(winInfo.Hwnd)
			fmt.Println(winInfo.Name, winInfo.Rect)
			//newFileName := utils.GetRandomFilename("..\\test1.png")
			time.Sleep(time.Second * 2)
			ww.bot.TakeScreenshot("..\\test1.png", &winInfo.Rect)
			response := hip.ProcessImage("..\\test1.png")
			fmt.Println(response)
			ww.pauseCh <- true
			windowClosed := false
			for _, box := range response {
				ww.bot.MoveByRelativeBox(&box, &winInfo.Rect)
				ww.bot.ClickNTimes(1)
				status, _, _ := wo.IsWindow(winInfo.Hwnd)
				if status == 1 {
					ww.initialWindowProc.AddProcessedWindow(winInfo)
					windowClosed = true
					break
				}
			}
			if !windowClosed {
				x := int(winInfo.Rect.Right+winInfo.Rect.Left) / 2
				y := int(winInfo.Rect.Bottom+winInfo.Rect.Top) / 2
				fmt.Println(x, y)
				ww.bot.MoveMouseFluent(x, y)
				ww.bot.ClickNTimes(1)
				ww.bot.Scroll()
				status, _, _ := wo.IsWindow(winInfo.Hwnd)
				if status == 1 {
					ww.initialWindowProc.AddProcessedWindow(winInfo)
					windowClosed = true
					break
				}
			}
			ww.continueCh <- true
		}
	}
}
