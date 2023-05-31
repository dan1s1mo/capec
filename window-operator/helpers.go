package windowoperator

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	user32                  = syscall.MustLoadDLL("user32.dll")
	procGetWindowTextW      = user32.MustFindProc("GetWindowTextW")      //done
	procBringWindowToTop    = user32.MustFindProc("BringWindowToTop")    //done
	procSetActiveWindow     = user32.MustFindProc("SetActiveWindow")     //done
	procSetForegroundWindow = user32.MustFindProc("SetForegroundWindow") //done
	procGetWindowRect       = user32.MustFindProc("GetWindowRect")       //done
	procGetClassNameW       = user32.MustFindProc("GetClassNameW")
	procEnumChildWindows    = user32.MustFindProc("EnumChildWindows")
	procIsWindow            = user32.MustFindProc("IsWindow")
)

func GetWindowText(hwnd windows.HWND, maxCount int32) (str string, err error) {
	b := make([]uint16, maxCount)
	r0, _, e1 := procGetWindowTextW.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&b[0])), uintptr(maxCount))
	len := int32(r0)
	if len == 0 {
		if e1 != nil {
			err = e1
		} else {
			err = syscall.EINVAL
		}
		return
	}
	str = syscall.UTF16ToString(b)
	return
}

func GetClassName(hwnd windows.HWND, maxCount int32) (str string, err error) {
	b := make([]uint16, maxCount)
	r0, _, e1 := procGetClassNameW.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&b[0])), uintptr(maxCount))
	len := int32(r0)
	if len == 0 {
		if e1 != nil {
			err = e1
		} else {
			err = syscall.EINVAL
		}
		return
	}
	str = syscall.UTF16ToString(b)
	return
}

func BringWindowToTop(hwnd windows.HWND) (err error) {
	r0, _, e1 := procBringWindowToTop.Call(uintptr(hwnd))
	if r0 == 0 {
		if e1 != nil {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetForegroundWindow(hwnd windows.HWND) (err error) {
	r0, _, e1 := procSetForegroundWindow.Call(uintptr(hwnd))
	if r0 == 0 {
		if e1 != nil {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetActiveWindow(hwnd windows.HWND) (prewHwnd windows.HWND, err error) {
	r0, _, e1 := procSetActiveWindow.Call(uintptr(hwnd))
	prewHwnd = windows.HWND(r0)
	if prewHwnd == 0 {
		if e1 != nil {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetWindowRect(hwnd windows.HWND) (rect TagRECT, err error) {
	r0, _, e1 := procGetWindowRect.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&rect)))
	if r0 == 0 {
		if e1 != nil {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func IsWindow(hwnd windows.HWND) (uintptr, uintptr, error) {
	r0, o, e1 := procIsWindow.Call(uintptr(hwnd))
	return r0, o, e1
}
