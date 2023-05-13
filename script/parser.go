package script

import (
	t "capec/types"
	u "capec/utils"
	"fmt"
)

type Parser struct {
	index   int
	tokens  []t.Token
	content []t.Command
}

func (p *Parser) findFirstTokenOfType(tpe string) (*t.Token, int, error) {
	for i, token := range p.tokens {
		if token.Tpe == tpe {
			return &token, i, nil
		}
	}
	return nil, -1, fmt.Errorf("No token of type " + tpe + " found")
}

func (p *Parser) formCommand(blocks []string) (int, t.Command, error) {
	comm := make(t.Command, 0)
	comm = append(comm, p.tokens[0])
	tOp, iOp, err := p.findFirstTokenOfType(OPARENTHESIS.typ)
	if err != nil || iOp != 1 {
		text := ""
		if p.index+1 <= len(p.tokens) {
			text = p.tokens[p.index+1].Text
		}
		return 0, nil, fmt.Errorf("expected token \"(\" on position %d, but received %s", (comm[0].Position + 1), text)
	}
	_, iCp, err := p.findFirstTokenOfType(CPARENTHESIS.typ)
	if err != nil {
		return 0, nil, fmt.Errorf("expected token \")\" to close \"(\" on position %d", tOp.Position)
	}

	for i := iOp + 1; i < iCp; i++ {
		token := p.tokens[i]
		if token.Tpe == NUMBER.typ || token.Tpe == TEXT.typ {
			if comm[0].Tpe == REPEAT.typ {
				if len(comm) == 1 {
					if token.Tpe != TEXT.typ {
						return 0, nil, fmt.Errorf("expected block name on position %d, but received %s, position: %d", token.Position, token.Text, comm[0].Position)
					}
					if token.Tpe == TEXT.typ && !u.Contains(blocks, token.Text) {
						return 0, nil, fmt.Errorf("cant use block with name %s dose not exists or not initiated, position: %d", token.Text, comm[0].Position)
					}
				}
				if len(comm) == 2 {
					if token.Tpe != NUMBER.typ {
						return 0, nil, fmt.Errorf("expected token with type %s on position %d, but received %s, position: %d", NUMBER.typ, token.Position, token.Text, comm[0].Position)
					}
				}
			}
			if comm[0].Tpe == CLICK.typ {
				if token.Tpe != NUMBER.typ && len(comm) == 1 {
					return 0, nil, fmt.Errorf("expected token with type %s on position %d, but received %s, position: %d", NUMBER.typ, token.Position, token.Text, comm[0].Position)
				}
				if len(comm) == 2 {
					return 0, nil, fmt.Errorf("function %s supports only one argument - amount of clicks, position: %d", CLICK.re, comm[0].Position)
				}
			}
			if comm[0].Tpe == MOVE.typ {
				if token.Tpe != NUMBER.typ && len(comm) <= 2 {
					return 0, nil, fmt.Errorf("expected token with type %s on position %d, but received %s, position: %d", NUMBER.typ, token.Position, token.Text, comm[0].Position)
				}
				if len(comm) == 3 {
					return 0, nil, fmt.Errorf("function %s supports only one argument - amount of clicks, position: %d", CLICK.re, comm[0].Position)
				}
			}

			comm = append(comm, token)
		}
	}
	if comm[0].Tpe == CLICK.typ && len(comm) != 2 {
		return 0, nil, fmt.Errorf("function %s expects to have one argument - amount of clicks, position: %d", CLICK.re, comm[0].Position)
	}
	if comm[0].Tpe == MOVE.typ && len(comm) != 3 {
		return 0, nil, fmt.Errorf("function %s expects to have two argument - x and y coordinates, position: %d", CLICK.re, comm[0].Position)
	}
	if comm[0].Tpe == REPEAT.typ && len(comm) != 3 {
		return 0, nil, fmt.Errorf("function %s expects to have two argument - block name and amount of repeats, position: %d", CLICK.re, comm[0].Position)
	}
	length := 3 + iCp - iOp
	return length, comm, nil
}

func (p *Parser) BlockContent(blocks []string) ([]t.Command, error) {
	for {
		if len(p.tokens) == 0 {
			break
		}
		index, comm, err := p.formCommand(blocks)
		if err != nil {
			return nil, err
		}
		p.content = append(p.content, comm)
		if index == len(p.tokens)+1 {
			index = len(p.tokens)
		}
		p.tokens = p.tokens[index:len(p.tokens)]
	}
	return p.content, nil
}
