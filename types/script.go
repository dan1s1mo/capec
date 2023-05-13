package types

type Command []Token

type Block struct {
	Name    string
	Content []Command
}

type Token struct {
	Position int
	Tpe      string
	Text     string
}
