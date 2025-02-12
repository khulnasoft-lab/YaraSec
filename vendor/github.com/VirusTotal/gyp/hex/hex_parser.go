// Code generated by goyacc -p hex -o hex/hex_parser.go hex/hex_grammar.y. DO NOT EDIT.

//line hex/hex_grammar.y:31
package hex

import __yyfmt__ "fmt"

//line hex/hex_grammar.y:31

import (
	"github.com/VirusTotal/gyp/ast"
	gyperror "github.com/VirusTotal/gyp/error"
)

const StringChainingThreshold int = 200

type byteWithMask struct {
	Value byte
	Mask  byte
	Not   bool
}

//line hex/hex_grammar.y:70
type hexSymType struct {
	yys     int
	integer int
	bm      byteWithMask
	token   ast.HexToken
	tokens  ast.HexTokens
	bytes   *ast.HexBytes
	hexor   *ast.HexOr
}

const _BYTE_ = 57346
const _NOT_BYTE_ = 57347
const _MASKED_BYTE_ = 57348
const _NUMBER_ = 57349
const _LBRACE_ = 57350
const _RBRACE_ = 57351
const _LBRACKET_ = 57352
const _RBRACKET_ = 57354
const _HYPHEN_ = 57355
const _LPARENS_ = 57356
const _RPARENS_ = 57357
const _PIPE_ = 57358

var hexToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"_BYTE_",
	"_NOT_BYTE_",
	"_MASKED_BYTE_",
	"_NUMBER_",
	"_LBRACE_",
	"_RBRACE_",
	"_LBRACKET_",
	"$token",
	"_RBRACKET_",
	"_HYPHEN_",
	"_LPARENS_",
	"_RPARENS_",
	"_PIPE_",
}

var hexStatenames = [...]string{}

const hexEofCode = 1
const hexErrCode = 2
const hexInitialStackSize = 16

//line hex/hex_grammar.y:282

//line yacctab:1
var hexExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 12,
	9, 3,
	15, 3,
	16, 3,
	-2, 7,
	-1, 19,
	9, 4,
	15, 4,
	16, 4,
	-2, 7,
}

const hexPrivate = 57344

const hexLast = 40

var hexAct = [...]int8{
	3, 8, 10, 9, 8, 10, 9, 16, 28, 29,
	4, 6, 21, 14, 6, 12, 25, 26, 22, 24,
	30, 33, 27, 11, 19, 31, 2, 20, 7, 18,
	32, 8, 10, 9, 17, 1, 5, 15, 23, 13,
}

var hexPact = [...]int16{
	18, -1000, 0, 14, -3, 27, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -3, -1000, -1000, 5, -1000, 0, -1000,
	-1000, 4, 10, -7, -1000, -1000, 13, -1000, -1000, 0,
	9, -1000, -1000, -1000,
}

var hexPgo = [...]int8{
	0, 0, 39, 13, 10, 38, 37, 36, 28, 35,
	29,
}

var hexR1 = [...]int8{
	0, 9, 1, 1, 1, 2, 2, 3, 3, 4,
	10, 4, 6, 6, 6, 6, 5, 5, 7, 7,
	8, 8, 8,
}

var hexR2 = [...]int8{
	0, 3, 1, 2, 3, 1, 2, 1, 1, 1,
	0, 4, 3, 5, 4, 3, 1, 3, 1, 2,
	1, 1, 1,
}

var hexChk = [...]int16{
	-1000, -9, 8, -1, -4, -7, 14, -8, 4, 6,
	5, 9, -4, -2, -3, -6, 10, -8, -10, -4,
	-3, 7, 13, -5, -1, 12, 13, 12, 15, 16,
	7, 12, -1, 12,
}

var hexDef = [...]int8{
	0, -2, 0, 0, 2, 9, 10, 18, 20, 21,
	22, 1, -2, 0, 5, 8, 0, 19, 0, -2,
	6, 0, 0, 0, 16, 12, 0, 15, 11, 0,
	0, 14, 17, 13,
}

var hexTok1 = [...]int8{
	1,
}

var hexTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16,
}

var hexTok3 = [...]int8{
	0,
}

var hexErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	hexDebug        = 0
	hexErrorVerbose = false
)

type hexLexer interface {
	Lex(lval *hexSymType) int
	Error(s string)
}

type hexParser interface {
	Parse(hexLexer) int
	Lookahead() int
}

type hexParserImpl struct {
	lval  hexSymType
	stack [hexInitialStackSize]hexSymType
	char  int
}

