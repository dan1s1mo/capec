package script

import (
	bot "capec/bot"
	t "capec/types"
	"strconv"
)

type Engine struct {
	bot          bot.Bot
	blocks       map[string][]t.Command
	loopCount    int
	contextStack Stack
}

func (e *Engine) ExecuteNextCommand() (bool, error) {
	command, next := e.GetCommand()
	if !next {
		return false, nil
	}
	if command[0].Tpe == MOVE.typ {
		X, err := strconv.Atoi(command[1].Text)
		if err != nil {
			return false, err
		}
		Y, err := strconv.Atoi(command[2].Text)
		if err != nil {
			return false, err
		}
		return next, e.bot.MoveMouseFluent(X, Y)
	}
	if command[0].Tpe == CLICK.typ {
		n, err := strconv.Atoi(command[1].Text)
		if err != nil {
			return false, err
		}
		return next, e.bot.ClickNTimes(n)
	}
	return false, nil
}

func (e *Engine) GetCommand() (t.Command, bool) {
	if e.contextStack.IsEmpty() {
		return nil, false
	}
	context, contextFinished, _ := e.contextStack.Pop()

	command := e.blocks[context.BlockName][context.Index]

	if command[0].Tpe == TEXT.typ {
		blockName := command[0].Text
		blockLength := len(e.blocks[blockName])
		e.contextStack.Push(ContextUnit{blockName, -1, blockLength})
		return e.GetCommand()
	}
	if command[0].Tpe == REPEAT.typ {
		blockName := command[1].Text
		blockLength := len(e.blocks[blockName])
		loopSize, _ := strconv.Atoi(command[2].Text)
		if e.loopCount < loopSize-1 {
			context.Index -= 1
		}
		if e.loopCount < loopSize {
			e.contextStack.Push(ContextUnit{blockName, -1, blockLength})
			loopCommand, loopContextFinished := e.GetCommand()
			if loopContextFinished {
				e.loopCount += 1
			}
			return loopCommand, loopContextFinished
		}
		e.loopCount = 0
		return e.GetCommand()
	}
	if context.BlockName == MAIN_BLOCK {
		return command, !contextFinished
	}
	return command, true
}

func (e *Engine) Init(blocks map[string][]t.Command, bot bot.Bot) {
	e.blocks = blocks
	e.loopCount = 0
	e.bot = bot
	initCu := ContextUnit{MAIN_BLOCK, -1, len(e.blocks[MAIN_BLOCK])}
	e.contextStack.Init(initCu)
}

func (e *Engine) Reset() {
	e.loopCount = 0
	initCu := ContextUnit{MAIN_BLOCK, -1, len(e.blocks[MAIN_BLOCK])}
	e.contextStack.Init(initCu)
}
