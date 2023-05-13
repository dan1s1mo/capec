package script

import (
	t "capec/types"
)

func getIdentifier(input string) string {
	for _, id := range CommonIdntifiers {
		if id.re.MatchString(input) {
			return id.typ
		}
	}
	return EMPTY.typ
}

func getTokenText(text string, id Identifier) (string, string) {
	match := id.re.FindString(text)
	if len(match) == 0 {
		return match, text
	}
	rest := id.re.ReplaceAllString(text, "")
	return match, rest
}

type Lexer struct {
	index  int
	text   string
	tokens [](t.Token)
}

func (l *Lexer) addTokenDefault(id Identifier) bool {
	if tokenText, newText := getTokenText(l.text, id); len(tokenText) != 0 {
		l.tokens = append(l.tokens, t.Token{Position: l.index, Tpe: id.typ, Text: tokenText})
		l.index += len(tokenText)
		l.text = newText
		return true
	}
	return false
}

func (l *Lexer) GetTokens(input string) []t.Token {
	l.text = input
	l.index = 0
	for {
		if len(l.text) == 0 {
			break
		}
		currentIdentifier := getIdentifier(l.text)
		if currentIdentifier == CommonIdntifiers["TEXT"].typ {
			if l.addTokenDefault(CLICK) {
				continue
			}
			if l.addTokenDefault(MOVE) {
				continue
			}
			if l.addTokenDefault(REPEAT) {
				continue
			}
			if l.addTokenDefault(BLOCK) {
				continue
			}
			if l.addTokenDefault(TEXT) {
				continue
			}
		}
		if currentIdentifier == CommonIdntifiers["DIGIT"].typ {
			if l.addTokenDefault(NUMBER) {
				continue
			}
		}
		if currentIdentifier == CommonIdntifiers["SCHAR"].typ {
			if l.addTokenDefault(OPARENTHESIS) {
				continue
			}
			if l.addTokenDefault(CPARENTHESIS) {
				continue
			}
			if l.addTokenDefault(COMA) {
				continue
			}
			if l.addTokenDefault(ENDOFLINE) {
				continue
			}
			if l.addTokenDefault(COLUMN) {
				continue
			}
			if l.addTokenDefault(LBRACKET) {
				continue
			}
			if l.addTokenDefault(RBRACKET) {
				continue
			}
		}
		if currentIdentifier == CommonIdntifiers["TABULATION"].typ {
			if l.addTokenDefault(NEWLINE) {
				continue
			}
			if l.addTokenDefault(SPACE) {
				continue
			}

			if l.addTokenDefault(TAB) {
				continue
			}
		}

	}
	return l.tokens
}
