package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var g = &grammar{
	rules: []*rule{
		{
			name: "Start",
			pos:  position{line: 5, col: 1, offset: 20},
			expr: &seqExpr{
				pos: position{line: 6, col: 6, offset: 32},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 6, col: 6, offset: 32},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 7, col: 8, offset: 43},
						alternatives: []interface{}{
							&actionExpr{
								pos: position{line: 7, col: 8, offset: 43},
								run: (*parser).callonStart4,
								expr: &seqExpr{
									pos: position{line: 7, col: 8, offset: 43},
									exprs: []interface{}{
										&labeledExpr{
											pos:   position{line: 7, col: 8, offset: 43},
											label: "ins",
											expr: &ruleRefExpr{
												pos:  position{line: 7, col: 12, offset: 47},
												name: "Instruction",
											},
										},
										&zeroOrOneExpr{
											pos: position{line: 7, col: 24, offset: 59},
											expr: &ruleRefExpr{
												pos:  position{line: 7, col: 24, offset: 59},
												name: "Comment",
											},
										},
									},
								},
							},
							&actionExpr{
								pos: position{line: 8, col: 8, offset: 104},
								run: (*parser).callonStart10,
								expr: &labeledExpr{
									pos:   position{line: 8, col: 8, offset: 104},
									label: "lbl",
									expr: &ruleRefExpr{
										pos:  position{line: 8, col: 12, offset: 108},
										name: "Label",
									},
								},
							},
							&actionExpr{
								pos: position{line: 9, col: 8, offset: 165},
								run: (*parser).callonStart13,
								expr: &labeledExpr{
									pos:   position{line: 9, col: 8, offset: 165},
									label: "comment",
									expr: &ruleRefExpr{
										pos:  position{line: 9, col: 16, offset: 173},
										name: "Comment",
									},
								},
							},
						},
					},
					&notExpr{
						pos: position{line: 10, col: 6, offset: 228},
						expr: &anyMatcher{
							line: 10, col: 7, offset: 229,
						},
					},
				},
			},
		},
		{
			name: "Instruction",
			pos:  position{line: 12, col: 1, offset: 232},
			expr: &choiceExpr{
				pos: position{line: 13, col: 6, offset: 250},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 13, col: 6, offset: 250},
						run: (*parser).callonInstruction2,
						expr: &seqExpr{
							pos: position{line: 13, col: 6, offset: 250},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 13, col: 6, offset: 250},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 13, col: 8, offset: 252},
									label: "set",
									expr: &zeroOrOneExpr{
										pos: position{line: 13, col: 12, offset: 256},
										expr: &litMatcher{
											pos:        position{line: 13, col: 12, offset: 256},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 13, col: 17, offset: 261},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 13, col: 19, offset: 263},
									val:        "add",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 13, col: 25, offset: 269},
									label: "cnd",
									expr: &zeroOrOneExpr{
										pos: position{line: 13, col: 29, offset: 273},
										expr: &ruleRefExpr{
											pos:  position{line: 13, col: 29, offset: 273},
											name: "Condition",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 13, col: 40, offset: 284},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 13, col: 42, offset: 286},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 13, col: 45, offset: 289},
										name: "Register",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 13, col: 54, offset: 298},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 13, col: 56, offset: 300},
									label: "ra",
									expr: &ruleRefExpr{
										pos:  position{line: 13, col: 59, offset: 303},
										name: "Register",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 13, col: 68, offset: 312},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 13, col: 70, offset: 314},
									label: "rb",
									expr: &ruleRefExpr{
										pos:  position{line: 13, col: 73, offset: 317},
										name: "Register",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 13, col: 82, offset: 326},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 14, col: 6, offset: 386},
						run: (*parser).callonInstruction23,
						expr: &seqExpr{
							pos: position{line: 14, col: 6, offset: 386},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 14, col: 6, offset: 386},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 14, col: 8, offset: 388},
									label: "set",
									expr: &zeroOrOneExpr{
										pos: position{line: 14, col: 12, offset: 392},
										expr: &litMatcher{
											pos:        position{line: 14, col: 12, offset: 392},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 14, col: 17, offset: 397},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 14, col: 19, offset: 399},
									val:        "add",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 14, col: 25, offset: 405},
									label: "cnd",
									expr: &zeroOrOneExpr{
										pos: position{line: 14, col: 29, offset: 409},
										expr: &ruleRefExpr{
											pos:  position{line: 14, col: 29, offset: 409},
											name: "Condition",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 14, col: 40, offset: 420},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 14, col: 42, offset: 422},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 14, col: 45, offset: 425},
										name: "Register",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 14, col: 54, offset: 434},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 14, col: 56, offset: 436},
									label: "ra",
									expr: &ruleRefExpr{
										pos:  position{line: 14, col: 59, offset: 439},
										name: "Register",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 14, col: 68, offset: 448},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 14, col: 70, offset: 450},
									label: "num",
									expr: &ruleRefExpr{
										pos:  position{line: 14, col: 74, offset: 454},
										name: "Number",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 14, col: 81, offset: 461},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 15, col: 6, offset: 525},
						run: (*parser).callonInstruction44,
						expr: &seqExpr{
							pos: position{line: 15, col: 6, offset: 525},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 15, col: 6, offset: 525},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 15, col: 8, offset: 527},
									label: "set",
									expr: &zeroOrOneExpr{
										pos: position{line: 15, col: 12, offset: 531},
										expr: &litMatcher{
											pos:        position{line: 15, col: 12, offset: 531},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 15, col: 17, offset: 536},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 15, col: 19, offset: 538},
									val:        "add",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 15, col: 25, offset: 544},
									label: "cnd",
									expr: &zeroOrOneExpr{
										pos: position{line: 15, col: 29, offset: 548},
										expr: &ruleRefExpr{
											pos:  position{line: 15, col: 29, offset: 548},
											name: "Condition",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 15, col: 40, offset: 559},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 15, col: 42, offset: 561},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 15, col: 45, offset: 564},
										name: "Register",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 15, col: 54, offset: 573},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 15, col: 56, offset: 575},
									label: "rb",
									expr: &ruleRefExpr{
										pos:  position{line: 15, col: 59, offset: 578},
										name: "Register",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 15, col: 68, offset: 587},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 16, col: 6, offset: 661},
						run: (*parser).callonInstruction62,
						expr: &seqExpr{
							pos: position{line: 16, col: 6, offset: 661},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 16, col: 6, offset: 661},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 16, col: 8, offset: 663},
									label: "set",
									expr: &zeroOrOneExpr{
										pos: position{line: 16, col: 12, offset: 667},
										expr: &litMatcher{
											pos:        position{line: 16, col: 12, offset: 667},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 16, col: 17, offset: 672},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 16, col: 19, offset: 674},
									val:        "add",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 16, col: 25, offset: 680},
									label: "cnd",
									expr: &zeroOrOneExpr{
										pos: position{line: 16, col: 29, offset: 684},
										expr: &ruleRefExpr{
											pos:  position{line: 16, col: 29, offset: 684},
											name: "Condition",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 16, col: 40, offset: 695},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 16, col: 42, offset: 697},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 16, col: 45, offset: 700},
										name: "Register",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 16, col: 54, offset: 709},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 16, col: 56, offset: 711},
									label: "num",
									expr: &ruleRefExpr{
										pos:  position{line: 16, col: 60, offset: 715},
										name: "Number",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 16, col: 67, offset: 722},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 18, col: 6, offset: 799},
						run: (*parser).callonInstruction80,
						expr: &seqExpr{
							pos: position{line: 18, col: 6, offset: 799},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 18, col: 6, offset: 799},
									label: "cmd",
									expr: &ruleRefExpr{
										pos:  position{line: 18, col: 10, offset: 803},
										name: "Command",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 18, col: 18, offset: 811},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 19, col: 6, offset: 818},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 19, col: 9, offset: 821},
										name: "Register",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 19, col: 18, offset: 830},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 20, col: 6, offset: 838},
									label: "ra",
									expr: &ruleRefExpr{
										pos:  position{line: 20, col: 9, offset: 841},
										name: "Register",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 20, col: 18, offset: 850},
									name: "_",
								},
								&zeroOrOneExpr{
									pos: position{line: 21, col: 6, offset: 858},
									expr: &seqExpr{
										pos: position{line: 21, col: 7, offset: 859},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 21, col: 7, offset: 859},
												val:        "[",
												ignoreCase: false,
											},
											&ruleRefExpr{
												pos:  position{line: 21, col: 11, offset: 863},
												name: "_",
											},
										},
									},
								},
								&labeledExpr{
									pos:   position{line: 22, col: 6, offset: 873},
									label: "rb",
									expr: &ruleRefExpr{
										pos:  position{line: 22, col: 9, offset: 876},
										name: "Register",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 22, col: 18, offset: 885},
									name: "_",
								},
								&zeroOrOneExpr{
									pos: position{line: 23, col: 6, offset: 893},
									expr: &seqExpr{
										pos: position{line: 23, col: 7, offset: 894},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 23, col: 7, offset: 894},
												val:        "]",
												ignoreCase: false,
											},
											&ruleRefExpr{
												pos:  position{line: 23, col: 11, offset: 898},
												name: "_",
											},
										},
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 25, col: 6, offset: 1028},
						run: (*parser).callonInstruction102,
						expr: &seqExpr{
							pos: position{line: 25, col: 6, offset: 1028},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 25, col: 6, offset: 1028},
									label: "cmd",
									expr: &ruleRefExpr{
										pos:  position{line: 25, col: 10, offset: 1032},
										name: "Command",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 25, col: 18, offset: 1040},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 26, col: 6, offset: 1047},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 26, col: 9, offset: 1050},
										name: "Register",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 26, col: 18, offset: 1059},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 27, col: 6, offset: 1067},
									label: "ra",
									expr: &ruleRefExpr{
										pos:  position{line: 27, col: 9, offset: 1070},
										name: "Register",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 27, col: 18, offset: 1079},
									name: "_",
								},
								&zeroOrOneExpr{
									pos: position{line: 28, col: 6, offset: 1087},
									expr: &seqExpr{
										pos: position{line: 28, col: 7, offset: 1088},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 28, col: 7, offset: 1088},
												val:        "[",
												ignoreCase: false,
											},
											&ruleRefExpr{
												pos:  position{line: 28, col: 11, offset: 1092},
												name: "_",
											},
										},
									},
								},
								&labeledExpr{
									pos:   position{line: 29, col: 6, offset: 1102},
									label: "num",
									expr: &ruleRefExpr{
										pos:  position{line: 29, col: 10, offset: 1106},
										name: "Number",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 29, col: 17, offset: 1113},
									name: "_",
								},
								&zeroOrOneExpr{
									pos: position{line: 30, col: 6, offset: 1121},
									expr: &seqExpr{
										pos: position{line: 30, col: 7, offset: 1122},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 30, col: 7, offset: 1122},
												val:        "]",
												ignoreCase: false,
											},
											&ruleRefExpr{
												pos:  position{line: 30, col: 11, offset: 1126},
												name: "_",
											},
										},
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 32, col: 6, offset: 1252},
						run: (*parser).callonInstruction124,
						expr: &seqExpr{
							pos: position{line: 32, col: 6, offset: 1252},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 32, col: 6, offset: 1252},
									label: "cmd",
									expr: &ruleRefExpr{
										pos:  position{line: 32, col: 10, offset: 1256},
										name: "Command",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 32, col: 18, offset: 1264},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 33, col: 6, offset: 1272},
									label: "lbl",
									expr: &ruleRefExpr{
										pos:  position{line: 33, col: 10, offset: 1276},
										name: "Label",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 33, col: 16, offset: 1282},
									name: "_",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Command",
			pos:  position{line: 35, col: 1, offset: 1365},
			expr: &actionExpr{
				pos: position{line: 36, col: 6, offset: 1379},
				run: (*parser).callonCommand1,
				expr: &seqExpr{
					pos: position{line: 36, col: 6, offset: 1379},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 36, col: 6, offset: 1379},
							label: "set",
							expr: &zeroOrOneExpr{
								pos: position{line: 36, col: 10, offset: 1383},
								expr: &litMatcher{
									pos:        position{line: 36, col: 10, offset: 1383},
									val:        "!",
									ignoreCase: false,
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 36, col: 15, offset: 1388},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 37, col: 6, offset: 1395},
							label: "cmd",
							expr: &ruleRefExpr{
								pos:  position{line: 37, col: 10, offset: 1399},
								name: "Mnemonic",
							},
						},
						&labeledExpr{
							pos:   position{line: 38, col: 6, offset: 1414},
							label: "cnd",
							expr: &zeroOrOneExpr{
								pos: position{line: 38, col: 10, offset: 1418},
								expr: &ruleRefExpr{
									pos:  position{line: 38, col: 10, offset: 1418},
									name: "Condition",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Mnemonic",
			pos:  position{line: 40, col: 1, offset: 1485},
			expr: &actionExpr{
				pos: position{line: 40, col: 13, offset: 1497},
				run: (*parser).callonMnemonic1,
				expr: &seqExpr{
					pos: position{line: 40, col: 13, offset: 1497},
					exprs: []interface{}{
						&charClassMatcher{
							pos:        position{line: 40, col: 13, offset: 1497},
							val:        "[a-z]",
							ranges:     []rune{'a', 'z'},
							ignoreCase: false,
							inverted:   false,
						},
						&charClassMatcher{
							pos:        position{line: 40, col: 18, offset: 1502},
							val:        "[a-z]",
							ranges:     []rune{'a', 'z'},
							ignoreCase: false,
							inverted:   false,
						},
						&charClassMatcher{
							pos:        position{line: 40, col: 23, offset: 1507},
							val:        "[a-z]",
							ranges:     []rune{'a', 'z'},
							ignoreCase: false,
							inverted:   false,
						},
					},
				},
			},
		},
		{
			name: "Condition",
			pos:  position{line: 42, col: 1, offset: 1553},
			expr: &actionExpr{
				pos: position{line: 42, col: 14, offset: 1566},
				run: (*parser).callonCondition1,
				expr: &seqExpr{
					pos: position{line: 42, col: 14, offset: 1566},
					exprs: []interface{}{
						&charClassMatcher{
							pos:        position{line: 42, col: 14, offset: 1566},
							val:        "[a-z]",
							ranges:     []rune{'a', 'z'},
							ignoreCase: false,
							inverted:   false,
						},
						&charClassMatcher{
							pos:        position{line: 42, col: 19, offset: 1571},
							val:        "[a-z]",
							ranges:     []rune{'a', 'z'},
							ignoreCase: false,
							inverted:   false,
						},
					},
				},
			},
		},
		{
			name: "Register",
			pos:  position{line: 44, col: 1, offset: 1621},
			expr: &actionExpr{
				pos: position{line: 44, col: 13, offset: 1633},
				run: (*parser).callonRegister1,
				expr: &choiceExpr{
					pos: position{line: 44, col: 14, offset: 1634},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 44, col: 14, offset: 1634},
							val:        "sp",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 44, col: 21, offset: 1641},
							val:        "rp",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 44, col: 28, offset: 1648},
							val:        "ip",
							ignoreCase: false,
						},
						&seqExpr{
							pos: position{line: 44, col: 35, offset: 1655},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 44, col: 35, offset: 1655},
									val:        "r",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 44, col: 39, offset: 1659},
									expr: &charClassMatcher{
										pos:        position{line: 44, col: 39, offset: 1659},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Number",
			pos:  position{line: 46, col: 1, offset: 1701},
			expr: &actionExpr{
				pos: position{line: 46, col: 11, offset: 1711},
				run: (*parser).callonNumber1,
				expr: &labeledExpr{
					pos:   position{line: 46, col: 11, offset: 1711},
					label: "num",
					expr: &choiceExpr{
						pos: position{line: 46, col: 16, offset: 1716},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 46, col: 16, offset: 1716},
								name: "HexNumber",
							},
							&ruleRefExpr{
								pos:  position{line: 46, col: 28, offset: 1728},
								name: "DecNumber",
							},
						},
					},
				},
			},
		},
		{
			name: "DecNumber",
			pos:  position{line: 48, col: 1, offset: 1765},
			expr: &actionExpr{
				pos: position{line: 48, col: 14, offset: 1778},
				run: (*parser).callonDecNumber1,
				expr: &seqExpr{
					pos: position{line: 48, col: 14, offset: 1778},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 48, col: 14, offset: 1778},
							expr: &litMatcher{
								pos:        position{line: 48, col: 14, offset: 1778},
								val:        "-",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 48, col: 19, offset: 1783},
							expr: &charClassMatcher{
								pos:        position{line: 48, col: 19, offset: 1783},
								val:        "[0-9]",
								ranges:     []rune{'0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
						},
					},
				},
			},
		},
		{
			name: "HexNumber",
			pos:  position{line: 50, col: 1, offset: 1837},
			expr: &actionExpr{
				pos: position{line: 50, col: 14, offset: 1850},
				run: (*parser).callonHexNumber1,
				expr: &seqExpr{
					pos: position{line: 50, col: 14, offset: 1850},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 50, col: 14, offset: 1850},
							val:        "0x",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 50, col: 19, offset: 1855},
							expr: &charClassMatcher{
								pos:        position{line: 50, col: 19, offset: 1855},
								val:        "[0-9a-f]",
								ranges:     []rune{'0', '9', 'a', 'f'},
								ignoreCase: false,
								inverted:   false,
							},
						},
					},
				},
			},
		},
		{
			name: "Label",
			pos:  position{line: 52, col: 1, offset: 1909},
			expr: &actionExpr{
				pos: position{line: 52, col: 10, offset: 1918},
				run: (*parser).callonLabel1,
				expr: &seqExpr{
					pos: position{line: 52, col: 10, offset: 1918},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 52, col: 10, offset: 1918},
							val:        "@",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 52, col: 14, offset: 1922},
							expr: &charClassMatcher{
								pos:        position{line: 52, col: 14, offset: 1922},
								val:        "[a-zA-Z0-9]",
								ranges:     []rune{'a', 'z', 'A', 'Z', '0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
						},
					},
				},
			},
		},
		{
			name: "Comment",
			pos:  position{line: 54, col: 1, offset: 1976},
			expr: &actionExpr{
				pos: position{line: 54, col: 12, offset: 1987},
				run: (*parser).callonComment1,
				expr: &seqExpr{
					pos: position{line: 54, col: 12, offset: 1987},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 54, col: 12, offset: 1987},
							val:        "//",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 54, col: 17, offset: 1992},
							expr: &charClassMatcher{
								pos:        position{line: 54, col: 17, offset: 1992},
								val:        "[^\\n\\r]",
								chars:      []rune{'\n', '\r'},
								ignoreCase: false,
								inverted:   true,
							},
						},
					},
				},
			},
		},
		{
			name:        "_",
			displayName: "\"whitespace\"",
			pos:         position{line: 56, col: 1, offset: 2039},
			expr: &zeroOrMoreExpr{
				pos: position{line: 56, col: 19, offset: 2057},
				expr: &charClassMatcher{
					pos:        position{line: 56, col: 19, offset: 2057},
					val:        "[ \\n\\r\\t]",
					chars:      []rune{' ', '\n', '\r', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
	},
}

func (c *current) onStart4(ins interface{}) (interface{}, error) {
	return Forward(ins)
}

func (p *parser) callonStart4() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStart4(stack["ins"])
}

func (c *current) onStart10(lbl interface{}) (interface{}, error) {
	return Forward(lbl)
}

func (p *parser) callonStart10() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStart10(stack["lbl"])
}

