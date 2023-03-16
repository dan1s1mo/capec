package main

import (
	"fmt"

	wo "tapok/window-operator"

	"github.com/go-vgo/robotgo"
)

func main() {
	initialWindowProc := wo.WinProcess{}
	initialWindowProc.GetProcesses()
	currentWindowProc := wo.WinProcess{}
	x, y := robotgo.GetScreenSize()
	fmt.Println(x, y)
	for ok := true; ok; {
		robotgo.Sleep(1)
		currentWindowProc.GetProcesses()
		newWindows := initialWindowProc.GetNew(&currentWindowProc)
		for _, winInfo := range newWindows {

			wo.SetForegroundWindow(winInfo.Hwnd)
			wo.SetActiveWindow(winInfo.Hwnd)
			wo.BringWindowToTop(winInfo.Hwnd)
			fmt.Println(winInfo.Name, winInfo.Rect)
		}
	}

}
