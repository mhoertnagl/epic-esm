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
				pos: position{line: 6, col: 6, offset: 31},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 6, col: 6, offset: 31},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 7, col: 8, offset: 42},
						alternatives: []interface{}{
							&actionExpr{
								pos: position{line: 7, col: 8, offset: 42},
								run: (*parser).callonStart4,
								expr: &seqExpr{
									pos: position{line: 7, col: 8, offset: 42},
									exprs: []interface{}{
										&labeledExpr{
											pos:   position{line: 7, col: 8, offset: 42},
											label: "ins",
											expr: &ruleRefExpr{
												pos:  position{line: 7, col: 12, offset: 46},
												name: "Instr",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 7, col: 18, offset: 52},
											name: "_",
										},
										&zeroOrOneExpr{
											pos: position{line: 7, col: 20, offset: 54},
											expr: &ruleRefExpr{
												pos:  position{line: 7, col: 20, offset: 54},
												name: "Comment",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 7, col: 29, offset: 63},
											name: "_",
										},
									},
								},
							},
							&actionExpr{
								pos: position{line: 8, col: 8, offset: 174},
								run: (*parser).callonStart12,
								expr: &seqExpr{
									pos: position{line: 8, col: 8, offset: 174},
									exprs: []interface{}{
										&labeledExpr{
											pos:   position{line: 8, col: 8, offset: 174},
											label: "lbl",
											expr: &ruleRefExpr{
												pos:  position{line: 8, col: 12, offset: 178},
												name: "Label",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 8, col: 18, offset: 184},
											name: "_",
										},
										&zeroOrOneExpr{
											pos: position{line: 8, col: 20, offset: 186},
											expr: &ruleRefExpr{
												pos:  position{line: 8, col: 20, offset: 186},
												name: "Comment",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 8, col: 29, offset: 195},
											name: "_",
										},
									},
								},
							},
							&actionExpr{
								pos: position{line: 9, col: 8, offset: 306},
								run: (*parser).callonStart20,
								expr: &seqExpr{
									pos: position{line: 9, col: 8, offset: 306},
									exprs: []interface{}{
										&labeledExpr{
											pos:   position{line: 9, col: 8, offset: 306},
											label: "cmt",
											expr: &ruleRefExpr{
												pos:  position{line: 9, col: 12, offset: 310},
												name: "Comment",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 9, col: 20, offset: 318},
											name: "_",
										},
									},
								},
							},
						},
					},
					&notExpr{
						pos: position{line: 10, col: 6, offset: 436},
						expr: &anyMatcher{
							line: 10, col: 7, offset: 437,
						},
					},
				},
			},
		},
		{
			name:        "Instr",
			displayName: "\"instruction\"",
			pos:         position{line: 12, col: 1, offset: 440},
			expr: &choiceExpr{
				pos: position{line: 13, col: 15, offset: 476},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 13, col: 15, offset: 476},
						run: (*parser).callonInstr2,
						expr: &seqExpr{
							pos: position{line: 13, col: 15, offset: 476},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 13, col: 15, offset: 476},
									val:        "nop",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 13, col: 22, offset: 483},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 13, col: 25, offset: 486},
										expr: &ruleRefExpr{
											pos:  position{line: 13, col: 25, offset: 486},
											name: "Cnd",
										},
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 15, col: 6, offset: 600},
						run: (*parser).callonInstr8,
						expr: &seqExpr{
							pos: position{line: 15, col: 6, offset: 600},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 15, col: 6, offset: 600},
									label: "s",
									expr: &zeroOrOneExpr{
										pos: position{line: 15, col: 8, offset: 602},
										expr: &litMatcher{
											pos:        position{line: 15, col: 8, offset: 602},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 15, col: 13, offset: 607},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 15, col: 15, offset: 609},
									val:        "clr",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 15, col: 22, offset: 616},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 15, col: 25, offset: 619},
										expr: &ruleRefExpr{
											pos:  position{line: 15, col: 25, offset: 619},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 15, col: 30, offset: 624},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 15, col: 32, offset: 626},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 15, col: 35, offset: 629},
										name: "Reg",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 19, col: 6, offset: 865},
						run: (*parser).callonInstr21,
						expr: &seqExpr{
							pos: position{line: 19, col: 6, offset: 865},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 19, col: 6, offset: 865},
									label: "s",
									expr: &zeroOrOneExpr{
										pos: position{line: 19, col: 8, offset: 867},
										expr: &litMatcher{
											pos:        position{line: 19, col: 8, offset: 867},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 19, col: 13, offset: 872},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 19, col: 15, offset: 874},
									label: "op",
									expr: &litMatcher{
										pos:        position{line: 19, col: 18, offset: 877},
										val:        "mov",
										ignoreCase: false,
									},
								},
								&labeledExpr{
									pos:   position{line: 19, col: 25, offset: 884},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 19, col: 28, offset: 887},
										expr: &ruleRefExpr{
											pos:  position{line: 19, col: 28, offset: 887},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 19, col: 33, offset: 892},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 19, col: 35, offset: 894},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 19, col: 38, offset: 897},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 19, col: 42, offset: 901},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 19, col: 44, offset: 903},
									label: "rb",
									expr: &ruleRefExpr{
										pos:  position{line: 19, col: 47, offset: 906},
										name: "Reg",
									},
								},
								&labeledExpr{
									pos:   position{line: 19, col: 51, offset: 910},
									label: "sh",
									expr: &zeroOrOneExpr{
										pos: position{line: 19, col: 54, offset: 913},
										expr: &ruleRefExpr{
											pos:  position{line: 19, col: 54, offset: 913},
											name: "Shift",
										},
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 20, col: 6, offset: 997},
						run: (*parser).callonInstr41,
						expr: &seqExpr{
							pos: position{line: 20, col: 6, offset: 997},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 20, col: 6, offset: 997},
									label: "s",
									expr: &zeroOrOneExpr{
										pos: position{line: 20, col: 8, offset: 999},
										expr: &litMatcher{
											pos:        position{line: 20, col: 8, offset: 999},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 20, col: 13, offset: 1004},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 20, col: 15, offset: 1006},
									label: "op",
									expr: &litMatcher{
										pos:        position{line: 20, col: 18, offset: 1009},
										val:        "ldc",
										ignoreCase: false,
									},
								},
								&labeledExpr{
									pos:   position{line: 20, col: 25, offset: 1016},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 20, col: 28, offset: 1019},
										expr: &ruleRefExpr{
											pos:  position{line: 20, col: 28, offset: 1019},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 20, col: 33, offset: 1024},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 20, col: 35, offset: 1026},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 20, col: 38, offset: 1029},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 20, col: 42, offset: 1033},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 20, col: 44, offset: 1035},
									label: "nm",
									expr: &ruleRefExpr{
										pos:  position{line: 20, col: 47, offset: 1038},
										name: "Num",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 21, col: 6, offset: 1129},
						run: (*parser).callonInstr58,
						expr: &seqExpr{
							pos: position{line: 21, col: 6, offset: 1129},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 21, col: 6, offset: 1129},
									label: "s",
									expr: &zeroOrOneExpr{
										pos: position{line: 21, col: 8, offset: 1131},
										expr: &litMatcher{
											pos:        position{line: 21, col: 8, offset: 1131},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 21, col: 13, offset: 1136},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 21, col: 15, offset: 1138},
									label: "op",
									expr: &litMatcher{
										pos:        position{line: 21, col: 18, offset: 1141},
										val:        "lda",
										ignoreCase: false,
									},
								},
								&labeledExpr{
									pos:   position{line: 21, col: 25, offset: 1148},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 21, col: 28, offset: 1151},
										expr: &ruleRefExpr{
											pos:  position{line: 21, col: 28, offset: 1151},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 21, col: 33, offset: 1156},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 21, col: 35, offset: 1158},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 21, col: 38, offset: 1161},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 21, col: 42, offset: 1165},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 21, col: 44, offset: 1167},
									label: "lb",
									expr: &ruleRefExpr{
										pos:  position{line: 21, col: 47, offset: 1170},
										name: "Label",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 28, col: 6, offset: 1663},
						run: (*parser).callonInstr75,
						expr: &seqExpr{
							pos: position{line: 28, col: 6, offset: 1663},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 28, col: 6, offset: 1663},
									label: "s",
									expr: &zeroOrOneExpr{
										pos: position{line: 28, col: 8, offset: 1665},
										expr: &litMatcher{
											pos:        position{line: 28, col: 8, offset: 1665},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 28, col: 13, offset: 1670},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 28, col: 15, offset: 1672},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 28, col: 18, offset: 1675},
										name: "SOp",
									},
								},
								&labeledExpr{
									pos:   position{line: 28, col: 22, offset: 1679},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 28, col: 25, offset: 1682},
										expr: &ruleRefExpr{
											pos:  position{line: 28, col: 25, offset: 1682},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 28, col: 30, offset: 1687},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 28, col: 32, offset: 1689},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 28, col: 35, offset: 1692},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 28, col: 39, offset: 1696},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 28, col: 41, offset: 1698},
									label: "rb",
									expr: &ruleRefExpr{
										pos:  position{line: 28, col: 44, offset: 1701},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 28, col: 48, offset: 1705},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 28, col: 50, offset: 1707},
									label: "nm",
									expr: &ruleRefExpr{
										pos:  position{line: 28, col: 53, offset: 1710},
										name: "Num",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 34, col: 6, offset: 1821},
						run: (*parser).callonInstr95,
						expr: &seqExpr{
							pos: position{line: 34, col: 6, offset: 1821},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 34, col: 6, offset: 1821},
									label: "s",
									expr: &zeroOrOneExpr{
										pos: position{line: 34, col: 8, offset: 1823},
										expr: &litMatcher{
											pos:        position{line: 34, col: 8, offset: 1823},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 34, col: 13, offset: 1828},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 34, col: 15, offset: 1830},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 34, col: 18, offset: 1833},
										name: "DOp",
									},
								},
								&labeledExpr{
									pos:   position{line: 34, col: 22, offset: 1837},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 34, col: 25, offset: 1840},
										expr: &ruleRefExpr{
											pos:  position{line: 34, col: 25, offset: 1840},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 34, col: 30, offset: 1845},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 34, col: 32, offset: 1847},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 34, col: 35, offset: 1850},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 34, col: 39, offset: 1854},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 34, col: 41, offset: 1856},
									label: "ra",
									expr: &ruleRefExpr{
										pos:  position{line: 34, col: 44, offset: 1859},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 34, col: 48, offset: 1863},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 34, col: 50, offset: 1865},
									label: "rb",
									expr: &ruleRefExpr{
										pos:  position{line: 34, col: 53, offset: 1868},
										name: "Reg",
									},
								},
								&labeledExpr{
									pos:   position{line: 34, col: 57, offset: 1872},
									label: "sh",
									expr: &zeroOrOneExpr{
										pos: position{line: 34, col: 60, offset: 1875},
										expr: &ruleRefExpr{
											pos:  position{line: 34, col: 60, offset: 1875},
											name: "Shift",
										},
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 35, col: 6, offset: 1953},
						run: (*parser).callonInstr118,
						expr: &seqExpr{
							pos: position{line: 35, col: 6, offset: 1953},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 35, col: 6, offset: 1953},
									label: "s",
									expr: &zeroOrOneExpr{
										pos: position{line: 35, col: 8, offset: 1955},
										expr: &litMatcher{
											pos:        position{line: 35, col: 8, offset: 1955},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 35, col: 13, offset: 1960},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 35, col: 15, offset: 1962},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 35, col: 18, offset: 1965},
										name: "DOp",
									},
								},
								&labeledExpr{
									pos:   position{line: 35, col: 22, offset: 1969},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 35, col: 25, offset: 1972},
										expr: &ruleRefExpr{
											pos:  position{line: 35, col: 25, offset: 1972},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 35, col: 30, offset: 1977},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 35, col: 32, offset: 1979},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 35, col: 35, offset: 1982},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 35, col: 39, offset: 1986},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 35, col: 41, offset: 1988},
									label: "ra",
									expr: &ruleRefExpr{
										pos:  position{line: 35, col: 44, offset: 1991},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 35, col: 48, offset: 1995},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 35, col: 50, offset: 1997},
									label: "nm",
									expr: &ruleRefExpr{
										pos:  position{line: 35, col: 53, offset: 2000},
										name: "Num",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 36, col: 6, offset: 2085},
						run: (*parser).callonInstr138,
						expr: &seqExpr{
							pos: position{line: 36, col: 6, offset: 2085},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 36, col: 6, offset: 2085},
									label: "s",
									expr: &zeroOrOneExpr{
										pos: position{line: 36, col: 8, offset: 2087},
										expr: &litMatcher{
											pos:        position{line: 36, col: 8, offset: 2087},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 36, col: 13, offset: 2092},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 36, col: 15, offset: 2094},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 36, col: 18, offset: 2097},
										name: "DOp",
									},
								},
								&labeledExpr{
									pos:   position{line: 36, col: 22, offset: 2101},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 36, col: 25, offset: 2104},
										expr: &ruleRefExpr{
											pos:  position{line: 36, col: 25, offset: 2104},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 36, col: 30, offset: 2109},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 36, col: 32, offset: 2111},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 36, col: 35, offset: 2114},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 36, col: 39, offset: 2118},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 36, col: 41, offset: 2120},
									label: "nm",
									expr: &ruleRefExpr{
										pos:  position{line: 36, col: 44, offset: 2123},
										name: "Num",
									},
								},
								&labeledExpr{
									pos:   position{line: 36, col: 48, offset: 2127},
									label: "up",
									expr: &zeroOrOneExpr{
										pos: position{line: 36, col: 51, offset: 2130},
										expr: &seqExpr{
											pos: position{line: 36, col: 52, offset: 2131},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 36, col: 52, offset: 2131},
													name: "_",
												},
												&litMatcher{
													pos:        position{line: 36, col: 54, offset: 2133},
													val:        "<<",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 36, col: 59, offset: 2138},
													name: "_",
												},
												&litMatcher{
													pos:        position{line: 36, col: 61, offset: 2140},
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
						pos: position{line: 38, col: 6, offset: 2218},
						run: (*parser).callonInstr162,
						expr: &seqExpr{
							pos: position{line: 38, col: 6, offset: 2218},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 38, col: 6, offset: 2218},
									label: "s",
									expr: &zeroOrOneExpr{
										pos: position{line: 38, col: 8, offset: 2220},
										expr: &litMatcher{
											pos:        position{line: 38, col: 8, offset: 2220},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 38, col: 13, offset: 2225},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 38, col: 15, offset: 2227},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 38, col: 18, offset: 2230},
										name: "COp",
									},
								},
								&labeledExpr{
									pos:   position{line: 38, col: 22, offset: 2234},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 38, col: 25, offset: 2237},
										expr: &ruleRefExpr{
											pos:  position{line: 38, col: 25, offset: 2237},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 38, col: 30, offset: 2242},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 38, col: 32, offset: 2244},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 38, col: 35, offset: 2247},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 38, col: 39, offset: 2251},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 38, col: 41, offset: 2253},
									label: "rb",
									expr: &ruleRefExpr{
										pos:  position{line: 38, col: 44, offset: 2256},
										name: "Reg",
									},
								},
								&labeledExpr{
									pos:   position{line: 38, col: 48, offset: 2260},
									label: "sh",
									expr: &zeroOrOneExpr{
										pos: position{line: 38, col: 51, offset: 2263},
										expr: &ruleRefExpr{
											pos:  position{line: 38, col: 51, offset: 2263},
											name: "Shift",
										},
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 39, col: 6, offset: 2350},
						run: (*parser).callonInstr182,
						expr: &seqExpr{
							pos: position{line: 39, col: 6, offset: 2350},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 39, col: 6, offset: 2350},
									label: "s",
									expr: &zeroOrOneExpr{
										pos: position{line: 39, col: 8, offset: 2352},
										expr: &litMatcher{
											pos:        position{line: 39, col: 8, offset: 2352},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 39, col: 13, offset: 2357},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 39, col: 15, offset: 2359},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 39, col: 18, offset: 2362},
										name: "COp",
									},
								},
								&labeledExpr{
									pos:   position{line: 39, col: 22, offset: 2366},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 39, col: 25, offset: 2369},
										expr: &ruleRefExpr{
											pos:  position{line: 39, col: 25, offset: 2369},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 39, col: 30, offset: 2374},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 39, col: 32, offset: 2376},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 39, col: 35, offset: 2379},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 39, col: 39, offset: 2383},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 39, col: 41, offset: 2385},
									label: "nm",
									expr: &ruleRefExpr{
										pos:  position{line: 39, col: 44, offset: 2388},
										name: "Num",
									},
								},
								&labeledExpr{
									pos:   position{line: 39, col: 48, offset: 2392},
									label: "up",
									expr: &zeroOrOneExpr{
										pos: position{line: 39, col: 51, offset: 2395},
										expr: &seqExpr{
											pos: position{line: 39, col: 52, offset: 2396},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 39, col: 52, offset: 2396},
													name: "_",
												},
												&litMatcher{
													pos:        position{line: 39, col: 54, offset: 2398},
													val:        "<<",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 39, col: 59, offset: 2403},
													name: "_",
												},
												&litMatcher{
													pos:        position{line: 39, col: 61, offset: 2405},
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
						pos: position{line: 41, col: 6, offset: 2483},
						run: (*parser).callonInstr206,
						expr: &seqExpr{
							pos: position{line: 41, col: 6, offset: 2483},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 41, col: 6, offset: 2483},
									label: "s",
									expr: &zeroOrOneExpr{
										pos: position{line: 41, col: 8, offset: 2485},
										expr: &litMatcher{
											pos:        position{line: 41, col: 8, offset: 2485},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 13, offset: 2490},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 41, col: 15, offset: 2492},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 41, col: 18, offset: 2495},
										name: "MOp",
									},
								},
								&labeledExpr{
									pos:   position{line: 41, col: 22, offset: 2499},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 41, col: 25, offset: 2502},
										expr: &ruleRefExpr{
											pos:  position{line: 41, col: 25, offset: 2502},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 30, offset: 2507},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 41, col: 32, offset: 2509},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 41, col: 35, offset: 2512},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 39, offset: 2516},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 41, col: 41, offset: 2518},
									label: "ra",
									expr: &ruleRefExpr{
										pos:  position{line: 41, col: 44, offset: 2521},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 48, offset: 2525},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 41, col: 50, offset: 2527},
									val:        "[",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 54, offset: 2531},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 41, col: 56, offset: 2533},
									label: "rb",
									expr: &ruleRefExpr{
										pos:  position{line: 41, col: 59, offset: 2536},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 63, offset: 2540},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 41, col: 65, offset: 2542},
									val:        "]",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 42, col: 6, offset: 2615},
						run: (*parser).callonInstr230,
						expr: &seqExpr{
							pos: position{line: 42, col: 6, offset: 2615},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 42, col: 6, offset: 2615},
									label: "s",
									expr: &zeroOrOneExpr{
										pos: position{line: 42, col: 8, offset: 2617},
										expr: &litMatcher{
											pos:        position{line: 42, col: 8, offset: 2617},
											val:        "!",
											ignoreCase: false,
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 42, col: 13, offset: 2622},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 42, col: 15, offset: 2624},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 42, col: 18, offset: 2627},
										name: "MOp",
									},
								},
								&labeledExpr{
									pos:   position{line: 42, col: 22, offset: 2631},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 42, col: 25, offset: 2634},
										expr: &ruleRefExpr{
											pos:  position{line: 42, col: 25, offset: 2634},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 42, col: 30, offset: 2639},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 42, col: 32, offset: 2641},
									label: "rd",
									expr: &ruleRefExpr{
										pos:  position{line: 42, col: 35, offset: 2644},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 42, col: 39, offset: 2648},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 42, col: 41, offset: 2650},
									label: "ra",
									expr: &ruleRefExpr{
										pos:  position{line: 42, col: 44, offset: 2653},
										name: "Reg",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 42, col: 48, offset: 2657},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 42, col: 50, offset: 2659},
									val:        "[",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 42, col: 54, offset: 2663},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 42, col: 56, offset: 2665},
									label: "nm",
									expr: &ruleRefExpr{
										pos:  position{line: 42, col: 59, offset: 2668},
										name: "Num",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 42, col: 63, offset: 2672},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 42, col: 65, offset: 2674},
									val:        "]",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 44, col: 15, offset: 2757},
						run: (*parser).callonInstr254,
						expr: &seqExpr{
							pos: position{line: 44, col: 15, offset: 2757},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 44, col: 15, offset: 2757},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 44, col: 18, offset: 2760},
										name: "BOp",
									},
								},
								&labeledExpr{
									pos:   position{line: 44, col: 22, offset: 2764},
									label: "cd",
									expr: &zeroOrOneExpr{
										pos: position{line: 44, col: 25, offset: 2767},
										expr: &ruleRefExpr{
											pos:  position{line: 44, col: 25, offset: 2767},
											name: "Cnd",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 44, col: 30, offset: 2772},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 44, col: 32, offset: 2774},
									label: "lb",
									expr: &ruleRefExpr{
										pos:  position{line: 44, col: 35, offset: 2777},
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
			pos:         position{line: 46, col: 1, offset: 2876},
			expr: &actionExpr{
				pos: position{line: 46, col: 39, offset: 2914},
				run: (*parser).callonShift1,
				expr: &seqExpr{
					pos: position{line: 46, col: 39, offset: 2914},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 46, col: 39, offset: 2914},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 41, offset: 2916},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 44, offset: 2919},
								name: "ZOp",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 48, offset: 2923},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 50, offset: 2925},
							label: "nm",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 53, offset: 2928},
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
			pos:         position{line: 48, col: 1, offset: 3014},
			expr: &actionExpr{
				pos: position{line: 48, col: 39, offset: 3052},
				run: (*parser).callonDOp1,
				expr: &choiceExpr{
					pos: position{line: 48, col: 40, offset: 3053},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 48, col: 40, offset: 3053},
							val:        "add",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 48, col: 48, offset: 3061},
							val:        "sub",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 48, col: 56, offset: 3069},
							val:        "mul",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 48, col: 64, offset: 3077},
							val:        "div",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 49, col: 40, offset: 3125},
							val:        "and",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 49, col: 48, offset: 3133},
							val:        "oor",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 49, col: 56, offset: 3141},
							val:        "xor",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 49, col: 64, offset: 3149},
							val:        "nor",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 50, col: 40, offset: 3196},
							val:        "adu",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 50, col: 48, offset: 3204},
							val:        "sbu",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 50, col: 56, offset: 3212},
							val:        "mlu",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 50, col: 64, offset: 3220},
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
			pos:         position{line: 51, col: 1, offset: 3289},
			expr: &actionExpr{
				pos: position{line: 51, col: 39, offset: 3327},
				run: (*parser).callonCOp1,
				expr: &choiceExpr{
					pos: position{line: 51, col: 40, offset: 3328},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 51, col: 40, offset: 3328},
							val:        "cmp",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 51, col: 48, offset: 3336},
							val:        "cpu",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 51, col: 56, offset: 3344},
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
			pos:         position{line: 52, col: 1, offset: 3421},
			expr: &actionExpr{
				pos: position{line: 52, col: 39, offset: 3459},
				run: (*parser).callonMOp1,
				expr: &choiceExpr{
					pos: position{line: 52, col: 40, offset: 3460},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 52, col: 40, offset: 3460},
							val:        "ldw",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 52, col: 48, offset: 3468},
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
			pos:         position{line: 53, col: 1, offset: 3553},
			expr: &actionExpr{
				pos: position{line: 53, col: 39, offset: 3591},
				run: (*parser).callonBOp1,
				expr: &choiceExpr{
					pos: position{line: 53, col: 40, offset: 3592},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 53, col: 40, offset: 3592},
							val:        "bra",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 53, col: 48, offset: 3600},
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
			pos:         position{line: 55, col: 1, offset: 3684},
			expr: &actionExpr{
				pos: position{line: 55, col: 39, offset: 3722},
				run: (*parser).callonZOp1,
				expr: &choiceExpr{
					pos: position{line: 55, col: 40, offset: 3723},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 55, col: 40, offset: 3723},
							val:        "<<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 55, col: 48, offset: 3731},
							val:        "<<>",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 55, col: 56, offset: 3739},
							val:        "<>>",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 55, col: 64, offset: 3747},
							val:        ">>",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 55, col: 72, offset: 3755},
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
			pos:         position{line: 56, col: 1, offset: 3816},
			expr: &actionExpr{
				pos: position{line: 56, col: 39, offset: 3854},
				run: (*parser).callonSOp1,
				expr: &choiceExpr{
					pos: position{line: 56, col: 40, offset: 3855},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 56, col: 40, offset: 3855},
							val:        "sll",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 56, col: 48, offset: 3863},
							val:        "rol",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 56, col: 56, offset: 3871},
							val:        "srl",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 56, col: 64, offset: 3879},
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
			pos:         position{line: 58, col: 1, offset: 3949},
			expr: &actionExpr{
				pos: position{line: 58, col: 39, offset: 3987},
				run: (*parser).callonCnd1,
				expr: &choiceExpr{
					pos: position{line: 58, col: 40, offset: 3988},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 58, col: 40, offset: 3988},
							val:        "nv",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 58, col: 47, offset: 3995},
							val:        "eq",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 58, col: 54, offset: 4002},
							val:        "lt",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 58, col: 61, offset: 4009},
							val:        "le",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 59, col: 40, offset: 4056},
							val:        "gt",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 59, col: 47, offset: 4063},
							val:        "ge",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 59, col: 54, offset: 4070},
							val:        "ne",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 59, col: 61, offset: 4077},
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
			pos:         position{line: 61, col: 1, offset: 4150},
			expr: &actionExpr{
				pos: position{line: 61, col: 39, offset: 4188},
				run: (*parser).callonReg1,
				expr: &choiceExpr{
					pos: position{line: 61, col: 40, offset: 4189},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 61, col: 40, offset: 4189},
							val:        "sp",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 61, col: 47, offset: 4196},
							val:        "rp",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 61, col: 54, offset: 4203},
							val:        "ip",
							ignoreCase: false,
						},
						&seqExpr{
							pos: position{line: 61, col: 61, offset: 4210},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 61, col: 61, offset: 4210},
									val:        "r",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 61, col: 65, offset: 4214},
									expr: &charClassMatcher{
										pos:        position{line: 61, col: 65, offset: 4214},
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
			pos:         position{line: 63, col: 1, offset: 4283},
			expr: &actionExpr{
				pos: position{line: 63, col: 39, offset: 4321},
				run: (*parser).callonNum1,
				expr: &labeledExpr{
					pos:   position{line: 63, col: 39, offset: 4321},
					label: "nm",
					expr: &choiceExpr{
						pos: position{line: 63, col: 43, offset: 4325},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 63, col: 43, offset: 4325},
								name: "BinNum",
							},
							&ruleRefExpr{
								pos:  position{line: 63, col: 52, offset: 4334},
								name: "HexNum",
							},
							&ruleRefExpr{
								pos:  position{line: 63, col: 61, offset: 4343},
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
			pos:         position{line: 64, col: 1, offset: 4415},
			expr: &actionExpr{
				pos: position{line: 64, col: 39, offset: 4453},
				run: (*parser).callonBinNum1,
				expr: &seqExpr{
					pos: position{line: 64, col: 39, offset: 4453},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 64, col: 39, offset: 4453},
							val:        "0b",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 64, col: 44, offset: 4458},
							expr: &charClassMatcher{
								pos:        position{line: 64, col: 44, offset: 4458},
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
			pos:         position{line: 65, col: 1, offset: 4547},
			expr: &actionExpr{
				pos: position{line: 65, col: 39, offset: 4585},
				run: (*parser).callonHexNum1,
				expr: &seqExpr{
					pos: position{line: 65, col: 39, offset: 4585},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 65, col: 39, offset: 4585},
							val:        "0x",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 65, col: 44, offset: 4590},
							expr: &charClassMatcher{
								pos:        position{line: 65, col: 44, offset: 4590},
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
			pos:         position{line: 66, col: 1, offset: 4679},
			expr: &actionExpr{
				pos: position{line: 66, col: 39, offset: 4717},
				run: (*parser).callonDecNum1,
				expr: &seqExpr{
					pos: position{line: 66, col: 39, offset: 4717},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 66, col: 39, offset: 4717},
							expr: &litMatcher{
								pos:        position{line: 66, col: 39, offset: 4717},
								val:        "-",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 66, col: 44, offset: 4722},
							expr: &charClassMatcher{
								pos:        position{line: 66, col: 44, offset: 4722},
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
			pos:         position{line: 68, col: 1, offset: 4812},
			expr: &actionExpr{
				pos: position{line: 68, col: 39, offset: 4850},
				run: (*parser).callonLabel1,
				expr: &seqExpr{
					pos: position{line: 68, col: 39, offset: 4850},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 68, col: 39, offset: 4850},
							val:        "@",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 68, col: 43, offset: 4854},
							expr: &charClassMatcher{
								pos:        position{line: 68, col: 43, offset: 4854},
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
			pos:  position{line: 70, col: 1, offset: 4945},
			expr: &actionExpr{
				pos: position{line: 70, col: 39, offset: 4983},
				run: (*parser).callonComment1,
				expr: &seqExpr{
					pos: position{line: 70, col: 39, offset: 4983},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 70, col: 39, offset: 4983},
							val:        "//",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 70, col: 44, offset: 4988},
							expr: &charClassMatcher{
								pos:        position{line: 70, col: 44, offset: 4988},
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
			pos:         position{line: 72, col: 1, offset: 5078},
			expr: &zeroOrMoreExpr{
				pos: position{line: 72, col: 39, offset: 5116},
				expr: &charClassMatcher{
					pos:        position{line: 72, col: 39, offset: 5116},
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
	return NewNopInstr(cd)
}

func (p *parser) callonInstr2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr2(stack["cd"])
}

func (c *current) onInstr8(s, cd, rd interface{}) (interface{}, error) {
	return NewClrInstr(s, cd, rd)
}

func (p *parser) callonInstr8() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr8(stack["s"], stack["cd"], stack["rd"])
}

func (c *current) onInstr21(s, op, cd, rd, rb, sh interface{}) (interface{}, error) {
	return NewRegInstr(s, op, cd, rd, "r0", rb, sh)
}

func (p *parser) callonInstr21() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr21(stack["s"], stack["op"], stack["cd"], stack["rd"], stack["rb"], stack["sh"])
}

func (c *current) onInstr41(s, op, cd, rd, nm interface{}) (interface{}, error) {
	return NewLdcInstr(s, cd, rd, nm)
}

func (p *parser) callonInstr41() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr41(stack["s"], stack["op"], stack["cd"], stack["rd"], stack["nm"])
}

func (c *current) onInstr58(s, op, cd, rd, lb interface{}) (interface{}, error) {
	return NewLdaInstr(s, cd, rd, lb)
}

func (p *parser) callonInstr58() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr58(stack["s"], stack["op"], stack["cd"], stack["rd"], stack["lb"])
}

func (c *current) onInstr75(s, op, cd, rd, rb, nm interface{}) (interface{}, error) {
	return NewShiftInstr(s, op, cd, rd, rb, nm)
}

func (p *parser) callonInstr75() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr75(stack["s"], stack["op"], stack["cd"], stack["rd"], stack["rb"], stack["nm"])
}

func (c *current) onInstr95(s, op, cd, rd, ra, rb, sh interface{}) (interface{}, error) {
	return NewRegInstr(s, op, cd, rd, ra, rb, sh)
}

func (p *parser) callonInstr95() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr95(stack["s"], stack["op"], stack["cd"], stack["rd"], stack["ra"], stack["rb"], stack["sh"])
}

func (c *current) onInstr118(s, op, cd, rd, ra, nm interface{}) (interface{}, error) {
	return NewI12Instr(s, op, cd, rd, ra, nm)
}

func (p *parser) callonInstr118() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr118(stack["s"], stack["op"], stack["cd"], stack["rd"], stack["ra"], stack["nm"])
}

func (c *current) onInstr138(s, op, cd, rd, nm, up interface{}) (interface{}, error) {
	return NewI16Instr(s, op, cd, up, rd, nm)
}

func (p *parser) callonInstr138() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr138(stack["s"], stack["op"], stack["cd"], stack["rd"], stack["nm"], stack["up"])
}

func (c *current) onInstr162(s, op, cd, rd, rb, sh interface{}) (interface{}, error) {
	return NewRegInstr(s, op, cd, rd, "r0", rb, sh)
}

func (p *parser) callonInstr162() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr162(stack["s"], stack["op"], stack["cd"], stack["rd"], stack["rb"], stack["sh"])
}