func (c *current) onStart13(comment interface{}) (interface{}, error) {
	return Forward(comment)
}

func (p *parser) callonStart13() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStart13(stack["comment"])
}

func (c *current) onInstruction2(set, cnd, rd, ra, rb interface{}) (interface{}, error) {
	return NewRegInstr(set, "add", cnd, rd, ra, rb)
}

func (p *parser) callonInstruction2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstruction2(stack["set"], stack["cnd"], stack["rd"], stack["ra"], stack["rb"])
}

func (c *current) onInstruction23(set, cnd, rd, ra, num interface{}) (interface{}, error) {
	return NewImm12Instr(set, "add", cnd, rd, ra, num)
}

func (p *parser) callonInstruction23() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstruction23(stack["set"], stack["cnd"], stack["rd"], stack["ra"], stack["num"])
}

func (c *current) onInstruction44(set, cnd, rd, rb interface{}) (interface{}, error) {
	return NewRegInstr(set, "add", cnd, rd, rd, rb)
}

func (p *parser) callonInstruction44() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstruction44(stack["set"], stack["cnd"], stack["rd"], stack["rb"])
}

func (c *current) onInstruction62(set, cnd, rd, num interface{}) (interface{}, error) {
	return NewImm16Instr(set, "add", cnd, rd, num)
}

