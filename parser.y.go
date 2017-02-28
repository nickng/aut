//line aut.y:2
package aut

import __yyfmt__ "fmt"

//line aut.y:2
import (
	"io"
)

var autfile = new(Aut)

//line aut.y:11
type autSymType struct {
	yys   int
	str   string
	num   int
	trans *Transition
}

const DES = 57346
const LPAREN = 57347
const RPAREN = 57348
const COMMA = 57349
const DQUOTE = 57350
const DIGITS = 57351
const LABEL = 57352

var autToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"DES",
	"LPAREN",
	"RPAREN",
	"COMMA",
	"DQUOTE",
	"DIGITS",
	"LABEL",
}
var autStatenames = [...]string{}

const autEofCode = 1
const autErrCode = 2
const autInitialStackSize = 16

//line aut.y:34

func Parse(r io.Reader) (*Aut, error) {
	l := NewLexer(r)
	autParse(l)
	select {
	case err := <-l.Errors:
		return nil, err
	default:
		return autfile, nil
	}
}

//line yacctab:1
var autExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const autNprod = 5
const autPrivate = 57344

var autTokenNames []string
var autStates []string

const autLast = 25

var autAct = [...]int{

	16, 14, 22, 21, 20, 15, 10, 9, 8, 19,
	18, 17, 13, 12, 11, 25, 24, 23, 7, 6,
	5, 3, 2, 1, 4,
}
var autPact = [...]int{

	17, -1000, 15, 14, 13, -1, -2, -3, 7, 6,
	5, -9, -4, -10, 4, 3, 2, -5, -6, -7,
	11, 10, 9, -1000, -1000, -1000,
}
var autPgo = [...]int{

	0, 24, 23, 22,
}
var autR1 = [...]int{

	0, 2, 3, 1, 1,
}
var autR2 = [...]int{

	0, 2, 8, 8, 7,
}
var autChk = [...]int{

	-1000, -2, -3, 4, -1, 5, 5, 5, 9, 9,
	9, 7, 7, 7, 10, 9, 10, 7, 7, 7,
	9, 9, 9, 6, 6, 6,
}
var autDef = [...]int{

	0, -2, 0, 0, 1, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 4, 2, 3,
}
var autTok1 = [...]int{

	1,
}
var autTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10,
}
var autTok3 = [...]int{
	0,
}

var autErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	autDebug        = 0
	autErrorVerbose = false
)

type autLexer interface {
	Lex(lval *autSymType) int
	Error(s string)
}

type autParser interface {
	Parse(autLexer) int
	Lookahead() int
}

type autParserImpl struct {
	lval  autSymType
	stack [autInitialStackSize]autSymType
	char  int
}

func (p *autParserImpl) Lookahead() int {
	return p.char
}

func autNewParser() autParser {
	return &autParserImpl{}
}

const autFlag = -1000

