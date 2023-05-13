package script

import (
	t "capec/types"
	u "capec/utils"
	"fmt"
)

type BlockParser struct {
	index  int
	tokens []t.Token
	Block  map[string]([]t.Command)
}

func (p *BlockParser) findBLock() (*t.Block, int, error) {
	if p.tokens[0].Tpe != BLOCK.typ {
		return nil, 0, fmt.Errorf("expected to start the block with %s keyword on position %d", BLOCK.re, p.tokens[0].Position)
	}
	if p.tokens[1].Tpe != COLUMN.typ {
		return nil, 0, fmt.Errorf("expected token with types %s on position %d, but received %s", COLUMN.typ, p.tokens[0].Position+1, p.tokens[1].Text)
	}
	if p.tokens[2].Tpe != TEXT.typ {
		return nil, 0, fmt.Errorf("expected token with types %s on position %d, but received %s", TEXT.typ, p.tokens[1].Position+1, p.tokens[2].Text)
	}
	if p.tokens[3].Tpe != LBRACKET.typ {
		return nil, 0, fmt.Errorf("expected token with types %s on position %d, but received %s", LBRACKET.typ, p.tokens[2].Position+1, p.tokens[3].Text)
	}
	_, iRb, err := p.findFirstTokenOfType(RBRACKET.typ)
	if err != nil {
		return nil, 0, fmt.Errorf("expected token \"}\" to close \"{\" on position %d", p.tokens[3].Position)
	}
	blockTokens := u.Filter(p.tokens[4:iRb], func(t t.Token) bool {
		isTabulation := t.Tpe == NEWLINE.typ || t.Tpe == SPACE.typ || t.Tpe == TAB.typ
		return !isTabulation
	})
	blockName := p.tokens[2].Text
	if len(p.Block[blockName]) != 0 {
		return nil, 0, fmt.Errorf("block with name \"%s\" is already created", blockName)
	}
	contentParser := Parser{1, blockTokens, make([]t.Command, 0)}
	content, err := contentParser.BlockContent(u.GetMapsKeys(p.Block))
	if err != nil {
		return nil, 0, err
	}
	newBlock := t.Block{Name: blockName, Content: content}
	return &newBlock, iRb + 1, nil
}

func (p *BlockParser) findFirstTokenOfType(tpe string) (*t.Token, int, error) {
	for i, token := range p.tokens {
		if token.Tpe == tpe {
			return &token, i, nil
		}
	}
	return nil, -1, fmt.Errorf("No token of type " + tpe + " found")
}

func (p *BlockParser) run() error {
	i := 0
	for {
		if len(p.tokens) == 0 {
			break
		}
		if p.tokens[i].Tpe == NEWLINE.typ || p.tokens[i].Tpe == SPACE.typ || p.tokens[i].Tpe == TAB.typ {
			i++
			continue
		}
		if p.tokens[i].Tpe == BLOCK.typ {
			p.tokens = p.tokens[i:len(p.tokens)]
			i = 0
			newBlock, index, err := p.findBLock()
			if err != nil {
				return err
			}
			if index == len(p.tokens)+1 {
				index = len(p.tokens)
			}
			p.tokens = p.tokens[index:len(p.tokens)]
			p.Block[newBlock.Name] = newBlock.Content
			continue
		}
	}
	if len(p.Block[MAIN_BLOCK]) == 0 {
		return fmt.Errorf("block with name \"%s\" is an entry point of a script and its is mandatory to be defined", MAIN_BLOCK)
	}
	return nil
}