func (p *parser) callonInstruction62() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstruction62(stack["set"], stack["cnd"], stack["rd"], stack["num"])
}

func (c *current) onInstruction80(cmd, rd, ra, rb interface{}) (interface{}, error) {
	return NewRegInstruction(cmd.(*Command), rd.(*Register), ra.(*Register), rb.(*Register))
}

func (p *parser) callonInstruction80() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstruction80(stack["cmd"], stack["rd"], stack["ra"], stack["rb"])
}

func (c *current) onInstruction102(cmd, rd, ra, num interface{}) (interface{}, error) {
	return NewImmInstruction(cmd.(*Command), rd.(*Register), ra.(*Register), num.(*Number))
}

func (p *parser) callonInstruction102() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstruction102(stack["cmd"], stack["rd"], stack["ra"], stack["num"])
}

func (c *current) onInstruction124(cmd, lbl interface{}) (interface{}, error) {
	return NewBraInstruction(cmd.(*Command), lbl.(*Label))
}

func (p *parser) callonInstruction124() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstruction124(stack["cmd"], stack["lbl"])
}

func (c *current) onCommand1(set, cmd, cnd interface{}) (interface{}, error) {
	return NewCommand(set, cmd, cnd)
}

func (p *parser) callonCommand1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCommand1(stack["set"], stack["cmd"], stack["cnd"])
}

