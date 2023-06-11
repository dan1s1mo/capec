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
		time.Sleep(time.Second * 5)
		currentWindowProc.GetProcesses()
		newWindows := ww.initialWindowProc.GetNew(&currentWindowProc)
		if len(newWindows) == 0 {
			continue
		}
		for _, winInfo := range newWindows {
			ww.pauseCh <- true
			wo.SetForegroundWindow(winInfo.Hwnd)
			wo.SetActiveWindow(winInfo.Hwnd)
			wo.BringWindowToTop(winInfo.Hwnd)
			fmt.Println(winInfo.Name, winInfo.Rect)
			windowClosed := false
			childWindows := wo.WinProcess{}
			childWindows.GetChildWindows(winInfo.Hwnd)

			for _, btn := range childWindows.Process() {
				fmt.Println(btn.Rect.Right-btn.Rect.Left, btn.Rect.Bottom-btn.Rect.Top)
				x := int(btn.Rect.Right+btn.Rect.Left) / 2
				y := int(btn.Rect.Bottom+btn.Rect.Top) / 2
				ww.bot.MoveMouseFluent(x, y)
				ww.bot.ClickNTimes(1)
				status, _, _ := wo.IsWindow(winInfo.Hwnd)
				if status == 0 {
					time.Sleep(time.Millisecond * 500)
					ww.initialWindowProc.AddProcessedWindow(winInfo)
					windowClosed = true
					break
				}
			}

			if !windowClosed {
				time.Sleep(time.Second * 2)
				ww.bot.TakeScreenshot("..\\test1.png", &winInfo.Rect)
				response := hip.ProcessImage("..\\test1.png")
				fmt.Println(response)
				ww.pauseCh <- true

				for _, box := range response {
					ww.bot.MoveByRelativeBox(&box, &winInfo.Rect)
					ww.bot.ClickNTimes(1)
					status, _, _ := wo.IsWindow(winInfo.Hwnd)
					if status == 0 {
						ww.initialWindowProc.AddProcessedWindow(winInfo)
						windowClosed = true
						break
					}
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
				if status == 0 {
					ww.initialWindowProc.AddProcessedWindow(winInfo)

				}
			}
			ww.continueCh <- true
		}
	}
}
