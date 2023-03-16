package windowoperator

import (
	"fmt"
	utils "tapok/utils"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	enumWindowsProc = windows.NewCallback(func(hwnd windows.HWND, lparam uintptr) uintptr {
		winProc := (*WinProcess)(unsafe.Pointer(lparam))
		if !windows.IsWindowVisible(hwnd) {
			return 1
		}
		rect, _ := GetWindowRect(hwnd)
		if rect.Top == 0 {
			return 1
		}
		name, _ := GetWindowText(hwnd, 200)
		if len(name) < 3 {
			return 1
		}
		winInfo := WinInfo{hwnd, rect, name}
		winProc.process = append(winProc.process, winInfo)
		return 1
	})
)

type TagRECT struct {
	Left, Top, Right, Bottom int32
}

type WinInfo struct {
	Hwnd windows.HWND
	Rect TagRECT
	Name string
}

type WinProcess struct {
	process []WinInfo
}

func (wp *WinProcess) GetNew(current_wp *WinProcess) []WinInfo {
	return utils.Filter(current_wp.process, func(wi WinInfo) bool {
		for _, v := range wp.process {
			if v.Hwnd == wi.Hwnd {
				return false
			}
		}
		return true
	})
}

func (wp *WinProcess) GetProcesses() error {
	wp.process = []WinInfo{}
	err := windows.EnumWindows(enumWindowsProc, unsafe.Pointer(wp))
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}

func (wp *WinProcess) TotalLen() int {
	return len(wp.process)
}