func (c *current) onMnemonic1() (interface{}, error) {
	return NewString(c.text)
}

func (p *parser) callonMnemonic1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMnemonic1()
}

func (c *current) onCondition1() (interface{}, error) {
	return NewString(c.text)
}

func (p *parser) callonCondition1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCondition1()
}

func (c *current) onRegister1() (interface{}, error) {
	return NewRegister(c.text)
}

func (p *parser) callonRegister1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRegister1()
}

func (c *current) onNumber1(num interface{}) (interface{}, error) {
	return Forward(num)
}

func (p *parser) callonNumber1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onNumber1(stack["num"])
}

func (c *current) onDecNumber1() (interface{}, error) {
	return NewNumber(c.text, 10)
}

func (p *parser) callonDecNumber1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDecNumber1()
}

func (c *current) onHexNumber1() (interface{}, error) {
	return NewNumber(c.text, 16)
}

func (p *parser) callonHexNumber1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onHexNumber1()
}

func (c *current) onLabel1() (interface{}, error) {
	return NewLabel(c.text)
}

func (p *parser) callonLabel1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLabel1()
}

func (c *current) onComment1() (interface{}, error) {
	return NewComment()
}

func (p *parser) callonComment1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onComment1()
}

var (
	// errNoRule is returned when the grammar to parse has no rule.
	errNoRule = errors.New("grammar has no rule")

	// errInvalidEntrypoint is returned when the specified entrypoint rule
	// does not exit.
	errInvalidEntrypoint = errors.New("invalid entrypoint")

	// errInvalidEncoding is returned when the source is not properly
	// utf8-encoded.
	errInvalidEncoding = errors.New("invalid encoding")

	// errMaxExprCnt is used to signal that the maximum number of
	// expressions have been parsed.
	errMaxExprCnt = errors.New("max number of expresssions parsed")
)

// Option is a function that can set an option on the parser. It returns
// the previous setting as an Option.
type Option func(*parser) Option

// MaxExpressions creates an Option to stop parsing after the provided
// number of expressions have been parsed, if the value is 0 then the parser will
// parse for as many steps as needed (possibly an infinite number).
//
// The default for maxExprCnt is 0.
func MaxExpressions(maxExprCnt uint64) Option {
	return func(p *parser) Option {
		oldMaxExprCnt := p.maxExprCnt
		p.maxExprCnt = maxExprCnt
		return MaxExpressions(oldMaxExprCnt)
	}
}

// Entrypoint creates an Option to set the rule name to use as entrypoint.
// The rule name must have been specified in the -alternate-entrypoints
// if generating the parser with the -optimize-grammar flag, otherwise
// it may have been optimized out. Passing an empty string sets the
// entrypoint to the first rule in the grammar.
//
// The default is to start parsing at the first rule in the grammar.
func Entrypoint(ruleName string) Option {
	return func(p *parser) Option {
		oldEntrypoint := p.entrypoint
		p.entrypoint = ruleName
		if ruleName == "" {
			p.entrypoint = g.rules[0].name
		}
		return Entrypoint(oldEntrypoint)
	}
}

// Statistics adds a user provided Stats struct to the parser to allow
// the user to process the results after the parsing has finished.
// Also the key for the "no match" counter is set.
//
// Example usage:
//
//     input := "input"
//     stats := Stats{}
//     _, err := Parse("input-file", []byte(input), Statistics(&stats, "no match"))
//     if err != nil {
//         log.Panicln(err)
//     }
//     b, err := json.MarshalIndent(stats.ChoiceAltCnt, "", "  ")
//     if err != nil {
//         log.Panicln(err)
//     }
//     fmt.Println(string(b))
//
func Statistics(stats *Stats, choiceNoMatch string) Option {
	return func(p *parser) Option {
		oldStats := p.Stats
		p.Stats = stats
		oldChoiceNoMatch := p.choiceNoMatch
		p.choiceNoMatch = choiceNoMatch
		if p.Stats.ChoiceAltCnt == nil {
			p.Stats.ChoiceAltCnt = make(map[string]map[string]int)
		}
		return Statistics(oldStats, oldChoiceNoMatch)
	}
}

// Debug creates an Option to set the debug flag to b. When set to true,
// debugging information is printed to stdout while parsing.
//
// The default is false.
func Debug(b bool) Option {
	return func(p *parser) Option {
		old := p.debug
		p.debug = b
		return Debug(old)
	}
}

// Memoize creates an Option to set the memoize flag to b. When set to true,
// the parser will cache all results so each expression is evaluated only
// once. This guarantees linear parsing time even for pathological cases,
// at the expense of more memory and slower times for typical cases.
//
// The default is false.
func Memoize(b bool) Option {
	return func(p *parser) Option {
		old := p.memoize
		p.memoize = b
		return Memoize(old)
	}
}

// AllowInvalidUTF8 creates an Option to allow invalid UTF-8 bytes.
// Every invalid UTF-8 byte is treated as a utf8.RuneError (U+FFFD)
// by character class matchers and is matched by the any matcher.
// The returned matched value, c.text and c.offset are NOT affected.
//
// The default is false.
func AllowInvalidUTF8(b bool) Option {
	return func(p *parser) Option {
		old := p.allowInvalidUTF8
		p.allowInvalidUTF8 = b
		return AllowInvalidUTF8(old)
	}
}