func (p *hexParserImpl) Lookahead() int {
	return p.char
}

func hexNewParser() hexParser {
	return &hexParserImpl{}
}

const hexFlag = -1000

func hexTokname(c int) string {
	if c >= 1 && c-1 < len(hexToknames) {
		if hexToknames[c-1] != "" {
			return hexToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func hexStatname(s int) string {
	if s >= 0 && s < len(hexStatenames) {
		if hexStatenames[s] != "" {
			return hexStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func hexErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !hexErrorVerbose {
		return "syntax error"
	}

	for _, e := range hexErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + hexTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(hexPact[state])
	for tok := TOKSTART; tok-1 < len(hexToknames); tok++ {
		if n := base + tok; n >= 0 && n < hexLast && int(hexChk[int(hexAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if hexDef[state] == -2 {
		i := 0
		for hexExca[i] != -1 || int(hexExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; hexExca[i] >= 0; i += 2 {
			tok := int(hexExca[i])
			if tok < TOKSTART || hexExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if hexExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += hexTokname(tok)
	}
	return res
}

func hexlex1(lex hexLexer, lval *hexSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(hexTok1[0])
		goto out
	}
	if char < len(hexTok1) {
		token = int(hexTok1[char])
		goto out
	}
	if char >= hexPrivate {
		if char < hexPrivate+len(hexTok2) {
			token = int(hexTok2[char-hexPrivate])
			goto out
		}
	}
	for i := 0; i < len(hexTok3); i += 2 {
		token = int(hexTok3[i+0])
		if token == char {
			token = int(hexTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(hexTok2[1]) /* unknown char */
	}
	if hexDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", hexTokname(token), uint(char))
	}
	return char, token
}

func hexParse(hexlex hexLexer) int {
	return hexNewParser().Parse(hexlex)
}

func (hexrcvr *hexParserImpl) Parse(hexlex hexLexer) int {
	var hexn int
	var hexVAL hexSymType
	var hexDollar []hexSymType
	_ = hexDollar // silence set and not used
	hexS := hexrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	hexstate := 0
	hexrcvr.char = -1
	hextoken := -1 // hexrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		hexstate = -1
		hexrcvr.char = -1
		hextoken = -1
	}()
	hexp := -1
	goto hexstack

ret0:
	return 0

ret1:
	return 1

hexstack:
	/* put a state and value onto the stack */
	if hexDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", hexTokname(hextoken), hexStatname(hexstate))
	}

	hexp++
	if hexp >= len(hexS) {
		nyys := make([]hexSymType, len(hexS)*2)
		copy(nyys, hexS)
		hexS = nyys
	}
	hexS[hexp] = hexVAL
	hexS[hexp].yys = hexstate

hexnewstate:
	hexn = int(hexPact[hexstate])
	if hexn <= hexFlag {
		goto hexdefault /* simple state */
	}
	if hexrcvr.char < 0 {
		hexrcvr.char, hextoken = hexlex1(hexlex, &hexrcvr.lval)
	}
	hexn += hextoken
	if hexn < 0 || hexn >= hexLast {
		goto hexdefault
	}
	hexn = int(hexAct[hexn])
	if int(hexChk[hexn]) == hextoken { /* valid shift */
		hexrcvr.char = -1
		hextoken = -1
		hexVAL = hexrcvr.lval
		hexstate = hexn
		if Errflag > 0 {
			Errflag--
		}
		goto hexstack
	}

hexdefault:
	/* default state action */
	hexn = int(hexDef[hexstate])
	if hexn == -2 {
		if hexrcvr.char < 0 {
			hexrcvr.char, hextoken = hexlex1(hexlex, &hexrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if hexExca[xi+0] == -1 && int(hexExca[xi+1]) == hexstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			hexn = int(hexExca[xi+0])
			if hexn < 0 || hexn == hextoken {
				break
			}
		}
		hexn = int(hexExca[xi+1])
		if hexn < 0 {
			goto ret0
		}
	}
	if hexn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			hexlex.Error(hexErrorMessage(hexstate, hextoken))
			Nerrs++
			if hexDebug >= 1 {
				__yyfmt__.Printf("%s", hexStatname(hexstate))
				__yyfmt__.Printf(" saw %s\n", hexTokname(hextoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for hexp >= 0 {
				hexn = int(hexPact[hexS[hexp].yys]) + hexErrCode
				if hexn >= 0 && hexn < hexLast {
					hexstate = int(hexAct[hexn]) /* simulate a shift of "error" */
					if int(hexChk[hexstate]) == hexErrCode {
						goto hexstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if hexDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", hexS[hexp].yys)
				}
				hexp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if hexDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", hexTokname(hextoken))
			}
			if hextoken == hexEofCode {
				goto ret1
			}
			hexrcvr.char = -1
			hextoken = -1
			goto hexnewstate /* try again in the same state */
		}
	}

	/* reduction by production hexn */
	if hexDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", hexn, hexStatname(hexstate))
	}

	hexnt := hexn
	hexpt := hexp
	_ = hexpt // guard against "declared and not used"

	hexp -= int(hexR2[hexn])
	// hexp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if hexp+1 >= len(hexS) {
		nyys := make([]hexSymType, len(hexS)*2)
		copy(nyys, hexS)
		hexS = nyys
	}
	hexVAL = hexS[hexp+1]

	/* consult goto table to find next state */
	hexn = int(hexR1[hexn])
	hexg := int(hexPgo[hexn])
	hexj := hexg + hexS[hexp].yys + 1

	if hexj >= hexLast {
		hexstate = int(hexAct[hexg])
	} else {
		hexstate = int(hexAct[hexj])
		if int(hexChk[hexstate]) != -hexn {
			hexstate = int(hexAct[hexg])
		}
	}
	// dummy call; replaced with literal code
	switch hexnt {

	case 1:
		hexDollar = hexS[hexpt-3 : hexpt+1]
//line hex/hex_grammar.y:83
		{
			asLexer(hexlex).hexTokens = hexDollar[2].tokens
		}
	case 2:
		hexDollar = hexS[hexpt-1 : hexpt+1]
//line hex/hex_grammar.y:91
		{
			hexVAL.tokens = []ast.HexToken{hexDollar[1].token}
		}
	case 3:
		hexDollar = hexS[hexpt-2 : hexpt+1]
//line hex/hex_grammar.y:95
		{
			hexVAL.tokens = []ast.HexToken{hexDollar[1].token, hexDollar[2].token}
		}
	case 4:
		hexDollar = hexS[hexpt-3 : hexpt+1]
//line hex/hex_grammar.y:99
		{
			hexVAL.tokens = append(append([]ast.HexToken{hexDollar[1].token}, hexDollar[2].tokens...), hexDollar[3].token)
		}
	case 5:
		hexDollar = hexS[hexpt-1 : hexpt+1]
//line hex/hex_grammar.y:107
		{
			hexVAL.tokens = []ast.HexToken{hexDollar[1].token}
		}
	case 6:
		hexDollar = hexS[hexpt-2 : hexpt+1]
//line hex/hex_grammar.y:111
		{
			hexVAL.tokens = append(hexDollar[1].tokens, hexDollar[2].token)
		}
	case 7:
		hexDollar = hexS[hexpt-1 : hexpt+1]
//line hex/hex_grammar.y:119
		{
			hexVAL.token = hexDollar[1].token
		}
	case 8:
		hexDollar = hexS[hexpt-1 : hexpt+1]
//line hex/hex_grammar.y:123
		{
			hexVAL.token = hexDollar[1].token
		}
	case 9:
		hexDollar = hexS[hexpt-1 : hexpt+1]
//line hex/hex_grammar.y:131
		{
			hexVAL.token = hexDollar[1].bytes
		}
	case 10:
		hexDollar = hexS[hexpt-1 : hexpt+1]
//line hex/hex_grammar.y:135
		{
			asLexer(hexlex).insideOr += 1
		}
	case 11:
		hexDollar = hexS[hexpt-4 : hexpt+1]
//line hex/hex_grammar.y:139
		{
			asLexer(hexlex).insideOr -= 1
			hexVAL.token = hexDollar[3].hexor
		}
	case 12:
		hexDollar = hexS[hexpt-3 : hexpt+1]
//line hex/hex_grammar.y:148
		{
			lexer := asLexer(hexlex)

			if hexDollar[2].integer <= 0 {
				return lexer.setError(
					gyperror.InvalidJumpLengthError,
					`invalid jump length: %d`, hexDollar[2].integer)
			}

			if lexer.insideOr > 0 && hexDollar[2].integer > StringChainingThreshold {
				return lexer.setError(
					gyperror.JumpTooLargeInsideAlternationError,
					`jump too large inside alternation: %d`, hexDollar[2].integer)
			}

			hexVAL.token = &ast.HexJump{
				Start: hexDollar[2].integer,
				End:   hexDollar[2].integer,
			}
		}
	case 13:
		hexDollar = hexS[hexpt-5 : hexpt+1]
//line hex/hex_grammar.y:169
		{
			lexer := asLexer(hexlex)

			if lexer.insideOr > 0 &&
				(hexDollar[2].integer > StringChainingThreshold || hexDollar[4].integer > StringChainingThreshold) {
				return lexer.setError(
					gyperror.JumpTooLargeInsideAlternationError,
					`jump too large inside alternation: %d-%d`, hexDollar[2].integer, hexDollar[4].integer)
			}

			if hexDollar[2].integer < 0 || hexDollar[4].integer < 0 {
				return lexer.setError(
					gyperror.NegativeJumpError,
					`negative jump: %d-%d`, hexDollar[2].integer, hexDollar[4].integer)
			}

			if hexDollar[2].integer > hexDollar[4].integer {
				return lexer.setError(
					gyperror.InvalidJumpRangeError,
					`jump too large inside alternation: %d-%d`, hexDollar[2].integer, hexDollar[4].integer)
			}

			hexVAL.token = &ast.HexJump{
				Start: hexDollar[2].integer,
				End:   hexDollar[4].integer,
			}
		}
	case 14:
		hexDollar = hexS[hexpt-4 : hexpt+1]
//line hex/hex_grammar.y:197
		{
			lexer := asLexer(hexlex)

			if lexer.insideOr > 0 {
				return lexer.setError(
					gyperror.UnboundedJumpInsideAlternationError,
					`unbounded jump inside alternation: %d`, hexDollar[2].integer)
			}

			if hexDollar[2].integer < 0 {
				return lexer.setError(
					gyperror.NegativeJumpError,
					`negative jump: %d`, hexDollar[2].integer)
			}

			hexVAL.token = &ast.HexJump{
				Start: hexDollar[2].integer,
			}
		}
	case 15:
		hexDollar = hexS[hexpt-3 : hexpt+1]
//line hex/hex_grammar.y:217
		{
			lexer := asLexer(hexlex)

			if lexer.insideOr > 0 {
				return lexer.setError(
					gyperror.UnboundedJumpInsideAlternationError,
					`unbounded jump inside alternation`)
			}

			hexVAL.token = &ast.HexJump{}
		}
	case 16:
		hexDollar = hexS[hexpt-1 : hexpt+1]
//line hex/hex_grammar.y:233
		{
			hexVAL.hexor = &ast.HexOr{
				Alternatives: ast.HexTokens{hexDollar[1].tokens},
			}
		}
	case 17:
		hexDollar = hexS[hexpt-3 : hexpt+1]
//line hex/hex_grammar.y:239
		{
			hexDollar[1].hexor.Alternatives = append(hexDollar[1].hexor.Alternatives, hexDollar[3].tokens)
			hexVAL.hexor = hexDollar[1].hexor
		}
	case 18:
		hexDollar = hexS[hexpt-1 : hexpt+1]
//line hex/hex_grammar.y:252
		{
			hexVAL.bytes = &ast.HexBytes{
				Bytes: []byte{hexDollar[1].bm.Value},
				Masks: []byte{hexDollar[1].bm.Mask},
				Nots:  []bool{hexDollar[1].bm.Not},
			}
		}
	case 19:
		hexDollar = hexS[hexpt-2 : hexpt+1]
//line hex/hex_grammar.y:260
		{
			hexDollar[1].bytes.Bytes = append(hexDollar[1].bytes.Bytes, hexDollar[2].bm.Value)
			hexDollar[1].bytes.Masks = append(hexDollar[1].bytes.Masks, hexDollar[2].bm.Mask)
			hexDollar[1].bytes.Nots = append(hexDollar[1].bytes.Nots, hexDollar[2].bm.Not)
		}
	case 20:
		hexDollar = hexS[hexpt-1 : hexpt+1]
//line hex/hex_grammar.y:269
		{
			hexVAL.bm = hexDollar[1].bm
		}
	case 21:
		hexDollar = hexS[hexpt-1 : hexpt+1]
//line hex/hex_grammar.y:273
		{
			hexVAL.bm = hexDollar[1].bm
		}
	case 22:
		hexDollar = hexS[hexpt-1 : hexpt+1]
//line hex/hex_grammar.y:277
		{
			hexVAL.bm = hexDollar[1].bm
		}
	}
	goto hexstack /* stack new state and value */
}
