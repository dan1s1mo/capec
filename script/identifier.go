package script

import (
	"regexp"
)

type Identifier struct {
	typ string
	re  *regexp.Regexp
}

var EMPTY Identifier = Identifier{"EMPTY", nil}

var CommonIdntifiers map[string]Identifier = map[string]Identifier{
	"SCHAR":      Identifier{"SCHAR", regexp.MustCompile(`^[(),;:\{\}]`)},
	"DIGIT":      Identifier{"DIGIT", regexp.MustCompile(`^\d`)},
	"TEXT":       Identifier{"TEXT", regexp.MustCompile(`^[A-Za-z]`)},
	"TABULATION": Identifier{"TEXT", regexp.MustCompile("^[\r\t ]")},
}

var (
	NUMBER       Identifier = Identifier{"NUMBER", regexp.MustCompile(`^\d+`)}
	CLICK                   = Identifier{"CLICK", regexp.MustCompile(`^click`)}
	TEXT                    = Identifier{"TEXT", regexp.MustCompile(`^\w*`)}
	MOVE                    = Identifier{"MOVE", regexp.MustCompile(`^move`)}
	REPEAT                  = Identifier{"REPEAT", regexp.MustCompile(`^repeat`)}
	COLUMN                  = Identifier{"COLUMN", regexp.MustCompile("^:")}
	BLOCK                   = Identifier{"BLOCK", regexp.MustCompile(`^block`)}
	OPARENTHESIS            = Identifier{"OPENPARENTHESIS", regexp.MustCompile(`^\(`)}
	CPARENTHESIS            = Identifier{"CLOSEPARENTHESIS", regexp.MustCompile(`^\)`)}
	ENDOFLINE               = Identifier{"ENDOFLINE", regexp.MustCompile(`^;`)}
	COMA                    = Identifier{"COMA", regexp.MustCompile(`^,`)}
	LBRACKET                = Identifier{"LEFTBRACKET", regexp.MustCompile(`^\{`)}
	RBRACKET                = Identifier{"RIGHTBRACKET", regexp.MustCompile(`^\}`)}
	NEWLINE                 = Identifier{"NEWLINE", regexp.MustCompile("^\r\n")}
	TAB                     = Identifier{"TAB", regexp.MustCompile("^\t")}
	SPACE                   = Identifier{"SPACE", regexp.MustCompile("^ ")}
)