// Recover creates an Option to set the recover flag to b. When set to
// true, this causes the parser to recover from panics and convert it
// to an error. Setting it to false can be useful while debugging to
// access the full stack trace.
//
// The default is true.
func Recover(b bool) Option {
	return func(p *parser) Option {
		old := p.recover
		p.recover = b
		return Recover(old)
	}
}

// GlobalStore creates an Option to set a key to a certain value in
// the globalStore.
func GlobalStore(key string, value interface{}) Option {
	return func(p *parser) Option {
		old := p.cur.globalStore[key]
		p.cur.globalStore[key] = value
		return GlobalStore(key, old)
	}
}

// InitState creates an Option to set a key to a certain value in
// the global "state" store.
func InitState(key string, value interface{}) Option {
	return func(p *parser) Option {
		old := p.cur.state[key]
		p.cur.state[key] = value
		return InitState(key, old)
	}
}

// ParseFile parses the file identified by filename.
func ParseFile(filename string, opts ...Option) (i interface{}, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			err = closeErr
		}
	}()
	return ParseReader(filename, f, opts...)
}

// ParseReader parses the data from r using filename as information in the
// error messages.
func ParseReader(filename string, r io.Reader, opts ...Option) (interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse(filename, b, opts...)
}

// Parse parses the data from b using filename as information in the
// error messages.
func Parse(filename string, b []byte, opts ...Option) (interface{}, error) {
	return newParser(filename, b, opts...).parse(g)
}

// position records a position in the text.
type position struct {
	line, col, offset int
}

func (p position) String() string {
	return fmt.Sprintf("%d:%d [%d]", p.line, p.col, p.offset)
}

// savepoint stores all state required to go back to this point in the
// parser.
type savepoint struct {
	position
	rn rune
	w  int
}

type current struct {
	pos  position // start position of the match
	text []byte   // raw text of the match

	// state is a store for arbitrary key,value pairs that the user wants to be
	// tied to the backtracking of the parser.
	// This is always rolled back if a parsing rule fails.
	state storeDict

	// globalStore is a general store for the user to store arbitrary key-value
	// pairs that they need to manage and that they do not want tied to the
	// backtracking of the parser. This is only modified by the user and never
	// rolled back by the parser. It is always up to the user to keep this in a
	// consistent state.
	globalStore storeDict
}

type storeDict map[string]interface{}

// the AST types...

type grammar struct {
	pos   position
	rules []*rule
}

type rule struct {
	pos         position
	name        string
	displayName string
	expr        interface{}
}

type choiceExpr struct {
	pos          position
	alternatives []interface{}
}

type actionExpr struct {
	pos  position
	expr interface{}
	run  func(*parser) (interface{}, error)
}

type recoveryExpr struct {
	pos          position
	expr         interface{}
	recoverExpr  interface{}
	failureLabel []string
}

type seqExpr struct {
	pos   position
	exprs []interface{}
}

type throwExpr struct {
	pos   position
	label string
}

type labeledExpr struct {
	pos   position
	label string
	expr  interface{}
}

type expr struct {
	pos  position
	expr interface{}
}

type andExpr expr
type notExpr expr
type zeroOrOneExpr expr
type zeroOrMoreExpr expr
type oneOrMoreExpr expr

type ruleRefExpr struct {
	pos  position
	name string
}

type stateCodeExpr struct {
	pos position
	run func(*parser) error
}

type andCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type notCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type litMatcher struct {
	pos        position
	val        string
	ignoreCase bool
}

type charClassMatcher struct {
	pos             position
	val             string
	basicLatinChars [128]bool
	chars           []rune
	ranges          []rune
	classes         []*unicode.RangeTable
	ignoreCase      bool
	inverted        bool
}

type anyMatcher position

// errList cumulates the errors found by the parser.
type errList []error

func (e *errList) add(err error) {
	*e = append(*e, err)
}

func (e errList) err() error {
	if len(e) == 0 {
		return nil
	}
	e.dedupe()
	return e
}

func (e *errList) dedupe() {
	var cleaned []error
	set := make(map[string]bool)
	for _, err := range *e {
		if msg := err.Error(); !set[msg] {
			set[msg] = true
			cleaned = append(cleaned, err)
		}
	}
	*e = cleaned
}

func (e errList) Error() string {
	switch len(e) {
	case 0:
		return ""
	case 1:
		return e[0].Error()
	default:
		var buf bytes.Buffer

		for i, err := range e {
			if i > 0 {
				buf.WriteRune('\n')
			}
			buf.WriteString(err.Error())
		}
		return buf.String()
	}
}

// parserError wraps an error with a prefix indicating the rule in which
// the error occurred. The original error is stored in the Inner field.
type parserError struct {
	Inner    error
	pos      position
	prefix   string
	expected []string
}

// Error returns the error message.
func (p *parserError) Error() string {
	return p.prefix + ": " + p.Inner.Error()
}

// newParser creates a parser with the specified input source and options.
func newParser(filename string, b []byte, opts ...Option) *parser {
	stats := Stats{
		ChoiceAltCnt: make(map[string]map[string]int),
	}

	p := &parser{
		filename: filename,
		errs:     new(errList),
		data:     b,
		pt:       savepoint{position: position{line: 1}},
		recover:  true,
		cur: current{
			state:       make(storeDict),
			globalStore: make(storeDict),
		},
		maxFailPos:      position{col: 1, line: 1},
		maxFailExpected: make([]string, 0, 20),
		Stats:           &stats,
		// start rule is rule [0] unless an alternate entrypoint is specified
		entrypoint: g.rules[0].name,
		emptyState: make(storeDict),
	}
	p.setOptions(opts)

	if p.maxExprCnt == 0 {
		p.maxExprCnt = math.MaxUint64
	}

	return p
}

// setOptions applies the options to the parser.
func (p *parser) setOptions(opts []Option) {
	for _, opt := range opts {
		opt(p)
	}
}

type resultTuple struct {
	v   interface{}
	b   bool
	end savepoint
}

const choiceNoMatch = -1

// Stats stores some statistics, gathered during parsing
type Stats struct {
	// ExprCnt counts the number of expressions processed during parsing
	// This value is compared to the maximum number of expressions allowed
	// (set by the MaxExpressions option).
	ExprCnt uint64

	// ChoiceAltCnt is used to count for each ordered choice expression,
	// which alternative is used how may times.
	// These numbers allow to optimize the order of the ordered choice expression
	// to increase the performance of the parser
	//
	// The outer key of ChoiceAltCnt is composed of the name of the rule as well
	// as the line and the column of the ordered choice.
	// The inner key of ChoiceAltCnt is the number (one-based) of the matching alternative.
	// For each alternative the number of matches are counted. If an ordered choice does not
	// match, a special counter is incremented. The name of this counter is set with
	// the parser option Statistics.
	// For an alternative to be included in ChoiceAltCnt, it has to match at least once.
	ChoiceAltCnt map[string]map[string]int
}

