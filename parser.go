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

	"github.com/mhoertnagl/epic-esm/gen"
)

var g = &grammar{
	rules: []*rule{
		{
			name: "Start",
			pos:  position{line: 7, col: 1, offset: 73},
			expr: &seqExpr{
				pos: position{line: 8, col: 6, offset: 85},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 8, col: 6, offset: 85},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 9, col: 8, offset: 97},
						alternatives: []interface{}{
							&actionExpr{
								pos: position{line: 9, col: 8, offset: 97},
								run: (*parser).callonStart4,
								expr: &seqExpr{
									pos: position{line: 9, col: 8, offset: 97},
									exprs: []interface{}{
										&labeledExpr{
											pos:   position{line: 9, col: 8, offset: 97},
											label: "ins",
											expr: &ruleRefExpr{
												pos:  position{line: 9, col: 12, offset: 101},
												name: "Instr",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 9, col: 18, offset: 107},
											name: "_",
										},
										&zeroOrOneExpr{
											pos: position{line: 9, col: 20, offset: 109},
											expr: &ruleRefExpr{
												pos:  position{line: 9, col: 20, offset: 109},
												name: "Comment",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 9, col: 29, offset: 118},
											name: "_",
										},
									},
								},
							},
							&actionExpr{
								pos: position{line: 10, col: 8, offset: 233},
								run: (*parser).callonStart12,
								expr: &seqExpr{
									pos: position{line: 10, col: 8, offset: 233},
									exprs: []interface{}{
										&labeledExpr{
											pos:   position{line: 10, col: 8, offset: 233},
											label: "lbl",
											expr: &ruleRefExpr{
												pos:  position{line: 10, col: 12, offset: 237},
												name: "Label",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 10, col: 18, offset: 243},
											name: "_",
										},
										&zeroOrOneExpr{
											pos: position{line: 10, col: 20, offset: 245},
											expr: &ruleRefExpr{
												pos:  position{line: 10, col: 20, offset: 245},
												name: "Comment",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 10, col: 29, offset: 254},
											name: "_",
										},
									},
								},
							},
							&actionExpr{
								pos: position{line: 11, col: 8, offset: 369},
								run: (*parser).callonStart20,
								expr: &seqExpr{
									pos: position{line: 11, col: 8, offset: 369},
									exprs: []interface{}{
										&labeledExpr{
											pos:   position{line: 11, col: 8, offset: 369},
											label: "cmt",
											expr: &ruleRefExpr{
												pos:  position{line: 11, col: 12, offset: 373},
												name: "Comment",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 11, col: 20, offset: 381},
											name: "_",
										},
									},
								},
							},
						},
					},
					&notExpr{
						pos: position{line: 12, col: 6, offset: 503},
						expr: &anyMatcher{
							line: 12, col: 7, offset: 504,
						},
					},
				},
			},
		},
		{
			name:        "Instr",
			displayName: "\"instruction\"",
			pos:         position{line: 14, col: 1, offset: 509},
			expr: &choiceExpr{
				pos: position{line: 15, col: 15, offset: 546},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 15, col: 15, offset: 546},
						run: (*parser).callonInstr2,
						expr: &seqExpr{
							pos: position{line: 15, col: 15, offset: 546},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 15, col: 15, offset: 546},
									val:        "nop",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 15, col: 22, offset: 553},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 15, col: 25, offset: 556},
										expr: &ruleRefExpr{
											pos:  position{line: 15, col: 25, offset: 556},
											name: "Cnd",
										},
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 17, col: 6, offset: 679},
						run: (*parser).callonInstr8,
						expr: &seqExpr{
							pos: position{line: 17, col: 6, offset: 679},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 17, col: 6, offset: 679},
									label: "s",
									expr: &zeroOrOneExpr{
										pos: position{line: 17, col: 8, offset: 681},
										expr: &litMatcher{
											pos:        position{line: 17, col: 8, offset: 681},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 17, col: 13, offset: 686},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 17, col: 15, offset: 688},
									val:        "clr",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 17, col: 22, offset: 695},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 17, col: 25, offset: 698},
										expr: &ruleRefExpr{
											pos:  position{line: 17, col: 25, offset: 698},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 17, col: 30, offset: 703},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 17, col: 32, offset: 705},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 17, col: 35, offset: 708},
										name: "Reg",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 21, col: 6, offset: 955},
						run: (*parser).callonInstr21,
						expr: &seqExpr{
							pos: position{line: 21, col: 6, offset: 955},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 21, col: 6, offset: 955},
									label: "s",
									expr: &zeroOrOneExpr{
										pos: position{line: 21, col: 8, offset: 957},
										expr: &litMatcher{
											pos:        position{line: 21, col: 8, offset: 957},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 21, col: 13, offset: 962},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 21, col: 15, offset: 964},
									val:        "mov",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 21, col: 22, offset: 971},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 21, col: 25, offset: 974},
										expr: &ruleRefExpr{
											pos:  position{line: 21, col: 25, offset: 974},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 21, col: 30, offset: 979},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 21, col: 32, offset: 981},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 21, col: 35, offset: 984},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 21, col: 39, offset: 988},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 21, col: 41, offset: 990},
									label: "rb",
									expr: &ruleRefExpr{
										pos:  position{line: 21, col: 44, offset: 993},
										name: "Reg",
									},
								},
								&labeledExpr{
									pos:   position{line: 21, col: 48, offset: 997},
									label: "sh",
									expr: &zeroOrOneExpr{
										pos: position{line: 21, col: 51, offset: 1000},
										expr: &ruleRefExpr{
											pos:  position{line: 21, col: 51, offset: 1000},
											name: "Shift",
										},
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 22, col: 6, offset: 1095},
						run: (*parser).callonInstr40,
						expr: &seqExpr{
							pos: position{line: 22, col: 6, offset: 1095},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 22, col: 6, offset: 1095},
									label: "s",
									expr: &zeroOrOneExpr{
										pos: position{line: 22, col: 8, offset: 1097},
										expr: &litMatcher{
											pos:        position{line: 22, col: 8, offset: 1097},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 22, col: 13, offset: 1102},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 22, col: 15, offset: 1104},
									val:        "ldc",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 22, col: 22, offset: 1111},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 22, col: 25, offset: 1114},
										expr: &ruleRefExpr{
											pos:  position{line: 22, col: 25, offset: 1114},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 22, col: 30, offset: 1119},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 22, col: 32, offset: 1121},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 22, col: 35, offset: 1124},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 22, col: 39, offset: 1128},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 22, col: 41, offset: 1130},
									label: "nm",
									expr: &ruleRefExpr{
										pos:  position{line: 22, col: 44, offset: 1133},
										name: "Num",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 23, col: 6, offset: 1235},
						run: (*parser).callonInstr56,
						expr: &seqExpr{
							pos: position{line: 23, col: 6, offset: 1235},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 23, col: 6, offset: 1235},
									label: "s",
									expr: &zeroOrOneExpr{
										pos: position{line: 23, col: 8, offset: 1237},
										expr: &litMatcher{
											pos:        position{line: 23, col: 8, offset: 1237},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 23, col: 13, offset: 1242},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 23, col: 15, offset: 1244},
									val:        "lda",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 23, col: 22, offset: 1251},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 23, col: 25, offset: 1254},
										expr: &ruleRefExpr{
											pos:  position{line: 23, col: 25, offset: 1254},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 23, col: 30, offset: 1259},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 23, col: 32, offset: 1261},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 23, col: 35, offset: 1264},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 23, col: 39, offset: 1268},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 23, col: 41, offset: 1270},
									label: "lb",
									expr: &ruleRefExpr{
										pos:  position{line: 23, col: 44, offset: 1273},
										name: "Label",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 30, col: 6, offset: 1783},
						run: (*parser).callonInstr72,
						expr: &seqExpr{
							pos: position{line: 30, col: 6, offset: 1783},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 30, col: 6, offset: 1783},
									label: "s",
									expr: &zeroOrOneExpr{
										pos: position{line: 30, col: 8, offset: 1785},
										expr: &litMatcher{
											pos:        position{line: 30, col: 8, offset: 1785},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 30, col: 13, offset: 1790},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 30, col: 15, offset: 1792},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 30, col: 18, offset: 1795},
										name: "SOp",
									},
								},
								&labeledExpr{
									pos:   position{line: 30, col: 22, offset: 1799},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 30, col: 25, offset: 1802},
										expr: &ruleRefExpr{
											pos:  position{line: 30, col: 25, offset: 1802},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 30, col: 30, offset: 1807},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 30, col: 32, offset: 1809},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 30, col: 35, offset: 1812},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 30, col: 39, offset: 1816},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 30, col: 41, offset: 1818},
									label: "rb",
									expr: &ruleRefExpr{
										pos:  position{line: 30, col: 44, offset: 1821},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 30, col: 48, offset: 1825},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 30, col: 50, offset: 1827},
									label: "nm",
									expr: &ruleRefExpr{
										pos:  position{line: 30, col: 53, offset: 1830},
										name: "Num",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 36, col: 6, offset: 1951},
						run: (*parser).callonInstr92,
						expr: &seqExpr{
							pos: position{line: 36, col: 6, offset: 1951},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 36, col: 6, offset: 1951},
									label: "s",
									expr: &zeroOrOneExpr{
										pos: position{line: 36, col: 8, offset: 1953},
										expr: &litMatcher{
											pos:        position{line: 36, col: 8, offset: 1953},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 36, col: 13, offset: 1958},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 36, col: 15, offset: 1960},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 36, col: 18, offset: 1963},
										name: "DOp",
									},
								},
								&labeledExpr{
									pos:   position{line: 36, col: 22, offset: 1967},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 36, col: 25, offset: 1970},
										expr: &ruleRefExpr{
											pos:  position{line: 36, col: 25, offset: 1970},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 36, col: 30, offset: 1975},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 36, col: 32, offset: 1977},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 36, col: 35, offset: 1980},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 36, col: 39, offset: 1984},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 36, col: 41, offset: 1986},
									label: "ra",
									expr: &ruleRefExpr{
										pos:  position{line: 36, col: 44, offset: 1989},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 36, col: 48, offset: 1993},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 36, col: 50, offset: 1995},
									label: "rb",
									expr: &ruleRefExpr{
										pos:  position{line: 36, col: 53, offset: 1998},
										name: "Reg",
									},
								},
								&labeledExpr{
									pos:   position{line: 36, col: 57, offset: 2002},
									label: "sh",
									expr: &zeroOrOneExpr{
										pos: position{line: 36, col: 60, offset: 2005},
										expr: &ruleRefExpr{
											pos:  position{line: 36, col: 60, offset: 2005},
											name: "Shift",
										},
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 37, col: 6, offset: 2091},
						run: (*parser).callonInstr115,
						expr: &seqExpr{
							pos: position{line: 37, col: 6, offset: 2091},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 37, col: 6, offset: 2091},
									label: "s",
									expr: &zeroOrOneExpr{
										pos: position{line: 37, col: 8, offset: 2093},
										expr: &litMatcher{
											pos:        position{line: 37, col: 8, offset: 2093},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 37, col: 13, offset: 2098},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 37, col: 15, offset: 2100},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 37, col: 18, offset: 2103},
										name: "DOp",
									},
								},
								&labeledExpr{
									pos:   position{line: 37, col: 22, offset: 2107},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 37, col: 25, offset: 2110},
										expr: &ruleRefExpr{
											pos:  position{line: 37, col: 25, offset: 2110},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 37, col: 30, offset: 2115},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 37, col: 32, offset: 2117},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 37, col: 35, offset: 2120},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 37, col: 39, offset: 2124},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 37, col: 41, offset: 2126},
									label: "ra",
									expr: &ruleRefExpr{
										pos:  position{line: 37, col: 44, offset: 2129},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 37, col: 48, offset: 2133},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 37, col: 50, offset: 2135},
									label: "nm",
									expr: &ruleRefExpr{
										pos:  position{line: 37, col: 53, offset: 2138},
										name: "Num",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 38, col: 6, offset: 2231},
						run: (*parser).callonInstr135,
						expr: &seqExpr{
							pos: position{line: 38, col: 6, offset: 2231},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 38, col: 6, offset: 2231},
									label: "s",
									expr: &zeroOrOneExpr{
										pos: position{line: 38, col: 8, offset: 2233},
										expr: &litMatcher{
											pos:        position{line: 38, col: 8, offset: 2233},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 38, col: 13, offset: 2238},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 38, col: 15, offset: 2240},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 38, col: 18, offset: 2243},
										name: "DOp",
									},
								},
								&labeledExpr{
									pos:   position{line: 38, col: 22, offset: 2247},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 38, col: 25, offset: 2250},
										expr: &ruleRefExpr{
											pos:  position{line: 38, col: 25, offset: 2250},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 38, col: 30, offset: 2255},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 38, col: 32, offset: 2257},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 38, col: 35, offset: 2260},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 38, col: 39, offset: 2264},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 38, col: 41, offset: 2266},
									label: "nm",
									expr: &ruleRefExpr{
										pos:  position{line: 38, col: 44, offset: 2269},
										name: "Num",
									},
								},
								&labeledExpr{
									pos:   position{line: 38, col: 48, offset: 2273},
									label: "up",
									expr: &zeroOrOneExpr{
										pos: position{line: 38, col: 51, offset: 2276},
										expr: &seqExpr{
											pos: position{line: 38, col: 52, offset: 2277},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 38, col: 52, offset: 2277},
													name: "_",
												},
												&litMatcher{
													pos:        position{line: 38, col: 54, offset: 2279},
													val:        "<<",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 38, col: 59, offset: 2284},
													name: "_",
												},
												&litMatcher{
													pos:        position{line: 38, col: 61, offset: 2286},
													val:        "16",
													ignoreCase: false,
												},
											},
										},
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 40, col: 6, offset: 2373},
						run: (*parser).callonInstr159,
						expr: &seqExpr{
							pos: position{line: 40, col: 6, offset: 2373},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 40, col: 6, offset: 2373},
									label: "s",
									expr: &zeroOrOneExpr{
										pos: position{line: 40, col: 8, offset: 2375},
										expr: &litMatcher{
											pos:        position{line: 40, col: 8, offset: 2375},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 40, col: 13, offset: 2380},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 40, col: 15, offset: 2382},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 40, col: 18, offset: 2385},
										name: "COp",
									},
								},
								&labeledExpr{
									pos:   position{line: 40, col: 22, offset: 2389},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 40, col: 25, offset: 2392},
										expr: &ruleRefExpr{
											pos:  position{line: 40, col: 25, offset: 2392},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 40, col: 30, offset: 2397},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 40, col: 32, offset: 2399},
									label: "ra",
									expr: &ruleRefExpr{
										pos:  position{line: 40, col: 35, offset: 2402},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 40, col: 39, offset: 2406},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 40, col: 41, offset: 2408},
									label: "rb",
									expr: &ruleRefExpr{
										pos:  position{line: 40, col: 44, offset: 2411},
										name: "Reg",
									},
								},
								&labeledExpr{
									pos:   position{line: 40, col: 48, offset: 2415},
									label: "sh",
									expr: &zeroOrOneExpr{
										pos: position{line: 40, col: 51, offset: 2418},
										expr: &ruleRefExpr{
											pos:  position{line: 40, col: 51, offset: 2418},
											name: "Shift",
										},
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 41, col: 6, offset: 2513},
						run: (*parser).callonInstr179,
						expr: &seqExpr{
							pos: position{line: 41, col: 6, offset: 2513},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 41, col: 6, offset: 2513},
									label: "s",
									expr: &zeroOrOneExpr{
										pos: position{line: 41, col: 8, offset: 2515},
										expr: &litMatcher{
											pos:        position{line: 41, col: 8, offset: 2515},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 13, offset: 2520},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 41, col: 15, offset: 2522},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 41, col: 18, offset: 2525},
										name: "COp",
									},
								},
								&labeledExpr{
									pos:   position{line: 41, col: 22, offset: 2529},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 41, col: 25, offset: 2532},
										expr: &ruleRefExpr{
											pos:  position{line: 41, col: 25, offset: 2532},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 30, offset: 2537},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 41, col: 32, offset: 2539},
									label: "ra",
									expr: &ruleRefExpr{
										pos:  position{line: 41, col: 35, offset: 2542},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 39, offset: 2546},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 41, col: 41, offset: 2548},
									label: "nm",
									expr: &ruleRefExpr{
										pos:  position{line: 41, col: 44, offset: 2551},
										name: "Num",
									},
								},
								&labeledExpr{
									pos:   position{line: 41, col: 48, offset: 2555},
									label: "up",
									expr: &zeroOrOneExpr{
										pos: position{line: 41, col: 51, offset: 2558},
										expr: &seqExpr{
											pos: position{line: 41, col: 52, offset: 2559},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 41, col: 52, offset: 2559},
													name: "_",
												},
												&litMatcher{
													pos:        position{line: 41, col: 54, offset: 2561},
													val:        "<<",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 41, col: 59, offset: 2566},
													name: "_",
												},
												&litMatcher{
													pos:        position{line: 41, col: 61, offset: 2568},
													val:        "16",
													ignoreCase: false,
												},
											},
										},
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 43, col: 6, offset: 2655},
						run: (*parser).callonInstr203,
						expr: &seqExpr{
							pos: position{line: 43, col: 6, offset: 2655},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 43, col: 6, offset: 2655},
									label: "s",
									expr: &zeroOrOneExpr{
										pos: position{line: 43, col: 8, offset: 2657},
										expr: &litMatcher{
											pos:        position{line: 43, col: 8, offset: 2657},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 43, col: 13, offset: 2662},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 43, col: 15, offset: 2664},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 43, col: 18, offset: 2667},
										name: "MOp",
									},
								},
								&labeledExpr{
									pos:   position{line: 43, col: 22, offset: 2671},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 43, col: 25, offset: 2674},
										expr: &ruleRefExpr{
											pos:  position{line: 43, col: 25, offset: 2674},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 43, col: 30, offset: 2679},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 43, col: 32, offset: 2681},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 43, col: 35, offset: 2684},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 43, col: 39, offset: 2688},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 43, col: 41, offset: 2690},
									label: "ra",
									expr: &ruleRefExpr{
										pos:  position{line: 43, col: 44, offset: 2693},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 43, col: 48, offset: 2697},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 43, col: 50, offset: 2699},
									val:        "[",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 43, col: 54, offset: 2703},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 43, col: 56, offset: 2705},
									label: "rb",
									expr: &ruleRefExpr{
										pos:  position{line: 43, col: 59, offset: 2708},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 43, col: 63, offset: 2712},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 43, col: 65, offset: 2714},
									val:        "]",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 44, col: 6, offset: 2795},
						run: (*parser).callonInstr227,
						expr: &seqExpr{
							pos: position{line: 44, col: 6, offset: 2795},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 44, col: 6, offset: 2795},
									label: "s",
									expr: &zeroOrOneExpr{
										pos: position{line: 44, col: 8, offset: 2797},
										expr: &litMatcher{
											pos:        position{line: 44, col: 8, offset: 2797},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 44, col: 13, offset: 2802},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 44, col: 15, offset: 2804},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 44, col: 18, offset: 2807},
										name: "MOp",
									},
								},
								&labeledExpr{
									pos:   position{line: 44, col: 22, offset: 2811},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 44, col: 25, offset: 2814},
										expr: &ruleRefExpr{
											pos:  position{line: 44, col: 25, offset: 2814},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 44, col: 30, offset: 2819},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 44, col: 32, offset: 2821},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 44, col: 35, offset: 2824},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 44, col: 39, offset: 2828},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 44, col: 41, offset: 2830},
									label: "ra",
									expr: &ruleRefExpr{
										pos:  position{line: 44, col: 44, offset: 2833},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 44, col: 48, offset: 2837},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 44, col: 50, offset: 2839},
									val:        "[",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 44, col: 54, offset: 2843},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 44, col: 56, offset: 2845},
									label: "nm",
									expr: &ruleRefExpr{
										pos:  position{line: 44, col: 59, offset: 2848},
										name: "Num",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 44, col: 63, offset: 2852},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 44, col: 65, offset: 2854},
									val:        "]",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 46, col: 15, offset: 2946},
						run: (*parser).callonInstr251,
						expr: &seqExpr{
							pos: position{line: 46, col: 15, offset: 2946},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 46, col: 15, offset: 2946},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 46, col: 18, offset: 2949},
										name: "BOp",
									},
								},
								&labeledExpr{
									pos:   position{line: 46, col: 22, offset: 2953},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 46, col: 25, offset: 2956},
										expr: &ruleRefExpr{
											pos:  position{line: 46, col: 25, offset: 2956},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 46, col: 30, offset: 2961},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 46, col: 32, offset: 2963},
									label: "lb",
									expr: &ruleRefExpr{
										pos:  position{line: 46, col: 35, offset: 2966},
										name: "Label",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name:        "Shift",
			displayName: "\"shift operation\"",
			pos:         position{line: 48, col: 1, offset: 3074},
			expr: &actionExpr{
				pos: position{line: 48, col: 39, offset: 3112},
				run: (*parser).callonShift1,
				expr: &seqExpr{
					pos: position{line: 48, col: 39, offset: 3112},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 48, col: 39, offset: 3112},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 48, col: 41, offset: 3114},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 48, col: 44, offset: 3117},
								name: "ZOp",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 48, col: 48, offset: 3121},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 48, col: 50, offset: 3123},
							label: "nm",
							expr: &ruleRefExpr{
								pos:  position{line: 48, col: 53, offset: 3126},
								name: "Num",
							},
						},
					},
				},
			},
		},
		{
			name:        "DOp",
			displayName: "\"data operator\"",
			pos:         position{line: 50, col: 1, offset: 3216},
			expr: &actionExpr{
				pos: position{line: 50, col: 39, offset: 3254},
				run: (*parser).callonDOp1,
				expr: &choiceExpr{
					pos: position{line: 50, col: 40, offset: 3255},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 50, col: 40, offset: 3255},
							val:        "add",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 50, col: 48, offset: 3263},
							val:        "sub",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 50, col: 56, offset: 3271},
							val:        "mul",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 50, col: 64, offset: 3279},
							val:        "div",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 51, col: 40, offset: 3327},
							val:        "and",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 51, col: 48, offset: 3335},
							val:        "oor",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 51, col: 56, offset: 3343},
							val:        "xor",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 51, col: 64, offset: 3351},
							val:        "nor",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 52, col: 40, offset: 3399},
							val:        "adu",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 52, col: 48, offset: 3407},
							val:        "sbu",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 52, col: 56, offset: 3415},
							val:        "mlu",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 52, col: 64, offset: 3423},
							val:        "dvu",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name:        "COp",
			displayName: "\"comparison operator\"",
			pos:         position{line: 53, col: 1, offset: 3496},
			expr: &actionExpr{
				pos: position{line: 53, col: 39, offset: 3534},
				run: (*parser).callonCOp1,
				expr: &choiceExpr{
					pos: position{line: 53, col: 40, offset: 3535},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 53, col: 40, offset: 3535},
							val:        "cmp",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 53, col: 48, offset: 3543},
							val:        "cpu",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 53, col: 56, offset: 3551},
							val:        "tst",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name:        "MOp",
			displayName: "\"memory operator\"",
			pos:         position{line: 54, col: 1, offset: 3632},
			expr: &actionExpr{
				pos: position{line: 54, col: 39, offset: 3670},
				run: (*parser).callonMOp1,
				expr: &choiceExpr{
					pos: position{line: 54, col: 40, offset: 3671},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 54, col: 40, offset: 3671},
							val:        "ldw",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 54, col: 48, offset: 3679},
							val:        "stw",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name:        "BOp",
			displayName: "\"branch operator\"",
			pos:         position{line: 55, col: 1, offset: 3768},
			expr: &actionExpr{
				pos: position{line: 55, col: 39, offset: 3806},
				run: (*parser).callonBOp1,
				expr: &choiceExpr{
					pos: position{line: 55, col: 40, offset: 3807},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 55, col: 40, offset: 3807},
							val:        "bra",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 55, col: 48, offset: 3815},
							val:        "brl",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name:        "ZOp",
			displayName: "\"shift operator\"",
			pos:         position{line: 56, col: 1, offset: 3904},
			expr: &actionExpr{
				pos: position{line: 56, col: 39, offset: 3942},
				run: (*parser).callonZOp1,
				expr: &choiceExpr{
					pos: position{line: 56, col: 40, offset: 3943},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 56, col: 40, offset: 3943},
							val:        "<<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 56, col: 48, offset: 3951},
							val:        "<<>",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 56, col: 56, offset: 3959},
							val:        "<>>",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 56, col: 64, offset: 3967},
							val:        ">>",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 56, col: 72, offset: 3975},
							val:        ">>>",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name:        "SOp",
			displayName: "\"shift operator\"",
			pos:         position{line: 57, col: 1, offset: 4040},
			expr: &actionExpr{
				pos: position{line: 57, col: 39, offset: 4078},
				run: (*parser).callonSOp1,
				expr: &choiceExpr{
					pos: position{line: 57, col: 40, offset: 4079},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 57, col: 40, offset: 4079},
							val:        "sll",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 57, col: 48, offset: 4087},
							val:        "rol",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 57, col: 56, offset: 4095},
							val:        "srl",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 57, col: 64, offset: 4103},
							val:        "sra",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name:        "Cnd",
			displayName: "\"condition flag\"",
			pos:         position{line: 59, col: 1, offset: 4178},
			expr: &actionExpr{
				pos: position{line: 59, col: 39, offset: 4216},
				run: (*parser).callonCnd1,
				expr: &choiceExpr{
					pos: position{line: 59, col: 40, offset: 4217},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 59, col: 40, offset: 4217},
							val:        "nv",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 59, col: 47, offset: 4224},
							val:        "eq",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 59, col: 54, offset: 4231},
							val:        "lt",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 59, col: 61, offset: 4238},
							val:        "le",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 60, col: 40, offset: 4285},
							val:        "gt",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 60, col: 47, offset: 4292},
							val:        "ge",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 60, col: 54, offset: 4299},
							val:        "ne",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 60, col: 61, offset: 4306},
							val:        "al",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name:        "Reg",
			displayName: "\"register\"",
			pos:         position{line: 62, col: 1, offset: 4384},
			expr: &actionExpr{
				pos: position{line: 62, col: 39, offset: 4422},
				run: (*parser).callonReg1,
				expr: &choiceExpr{
					pos: position{line: 62, col: 40, offset: 4423},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 62, col: 40, offset: 4423},
							val:        "sp",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 62, col: 47, offset: 4430},
							val:        "rp",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 62, col: 54, offset: 4437},
							val:        "ip",
							ignoreCase: false,
						},
						&seqExpr{
							pos: position{line: 62, col: 61, offset: 4444},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 62, col: 61, offset: 4444},
									val:        "r",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 62, col: 65, offset: 4448},
									expr: &charClassMatcher{
										pos:        position{line: 62, col: 65, offset: 4448},
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
			name:        "Num",
			displayName: "\"number\"",
			pos:         position{line: 64, col: 1, offset: 4522},
			expr: &actionExpr{
				pos: position{line: 64, col: 39, offset: 4560},
				run: (*parser).callonNum1,
				expr: &labeledExpr{
					pos:   position{line: 64, col: 39, offset: 4560},
					label: "nm",
					expr: &choiceExpr{
						pos: position{line: 64, col: 43, offset: 4564},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 64, col: 43, offset: 4564},
								name: "BinNum",
							},
							&ruleRefExpr{
								pos:  position{line: 64, col: 52, offset: 4573},
								name: "HexNum",
							},
							&ruleRefExpr{
								pos:  position{line: 64, col: 61, offset: 4582},
								name: "DecNum",
							},
						},
					},
				},
			},
		},
		{
			name:        "BinNum",
			displayName: "\"binary number\"",
			pos:         position{line: 65, col: 1, offset: 4658},
			expr: &actionExpr{
				pos: position{line: 65, col: 39, offset: 4696},
				run: (*parser).callonBinNum1,
				expr: &seqExpr{
					pos: position{line: 65, col: 39, offset: 4696},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 65, col: 39, offset: 4696},
							val:        "0b",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 65, col: 44, offset: 4701},
							expr: &charClassMatcher{
								pos:        position{line: 65, col: 44, offset: 4701},
								val:        "[01]",
								chars:      []rune{'0', '1'},
								ignoreCase: false,
								inverted:   false,
							},
						},
					},
				},
			},
		},
		{
			name:        "HexNum",
			displayName: "\"hexadecimal number\"",
			pos:         position{line: 66, col: 1, offset: 4794},
			expr: &actionExpr{
				pos: position{line: 66, col: 39, offset: 4832},
				run: (*parser).callonHexNum1,
				expr: &seqExpr{
					pos: position{line: 66, col: 39, offset: 4832},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 66, col: 39, offset: 4832},
							val:        "0x",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 66, col: 44, offset: 4837},
							expr: &charClassMatcher{
								pos:        position{line: 66, col: 44, offset: 4837},
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
			name:        "DecNum",
			displayName: "\"decimal number\"",
			pos:         position{line: 67, col: 1, offset: 4930},
			expr: &actionExpr{
				pos: position{line: 67, col: 39, offset: 4968},
				run: (*parser).callonDecNum1,
				expr: &seqExpr{
					pos: position{line: 67, col: 39, offset: 4968},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 67, col: 39, offset: 4968},
							expr: &litMatcher{
								pos:        position{line: 67, col: 39, offset: 4968},
								val:        "-",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 67, col: 44, offset: 4973},
							expr: &charClassMatcher{
								pos:        position{line: 67, col: 44, offset: 4973},
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
			name:        "Label",
			displayName: "\"label\"",
			pos:         position{line: 69, col: 1, offset: 5068},
			expr: &actionExpr{
				pos: position{line: 69, col: 39, offset: 5106},
				run: (*parser).callonLabel1,
				expr: &seqExpr{
					pos: position{line: 69, col: 39, offset: 5106},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 69, col: 39, offset: 5106},
							val:        "@",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 69, col: 43, offset: 5110},
							expr: &charClassMatcher{
								pos:        position{line: 69, col: 43, offset: 5110},
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
			pos:  position{line: 71, col: 1, offset: 5210},
			expr: &actionExpr{
				pos: position{line: 71, col: 39, offset: 5248},
				run: (*parser).callonComment1,
				expr: &seqExpr{
					pos: position{line: 71, col: 39, offset: 5248},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 71, col: 39, offset: 5248},
							val:        "//",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 71, col: 44, offset: 5253},
							expr: &charClassMatcher{
								pos:        position{line: 71, col: 44, offset: 5253},
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
			pos:         position{line: 73, col: 1, offset: 5352},
			expr: &zeroOrMoreExpr{
				pos: position{line: 73, col: 39, offset: 5390},
				expr: &charClassMatcher{
					pos:        position{line: 73, col: 39, offset: 5390},
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

func (c *current) onStart12(lbl interface{}) (interface{}, error) {
	return Forward(lbl)
}

func (p *parser) callonStart12() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStart12(stack["lbl"])
}

func (c *current) onStart20(cmt interface{}) (interface{}, error) {
	return Forward(cmt)
}

func (p *parser) callonStart20() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStart20(stack["cmt"])
}

func (c *current) onInstr2(cd interface{}) (interface{}, error) {
	return gen.NewNopInstr(cd)
}

func (p *parser) callonInstr2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr2(stack["cd"])
}

func (c *current) onInstr8(s, cd, rd interface{}) (interface{}, error) {
	return gen.NewClrInstr(s, cd, rd)
}

func (p *parser) callonInstr8() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr8(stack["s"], stack["cd"], stack["rd"])
}

func (c *current) onInstr21(s, cd, rd, rb, sh interface{}) (interface{}, error) {
	return gen.NewRegInstr(s, "mov", cd, rd, "r0", rb, sh)
}

func (p *parser) callonInstr21() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr21(stack["s"], stack["cd"], stack["rd"], stack["rb"], stack["sh"])
}

func (c *current) onInstr40(s, cd, rd, nm interface{}) (interface{}, error) {
	return gen.NewLdcInstr(s, cd, rd, nm)
}

func (p *parser) callonInstr40() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr40(stack["s"], stack["cd"], stack["rd"], stack["nm"])
}

func (c *current) onInstr56(s, cd, rd, lb interface{}) (interface{}, error) {
	return gen.NewLdaInstr(s, cd, rd, lb)
}

func (p *parser) callonInstr56() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr56(stack["s"], stack["cd"], stack["rd"], stack["lb"])
}

func (c *current) onInstr72(s, op, cd, rd, rb, nm interface{}) (interface{}, error) {
	return gen.NewShiftInstr(s, op, cd, rd, rb, nm)
}

func (p *parser) callonInstr72() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr72(stack["s"], stack["op"], stack["cd"], stack["rd"], stack["rb"], stack["nm"])
}

