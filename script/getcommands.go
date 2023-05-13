package script

import (
	bot "capec/bot"
	t "capec/types"
	u "capec/utils"
	"io/ioutil"
)

func InitEngine(path string, bot bot.Bot) (*Engine, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	text := string(data)

	l := Lexer{}
	tokens := l.GetTokens(text)
	a := make(map[string]([]t.Command), 0)
	iT := u.Filter(tokens, func(t t.Token) bool {
		isTabulation := t.Tpe == NEWLINE.typ || t.Tpe == SPACE.typ || t.Tpe == TAB.typ
		return !isTabulation
	})

	p := BlockParser{1, iT, a}
	err = p.run()
	if err != nil {
		return nil, err
	}
	e := Engine{}
	e.Init(p.Block, bot)
	return &e, nil
}