type parser struct {
	filename string
	pt       savepoint
	cur      current

	data []byte
	errs *errList

	depth   int
	recover bool
	debug   bool

	memoize bool
	// memoization table for the packrat algorithm:
	// map[offset in source] map[expression or rule] {value, match}
	memo map[int]map[interface{}]resultTuple

	// rules table, maps the rule identifier to the rule node
	rules map[string]*rule
	// variables stack, map of label to value
	vstack []map[string]interface{}
	// rule stack, allows identification of the current rule in errors
	rstack []*rule

	// parse fail
	maxFailPos            position
	maxFailExpected       []string
	maxFailInvertExpected bool

	// max number of expressions to be parsed
	maxExprCnt uint64
	// entrypoint for the parser
	entrypoint string

	allowInvalidUTF8 bool

	*Stats

	choiceNoMatch string
	// recovery expression stack, keeps track of the currently available recovery expression, these are traversed in reverse
	recoveryStack []map[string]interface{}

	// emptyState contains an empty storeDict, which is used to optimize cloneState if global "state" store is not used.
	emptyState storeDict
}

// push a variable set on the vstack.
func (p *parser) pushV() {
	if cap(p.vstack) == len(p.vstack) {
		// create new empty slot in the stack
		p.vstack = append(p.vstack, nil)
	} else {
		// slice to 1 more
		p.vstack = p.vstack[:len(p.vstack)+1]
	}

	// get the last args set
	m := p.vstack[len(p.vstack)-1]
	if m != nil && len(m) == 0 {
		// empty map, all good
		return
	}

	m = make(map[string]interface{})
	p.vstack[len(p.vstack)-1] = m
}

// pop a variable set from the vstack.
func (p *parser) popV() {
	// if the map is not empty, clear it
	m := p.vstack[len(p.vstack)-1]
	if len(m) > 0 {
		// GC that map
		p.vstack[len(p.vstack)-1] = nil
	}
	p.vstack = p.vstack[:len(p.vstack)-1]
}

// push a recovery expression with its labels to the recoveryStack
func (p *parser) pushRecovery(labels []string, expr interface{}) {
	if cap(p.recoveryStack) == len(p.recoveryStack) {
		// create new empty slot in the stack
		p.recoveryStack = append(p.recoveryStack, nil)
	} else {
		// slice to 1 more
		p.recoveryStack = p.recoveryStack[:len(p.recoveryStack)+1]
	}

	m := make(map[string]interface{}, len(labels))
	for _, fl := range labels {
		m[fl] = expr
	}
	p.recoveryStack[len(p.recoveryStack)-1] = m
}

// pop a recovery expression from the recoveryStack
func (p *parser) popRecovery() {
	// GC that map
	p.recoveryStack[len(p.recoveryStack)-1] = nil

	p.recoveryStack = p.recoveryStack[:len(p.recoveryStack)-1]
}

func (p *parser) print(prefix, s string) string {
	if !p.debug {
		return s
	}

	fmt.Printf("%s %d:%d:%d: %s [%#U]\n",
		prefix, p.pt.line, p.pt.col, p.pt.offset, s, p.pt.rn)
	return s
}

func (p *parser) in(s string) string {
	p.depth++
	return p.print(strings.Repeat(" ", p.depth)+">", s)
}

func (p *parser) out(s string) string {
	p.depth--
	return p.print(strings.Repeat(" ", p.depth)+"<", s)
}

func (p *parser) addErr(err error) {
	p.addErrAt(err, p.pt.position, []string{})
}

func (p *parser) addErrAt(err error, pos position, expected []string) {
	var buf bytes.Buffer
	if p.filename != "" {
		buf.WriteString(p.filename)
	}
	if buf.Len() > 0 {
		buf.WriteString(":")
	}
	buf.WriteString(fmt.Sprintf("%d:%d (%d)", pos.line, pos.col, pos.offset))
	if len(p.rstack) > 0 {
		if buf.Len() > 0 {
			buf.WriteString(": ")
		}
		rule := p.rstack[len(p.rstack)-1]
		if rule.displayName != "" {
			buf.WriteString("rule " + rule.displayName)
		} else {
			buf.WriteString("rule " + rule.name)
		}
	}
	pe := &parserError{Inner: err, pos: pos, prefix: buf.String(), expected: expected}
	p.errs.add(pe)
}

func (p *parser) failAt(fail bool, pos position, want string) {
	// process fail if parsing fails and not inverted or parsing succeeds and invert is set
	if fail == p.maxFailInvertExpected {
		if pos.offset < p.maxFailPos.offset {
			return
		}

		if pos.offset > p.maxFailPos.offset {
			p.maxFailPos = pos
			p.maxFailExpected = p.maxFailExpected[:0]
		}

		if p.maxFailInvertExpected {
			want = "!" + want
		}
		p.maxFailExpected = append(p.maxFailExpected, want)
	}
}

// read advances the parser to the next rune.
func (p *parser) read() {
	p.pt.offset += p.pt.w
	rn, n := utf8.DecodeRune(p.data[p.pt.offset:])
	p.pt.rn = rn
	p.pt.w = n
	p.pt.col++
	if rn == '\n' {
		p.pt.line++
		p.pt.col = 0
	}

	if rn == utf8.RuneError && n == 1 { // see utf8.DecodeRune
		if !p.allowInvalidUTF8 {
			p.addErr(errInvalidEncoding)
		}
	}
}

// restore parser position to the savepoint pt.
func (p *parser) restore(pt savepoint) {
	if p.debug {
		defer p.out(p.in("restore"))
	}
	if pt.offset == p.pt.offset {
		return
	}
	p.pt = pt
}

// Cloner is implemented by any value that has a Clone method, which returns a
// copy of the value. This is mainly used for types which are not passed by
// value (e.g map, slice, chan) or structs that contain such types.
//
// This is used in conjunction with the global state feature to create proper
// copies of the state to allow the parser to properly restore the state in
// the case of backtracking.
type Cloner interface {
	Clone() interface{}
}

// clone and return parser current state.
func (p *parser) cloneState() storeDict {
	if p.debug {
		defer p.out(p.in("cloneState"))
	}

	if len(p.cur.state) == 0 {
		if len(p.emptyState) > 0 {
			p.emptyState = make(storeDict)
		}
		return p.emptyState
	}

	state := make(storeDict, len(p.cur.state))
	for k, v := range p.cur.state {
		if c, ok := v.(Cloner); ok {
			state[k] = c.Clone()
		} else {
			state[k] = v
		}
	}
	return state
}

// restore parser current state to the state storeDict.
// every restoreState should applied only one time for every cloned state
func (p *parser) restoreState(state storeDict) {
	if p.debug {
		defer p.out(p.in("restoreState"))
	}
	p.cur.state = state
}