func autTokname(c int) string {
	if c >= 1 && c-1 < len(autToknames) {
		if autToknames[c-1] != "" {
			return autToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func autStatname(s int) string {
	if s >= 0 && s < len(autStatenames) {
		if autStatenames[s] != "" {
			return autStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func autErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !autErrorVerbose {
		return "syntax error"
	}

	for _, e := range autErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + autTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := autPact[state]
	for tok := TOKSTART; tok-1 < len(autToknames); tok++ {
		if n := base + tok; n >= 0 && n < autLast && autChk[autAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if autDef[state] == -2 {
		i := 0
		for autExca[i] != -1 || autExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; autExca[i] >= 0; i += 2 {
			tok := autExca[i]
			if tok < TOKSTART || autExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if autExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += autTokname(tok)
	}
	return res
}

func autlex1(lex autLexer, lval *autSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = autTok1[0]
		goto out
	}
	if char < len(autTok1) {
		token = autTok1[char]
		goto out
	}
	if char >= autPrivate {
		if char < autPrivate+len(autTok2) {
			token = autTok2[char-autPrivate]
			goto out
		}
	}
	for i := 0; i < len(autTok3); i += 2 {
		token = autTok3[i+0]
		if token == char {
			token = autTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = autTok2[1] /* unknown char */
	}
	if autDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", autTokname(token), uint(char))
	}
	return char, token
}

func autParse(autlex autLexer) int {
	return autNewParser().Parse(autlex)
}

func (autrcvr *autParserImpl) Parse(autlex autLexer) int {
	var autn int
	var autVAL autSymType
	var autDollar []autSymType
	_ = autDollar // silence set and not used
	autS := autrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	autstate := 0
	autrcvr.char = -1
	auttoken := -1 // autrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		autstate = -1
		autrcvr.char = -1
		auttoken = -1
	}()
	autp := -1
	goto autstack

ret0:
	return 0

ret1:
	return 1

autstack:
	/* put a state and value onto the stack */
	if autDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", autTokname(auttoken), autStatname(autstate))
	}

	autp++
	if autp >= len(autS) {
		nyys := make([]autSymType, len(autS)*2)
		copy(nyys, autS)
		autS = nyys
	}
	autS[autp] = autVAL
	autS[autp].yys = autstate

autnewstate:
	autn = autPact[autstate]
	if autn <= autFlag {
		goto autdefault /* simple state */
	}
	if autrcvr.char < 0 {
		autrcvr.char, auttoken = autlex1(autlex, &autrcvr.lval)
	}
	autn += auttoken
	if autn < 0 || autn >= autLast {
		goto autdefault
	}
	autn = autAct[autn]
	if autChk[autn] == auttoken { /* valid shift */
		autrcvr.char = -1
		auttoken = -1
		autVAL = autrcvr.lval
		autstate = autn
		if Errflag > 0 {
			Errflag--
		}
		goto autstack
	}

autdefault:
	/* default state action */
	autn = autDef[autstate]
	if autn == -2 {
		if autrcvr.char < 0 {
			autrcvr.char, auttoken = autlex1(autlex, &autrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if autExca[xi+0] == -1 && autExca[xi+1] == autstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			autn = autExca[xi+0]
			if autn < 0 || autn == auttoken {
				break
			}
		}
		autn = autExca[xi+1]
		if autn < 0 {
			goto ret0
		}
	}
	if autn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			autlex.Error(autErrorMessage(autstate, auttoken))
			Nerrs++
			if autDebug >= 1 {
				__yyfmt__.Printf("%s", autStatname(autstate))
				__yyfmt__.Printf(" saw %s\n", autTokname(auttoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for autp >= 0 {
				autn = autPact[autS[autp].yys] + autErrCode
				if autn >= 0 && autn < autLast {
					autstate = autAct[autn] /* simulate a shift of "error" */
					if autChk[autstate] == autErrCode {
						goto autstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if autDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", autS[autp].yys)
				}
				autp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if autDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", autTokname(auttoken))
			}
			if auttoken == autEofCode {
				goto ret1
			}
			autrcvr.char = -1
			auttoken = -1
			goto autnewstate /* try again in the same state */
		}
	}

	/* reduction by production autn */
	if autDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", autn, autStatname(autstate))
	}

	autnt := autn
	autpt := autp
	_ = autpt // guard against "declared and not used"

	autp -= autR2[autn]
	// autp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if autp+1 >= len(autS) {
		nyys := make([]autSymType, len(autS)*2)
		copy(nyys, autS)
		autS = nyys
	}
	autVAL = autS[autp+1]

	/* consult goto table to find next state */
	autn = autR1[autn]
	autg := autPgo[autn]
	autj := autg + autS[autp].yys + 1

	if autj >= autLast {
		autstate = autAct[autg]
	} else {
		autstate = autAct[autj]
		if autChk[autstate] != -autn {
			autstate = autAct[autg]
		}
	}
	// dummy call; replaced with literal code
	switch autnt {

	case 2:
		autDollar = autS[autpt-8 : autpt+1]
		//line aut.y:27
		{
			autfile.SetDes(autDollar[3].num, autDollar[5].num, autDollar[7].num)
		}
	case 3:
		autDollar = autS[autpt-8 : autpt+1]
		//line aut.y:30
		{
			autfile.AddTransition(State(autDollar[3].num), autDollar[5].str, State(autDollar[7].num))
		}
	case 4:
		autDollar = autS[autpt-7 : autpt+1]
		//line aut.y:31
		{
			autfile.AddTransition(State(autDollar[2].num), autDollar[4].str, State(autDollar[6].num))
		}
	}
	goto autstack /* stack new state and value */
}
