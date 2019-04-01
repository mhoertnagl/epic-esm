package gengen

import (
	"testing"
  "fmt"
  "strconv"
  "strings"
)

func TestBitPatterns(t *testing.T) {
  seq := &PatternSeq{
    Patterns: []Pattern{
      &BitsPattern{ Val: 8, Len: 4 },
      &BitsPattern{ Val: 7, Len: 4 },
      &BitsPattern{ Val: 6, Len: 4 },
      &BitsPattern{ Val: 5, Len: 4 },
      &BitsPattern{ Val: 4, Len: 4 },
      &BitsPattern{ Val: 3, Len: 4 },
      &BitsPattern{ Val: 2, Len: 4 },
      &BitsPattern{ Val: 1, Len: 4 },
    },
  }
  env := map[string]*BitsPattern{
    
  }

  test(t, seq, env, "1000 0111 0110 0101 0100 0011 0010 0001")
}

func TestVarPatterns(t *testing.T) {
  seq := &PatternSeq{
    Patterns: []Pattern{
      &VarPattern{ Name: "h" },
      &VarPattern{ Name: "g" },
      &VarPattern{ Name: "f" },
      &VarPattern{ Name: "e" },
      &VarPattern{ Name: "d" },
      &VarPattern{ Name: "c" },
      &VarPattern{ Name: "b" },
      &VarPattern{ Name: "a" },
    },
  }
  env := map[string]*BitsPattern{
    "a": &BitsPattern{ Val: 8, Len: 4 },
    "b": &BitsPattern{ Val: 7, Len: 4 },
    "c": &BitsPattern{ Val: 6, Len: 4 },
    "d": &BitsPattern{ Val: 5, Len: 4 },
    "e": &BitsPattern{ Val: 4, Len: 4 },
    "f": &BitsPattern{ Val: 3, Len: 4 },
    "g": &BitsPattern{ Val: 2, Len: 4 },
    "h": &BitsPattern{ Val: 1, Len: 4 },    
  }

  test(t, seq, env, "0001 0010 0011 0100 0101 0110 0111 1000")
}

func test(t *testing.T, seq *PatternSeq, env Env, expected string) {
  ctx := NewAsmContext()
  code := ctx.GeneratePat(seq, env)
  actual := fmt.Sprintf("%032s", strconv.FormatUint(uint64(code), 2))
  condensed := strings.Replace(expected, " ", "", 7)
	if actual != condensed {
		t.Errorf("Expected [%s] but got [%s].", condensed, actual)
	}
}