func (c *current) onInstr92(s, op, cd, rd, ra, rb, sh interface{}) (interface{}, error) {
	return gen.NewRegInstr(s, op, cd, rd, ra, rb, sh)
}

func (p *parser) callonInstr92() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr92(stack["s"], stack["op"], stack["cd"], stack["rd"], stack["ra"], stack["rb"], stack["sh"])
}

func (c *current) onInstr115(s, op, cd, rd, ra, nm interface{}) (interface{}, error) {
	return gen.NewI12Instr(s, op, cd, rd, ra, nm)
}

func (p *parser) callonInstr115() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr115(stack["s"], stack["op"], stack["cd"], stack["rd"], stack["ra"], stack["nm"])
}

func (c *current) onInstr135(s, op, cd, rd, nm, up interface{}) (interface{}, error) {
	return gen.NewI16Instr(s, op, cd, up, rd, nm)
}

func (p *parser) callonInstr135() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr135(stack["s"], stack["op"], stack["cd"], stack["rd"], stack["nm"], stack["up"])
}

func (c *current) onInstr159(s, op, cd, ra, rb, sh interface{}) (interface{}, error) {
	return gen.NewRegInstr(s, op, cd, "r0", ra, rb, sh)
}

func (p *parser) callonInstr159() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr159(stack["s"], stack["op"], stack["cd"], stack["ra"], stack["rb"], stack["sh"])
}

