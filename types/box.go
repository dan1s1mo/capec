package types

type Box struct {
	Text   string `json:"text"`
	Left   int    `json:"left"`
	Top    int    `json:"top"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func (b *Box) GetCenter() (int, int) {
	x := b.Left + (b.Width / 2)
	y := b.Top + (b.Height / 2)
	return x, y
}