// get the slice of bytes from the savepoint start to the current position.
func (p *parser) sliceFrom(start savepoint) []byte {
	return p.data[start.position.offset:p.pt.position.offset]
}

func (p *parser) getMemoized(node interface{}) (resultTuple, bool) {
	if len(p.memo) == 0 {
		return resultTuple{}, false
	}
	m := p.memo[p.pt.offset]
	if len(m) == 0 {
		return resultTuple{}, false
	}
	res, ok := m[node]
	return res, ok
}

func (p *parser) setMemoized(pt savepoint, node interface{}, tuple resultTuple) {
	if p.memo == nil {
		p.memo = make(map[int]map[interface{}]resultTuple)
	}
	m := p.memo[pt.offset]
	if m == nil {
		m = make(map[interface{}]resultTuple)
		p.memo[pt.offset] = m
	}
	m[node] = tuple
}

func (p *parser) buildRulesTable(g *grammar) {
	p.rules = make(map[string]*rule, len(g.rules))
	for _, r := range g.rules {
		p.rules[r.name] = r
	}
}

func (p *parser) parse(g *grammar) (val interface{}, err error) {
	if len(g.rules) == 0 {
		p.addErr(errNoRule)
		return nil, p.errs.err()
	}

	// TODO : not super critical but this could be generated
	p.buildRulesTable(g)

	if p.recover {
		// panic can be used in action code to stop parsing immediately
		// and return the panic as an error.
		defer func() {
			if e := recover(); e != nil {
				if p.debug {
					defer p.out(p.in("panic handler"))
				}
				val = nil
				switch e := e.(type) {
				case error:
					p.addErr(e)
				default:
					p.addErr(fmt.Errorf("%v", e))
				}
				err = p.errs.err()
			}
		}()
	}

	startRule, ok := p.rules[p.entrypoint]
	if !ok {
		p.addErr(errInvalidEntrypoint)
		return nil, p.errs.err()
	}

	p.read() // advance to first rune
	val, ok = p.parseRule(startRule)
	if !ok {
		if len(*p.errs) == 0 {
			// If parsing fails, but no errors have been recorded, the expected values
			// for the farthest parser position are returned as error.
			maxFailExpectedMap := make(map[string]struct{}, len(p.maxFailExpected))
			for _, v := range p.maxFailExpected {
				maxFailExpectedMap[v] = struct{}{}
			}
			expected := make([]string, 0, len(maxFailExpectedMap))
			eof := false
			if _, ok := maxFailExpectedMap["!."]; ok {
				delete(maxFailExpectedMap, "!.")
				eof = true
			}
			for k := range maxFailExpectedMap {
				expected = append(expected, k)
			}
			sort.Strings(expected)
			if eof {
				expected = append(expected, "EOF")
			}
			p.addErrAt(errors.New("no match found, expected: "+listJoin(expected, ", ", "or")), p.maxFailPos, expected)
		}

		return nil, p.errs.err()
	}
	return val, p.errs.err()
}

func listJoin(list []string, sep string, lastSep string) string {
	switch len(list) {
	case 0:
		return ""
	case 1:
		return list[0]
	default:
		return fmt.Sprintf("%s %s %s", strings.Join(list[:len(list)-1], sep), lastSep, list[len(list)-1])
	}
}

func (p *parser) parseRule(rule *rule) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRule " + rule.name))
	}

	if p.memoize {
		res, ok := p.getMemoized(rule)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
	}

	start := p.pt
	p.rstack = append(p.rstack, rule)
	p.pushV()
	val, ok := p.parseExpr(rule.expr)
	p.popV()
	p.rstack = p.rstack[:len(p.rstack)-1]
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}

	if p.memoize {
		p.setMemoized(start, rule, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseExpr(expr interface{}) (interface{}, bool) {
	var pt savepoint

	if p.memoize {
		res, ok := p.getMemoized(expr)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
		pt = p.pt
	}

	p.ExprCnt++
	if p.ExprCnt > p.maxExprCnt {
		panic(errMaxExprCnt)
	}

	var val interface{}
	var ok bool
	switch expr := expr.(type) {
	case *actionExpr:
		val, ok = p.parseActionExpr(expr)
	case *andCodeExpr:
		val, ok = p.parseAndCodeExpr(expr)
	case *andExpr:
		val, ok = p.parseAndExpr(expr)
	case *anyMatcher:
		val, ok = p.parseAnyMatcher(expr)
	case *charClassMatcher:
		val, ok = p.parseCharClassMatcher(expr)
	case *choiceExpr:
		val, ok = p.parseChoiceExpr(expr)
	case *labeledExpr:
		val, ok = p.parseLabeledExpr(expr)
	case *litMatcher:
		val, ok = p.parseLitMatcher(expr)
	case *notCodeExpr:
		val, ok = p.parseNotCodeExpr(expr)
	case *notExpr:
		val, ok = p.parseNotExpr(expr)
	case *oneOrMoreExpr:
		val, ok = p.parseOneOrMoreExpr(expr)
	case *recoveryExpr:
		val, ok = p.parseRecoveryExpr(expr)
	case *ruleRefExpr:
		val, ok = p.parseRuleRefExpr(expr)
	case *seqExpr:
		val, ok = p.parseSeqExpr(expr)
	case *stateCodeExpr:
		val, ok = p.parseStateCodeExpr(expr)
	case *throwExpr:
		val, ok = p.parseThrowExpr(expr)
	case *zeroOrMoreExpr:
		val, ok = p.parseZeroOrMoreExpr(expr)
	case *zeroOrOneExpr:
		val, ok = p.parseZeroOrOneExpr(expr)
	default:
		panic(fmt.Sprintf("unknown expression type %T", expr))
	}
	if p.memoize {
		p.setMemoized(pt, expr, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseActionExpr(act *actionExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseActionExpr"))
	}

	start := p.pt
	val, ok := p.parseExpr(act.expr)
	if ok {
		p.cur.pos = start.position
		p.cur.text = p.sliceFrom(start)
		state := p.cloneState()
		actVal, err := act.run(p)
		if err != nil {
			p.addErrAt(err, start.position, []string{})
		}
		p.restoreState(state)

		val = actVal
	}
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}
	return val, ok
}

func (p *parser) parseAndCodeExpr(and *andCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndCodeExpr"))
	}

	state := p.cloneState()

	ok, err := and.run(p)
	if err != nil {
		p.addErr(err)
	}
	p.restoreState(state)

	return nil, ok
}

func (p *parser) parseAndExpr(and *andExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndExpr"))
	}

	pt := p.pt
	state := p.cloneState()
	p.pushV()
	_, ok := p.parseExpr(and.expr)
	p.popV()
	p.restoreState(state)
	p.restore(pt)

	return nil, ok
}

func (p *parser) parseAnyMatcher(any *anyMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAnyMatcher"))
	}

	if p.pt.rn == utf8.RuneError && p.pt.w == 0 {
		// EOF - see utf8.DecodeRune
		p.failAt(false, p.pt.position, ".")
		return nil, false
	}
	start := p.pt
	p.read()
	p.failAt(true, start.position, ".")
	return p.sliceFrom(start), true
}