func (c *current) onInstr179(s, op, cd, ra, nm, up interface{}) (interface{}, error) {
	return gen.NewI16Instr(s, op, cd, up, ra, nm)
}

func (p *parser) callonInstr179() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr179(stack["s"], stack["op"], stack["cd"], stack["ra"], stack["nm"], stack["up"])
}

func (c *current) onInstr203(s, op, cd, rd, ra, rb interface{}) (interface{}, error) {
	return gen.NewMemRegInstr(s, op, cd, rd, ra, rb)
}

func (p *parser) callonInstr203() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr203(stack["s"], stack["op"], stack["cd"], stack["rd"], stack["ra"], stack["rb"])
}

func (c *current) onInstr227(s, op, cd, rd, ra, nm interface{}) (interface{}, error) {
	return gen.NewMemI12Instr(s, op, cd, rd, ra, nm)
}

func (p *parser) callonInstr227() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr227(stack["s"], stack["op"], stack["cd"], stack["rd"], stack["ra"], stack["nm"])
}

func (c *current) onInstr251(op, cd, lb interface{}) (interface{}, error) {
	return gen.NewBraInstr(op, cd, lb)
}

func (p *parser) callonInstr251() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr251(stack["op"], stack["cd"], stack["lb"])
}

func (c *current) onShift1(op, nm interface{}) (interface{}, error) {
	return gen.NewNumShift(op, nm)
}

