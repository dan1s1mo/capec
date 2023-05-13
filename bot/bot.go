package bot

import (
	"capec/types"
	wo "capec/window-operator"
	"fmt"

	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/imgo"
)

type Bot interface {
	MoveByRelativeBox(textBox *types.Box, windowBox *wo.TagRECT) error
	MoveMouseFluent(x int, y int) error
	ClickNTimes(n int) error
	TakeScreenshot(path string, windowBox *wo.TagRECT) error
	Scroll() error
}

type BotVOne struct {
	X int
	Y int
}

func (b *BotVOne) MoveMouseFluent(x int, y int) error {
	valid := x < b.X && y < b.Y
	if !valid {
		return nil
	}
	robotgo.MoveSmooth(x, y, 0.1, 1.0)
	return nil
}

func CreateBot() BotVOne {
	x, y := robotgo.GetScreenSize()
	return BotVOne{X: x, Y: y}
}

func (b *BotVOne) MoveByRelativeBox(textBox *types.Box, windowBox *wo.TagRECT) error {
	x, y := textBox.GetCenter()
	X := x + int(windowBox.Left)
	Y := y + int(windowBox.Top)
	fmt.Printf("move to", X, Y, "\n")
	b.MoveMouseFluent(X, Y)
	return nil
}

func (b *BotVOne) Scroll() error {
	robotgo.ScrollSmooth(-10, 6, 200, -10)
	return nil
}

func (b *BotVOne) TakeScreenshot(path string, Rect *wo.TagRECT) error {
	bit := robotgo.CaptureScreen(
		int(Rect.Left),
		int(Rect.Top),
		int(Rect.Right-Rect.Left),
		int(Rect.Bottom-Rect.Top),
	)
	img := robotgo.ToImage(bit)
	defer robotgo.FreeBitmap(bit)
	err := imgo.Save(path, img)
	return err
}

func (b *BotVOne) ClickNTimes(n int) error {
	for i := 0; i < n; i++ {
		robotgo.Click()
	}
	return nil
}