func (p *parser) parseCharClassMatcher(chr *charClassMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseCharClassMatcher"))
	}

	cur := p.pt.rn
	start := p.pt

	// can't match EOF
	if cur == utf8.RuneError && p.pt.w == 0 { // see utf8.DecodeRune
		p.failAt(false, start.position, chr.val)
		return nil, false
	}

	if chr.ignoreCase {
		cur = unicode.ToLower(cur)
	}

	// try to match in the list of available chars
	for _, rn := range chr.chars {
		if rn == cur {
			if chr.inverted {
				p.failAt(false, start.position, chr.val)
				return nil, false
			}
			p.read()
			p.failAt(true, start.position, chr.val)
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of ranges
	for i := 0; i < len(chr.ranges); i += 2 {
		if cur >= chr.ranges[i] && cur <= chr.ranges[i+1] {
			if chr.inverted {
				p.failAt(false, start.position, chr.val)
				return nil, false
			}
			p.read()
			p.failAt(true, start.position, chr.val)
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of Unicode classes
	for _, cl := range chr.classes {
		if unicode.Is(cl, cur) {
			if chr.inverted {
				p.failAt(false, start.position, chr.val)
				return nil, false
			}
			p.read()
			p.failAt(true, start.position, chr.val)
			return p.sliceFrom(start), true
		}
	}

	if chr.inverted {
		p.read()
		p.failAt(true, start.position, chr.val)
		return p.sliceFrom(start), true
	}
	p.failAt(false, start.position, chr.val)
	return nil, false
}

func (p *parser) incChoiceAltCnt(ch *choiceExpr, altI int) {
	choiceIdent := fmt.Sprintf("%s %d:%d", p.rstack[len(p.rstack)-1].name, ch.pos.line, ch.pos.col)
	m := p.ChoiceAltCnt[choiceIdent]
	if m == nil {
		m = make(map[string]int)
		p.ChoiceAltCnt[choiceIdent] = m
	}
	// We increment altI by 1, so the keys do not start at 0
	alt := strconv.Itoa(altI + 1)
	if altI == choiceNoMatch {
		alt = p.choiceNoMatch
	}
	m[alt]++
}

func (p *parser) parseChoiceExpr(ch *choiceExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseChoiceExpr"))
	}

	for altI, alt := range ch.alternatives {
		// dummy assignment to prevent compile error if optimized
		_ = altI

		state := p.cloneState()

		p.pushV()
		val, ok := p.parseExpr(alt)
		p.popV()
		if ok {
			p.incChoiceAltCnt(ch, altI)
			return val, ok
		}
		p.restoreState(state)
	}
	p.incChoiceAltCnt(ch, choiceNoMatch)
	return nil, false
}

func (p *parser) parseLabeledExpr(lab *labeledExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLabeledExpr"))
	}

	p.pushV()
	val, ok := p.parseExpr(lab.expr)
	p.popV()
	if ok && lab.label != "" {
		m := p.vstack[len(p.vstack)-1]
		m[lab.label] = val
	}
	return val, ok
}

func (p *parser) parseLitMatcher(lit *litMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLitMatcher"))
	}

	ignoreCase := ""
	if lit.ignoreCase {
		ignoreCase = "i"
	}
	val := fmt.Sprintf("%q%s", lit.val, ignoreCase)
	start := p.pt
	for _, want := range lit.val {
		cur := p.pt.rn
		if lit.ignoreCase {
			cur = unicode.ToLower(cur)
		}
		if cur != want {
			p.failAt(false, start.position, val)
			p.restore(start)
			return nil, false
		}
		p.read()
	}
	p.failAt(true, start.position, val)
	return p.sliceFrom(start), true
}

func (p *parser) parseNotCodeExpr(not *notCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotCodeExpr"))
	}

	state := p.cloneState()

	ok, err := not.run(p)
	if err != nil {
		p.addErr(err)
	}
	p.restoreState(state)

	return nil, !ok
}

func (p *parser) parseNotExpr(not *notExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotExpr"))
	}

	pt := p.pt
	state := p.cloneState()
	p.pushV()
	p.maxFailInvertExpected = !p.maxFailInvertExpected
	_, ok := p.parseExpr(not.expr)
	p.maxFailInvertExpected = !p.maxFailInvertExpected
	p.popV()
	p.restoreState(state)
	p.restore(pt)

	return nil, !ok
}

func (p *parser) parseOneOrMoreExpr(expr *oneOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseOneOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			if len(vals) == 0 {
				// did not match once, no match
				return nil, false
			}
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseRecoveryExpr(recover *recoveryExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRecoveryExpr (" + strings.Join(recover.failureLabel, ",") + ")"))
	}

	p.pushRecovery(recover.failureLabel, recover.recoverExpr)
	val, ok := p.parseExpr(recover.expr)
	p.popRecovery()

	return val, ok
}

func (p *parser) parseRuleRefExpr(ref *ruleRefExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRuleRefExpr " + ref.name))
	}

	if ref.name == "" {
		panic(fmt.Sprintf("%s: invalid rule: missing name", ref.pos))
	}

	rule := p.rules[ref.name]
	if rule == nil {
		p.addErr(fmt.Errorf("undefined rule: %s", ref.name))
		return nil, false
	}
	return p.parseRule(rule)
}

func (p *parser) parseSeqExpr(seq *seqExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseSeqExpr"))
	}

	vals := make([]interface{}, 0, len(seq.exprs))

	pt := p.pt
	state := p.cloneState()
	for _, expr := range seq.exprs {
		val, ok := p.parseExpr(expr)
		if !ok {
			p.restoreState(state)
			p.restore(pt)
			return nil, false
		}
		vals = append(vals, val)
	}
	return vals, true
}

func (p *parser) parseStateCodeExpr(state *stateCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseStateCodeExpr"))
	}

	err := state.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, true
}

func (p *parser) parseThrowExpr(expr *throwExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseThrowExpr"))
	}

	for i := len(p.recoveryStack) - 1; i >= 0; i-- {
		if recoverExpr, ok := p.recoveryStack[i][expr.label]; ok {
			if val, ok := p.parseExpr(recoverExpr); ok {
				return val, ok
			}
		}
	}

	return nil, false
}

func (p *parser) parseZeroOrMoreExpr(expr *zeroOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseZeroOrOneExpr(expr *zeroOrOneExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrOneExpr"))
	}

	p.pushV()
	val, _ := p.parseExpr(expr.expr)
	p.popV()
	// whether it matched or not, consider it a match
	return val, true
}