func (p *parser) callonShift1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onShift1(stack["op"], stack["nm"])
}

func (c *current) onDOp1() (interface{}, error) {
	return NewString(c.text)
}

func (p *parser) callonDOp1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDOp1()
}

func (c *current) onCOp1() (interface{}, error) {
	return NewString(c.text)
}

func (p *parser) callonCOp1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCOp1()
}

func (c *current) onMOp1() (interface{}, error) {
	return NewString(c.text)
}

func (p *parser) callonMOp1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMOp1()
}

func (c *current) onBOp1() (interface{}, error) {
	return NewString(c.text)
}

func (p *parser) callonBOp1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBOp1()
}

func (c *current) onZOp1() (interface{}, error) {
	return NewString(c.text)
}

func (p *parser) callonZOp1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onZOp1()
}

func (c *current) onSOp1() (interface{}, error) {
	return NewString(c.text)
}

func (p *parser) callonSOp1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSOp1()
}

func (c *current) onCnd1() (interface{}, error) {
	return NewString(c.text)
}

func (p *parser) callonCnd1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCnd1()
}

func (c *current) onReg1() (interface{}, error) {
	return NewString(c.text)
}

func (p *parser) callonReg1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onReg1()
}

func (c *current) onNum1(nm interface{}) (interface{}, error) {
	return Forward(nm)
}

func (p *parser) callonNum1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onNum1(stack["nm"])
}

func (c *current) onBinNum1() (interface{}, error) {
	return NewString(c.text)
}

func (p *parser) callonBinNum1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBinNum1()
}

func (c *current) onHexNum1() (interface{}, error) {
	return NewString(c.text)
}

func (p *parser) callonHexNum1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onHexNum1()
}

func (c *current) onDecNum1() (interface{}, error) {
	return NewString(c.text)
}

func (p *parser) callonDecNum1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDecNum1()
}

func (c *current) onLabel1() (interface{}, error) {
	return gen.NewLabel(c.text)
}

func (p *parser) callonLabel1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLabel1()
}

func (c *current) onComment1() (interface{}, error) {
	return gen.NewComment()
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