func (c *current) onInstr182(s, op, cd, rd, nm, up interface{}) (interface{}, error) {
	return NewI16Instr(s, op, cd, up, rd, nm)
}

func (p *parser) callonInstr182() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr182(stack["s"], stack["op"], stack["cd"], stack["rd"], stack["nm"], stack["up"])
}

func (c *current) onInstr206(s, op, cd, rd, ra, rb interface{}) (interface{}, error) {
	return NewMemRegInstr(s, op, cd, rd, ra, rb)
}

func (p *parser) callonInstr206() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr206(stack["s"], stack["op"], stack["cd"], stack["rd"], stack["ra"], stack["rb"])
}

func (c *current) onInstr230(s, op, cd, rd, ra, nm interface{}) (interface{}, error) {
	return NewMemI12Instr(s, op, cd, rd, ra, nm)
}

func (p *parser) callonInstr230() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr230(stack["s"], stack["op"], stack["cd"], stack["rd"], stack["ra"], stack["nm"])
}

func (c *current) onInstr254(op, cd, lb interface{}) (interface{}, error) {
	return NewBraInstr(op, cd, lb)
}

func (p *parser) callonInstr254() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInstr254(stack["op"], stack["cd"], stack["lb"])
}

func (c *current) onShift1(op, nm interface{}) (interface{}, error) {
	return NewNumShift(op, nm)
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
