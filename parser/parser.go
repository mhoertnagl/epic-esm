package parser

import (
	"fmt"

	"github.com/mhoertnagl/epic-esm/lexer"
	"github.com/mhoertnagl/epic-esm/token"
)

type Parser struct {
	lexer          *lexer.Lexer
	curToken       token.Token
	nxtToken       token.Token
	errors         []string
}

func NewParser(lexer *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:  lexer,
		errors: []string{},
	}

	// Sets the parsers current and next tokens.
	p.next()
	p.next()
	return p
}

func (p *Parser) debug(prefix string) {
  fmt.Printf("%s: Current: %s | Next: %s\n", prefix, p.curToken, p.nxtToken)
}

func (p *Parser) next() {
	p.curToken = p.nxtToken
	p.nxtToken = p.lexer.Next()
}

func (p *Parser) curTokenIs(exp token.TokenType) bool {
	return p.curToken.Typ == exp
}

func (p *Parser) nxtTokenIs(exp token.TokenType) bool {
	return p.nxtToken.Typ == exp
}

func (p *Parser) expectNext(exp token.TokenType) bool {
	if p.nxtTokenIs(exp) {
		p.next()
		return true
	}
	p.errorNext(exp)
	return false
}

func (p *Parser) HasErrors() bool {
	return len(p.errors) > 0
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) error(format string, a ...interface{}) {
	p.errors = append(p.errors, fmt.Sprintf(format, a...))
}

func (p *Parser) errorNext(exp token.TokenType) {
	p.error("Expected token [%s] but got [%s].", exp, p.nxtToken.Typ)
}

func (p *Parser) Parse() Node {
  switch p.curToken.Typ {
  case token.SET:
  case token.ID:
    return p.parseInstruction()
  }
	return nil
}

func (p *Parser) parseInstruction() Node {
  ins := &Instr{}
  if p.curTokenIs(token.SET) {
    ins.Set = true
    p.next()
  }
  if p.curTokenIs(token.ID) {
    c := p.curToken.Literal
    if len(c) == 5 {
      ins.Cmd = c[:3]
      ins.Cond = c[3:]
    } else {
      ins.Cmd = c
      ins.Cond = "al"      
    }
    p.next()
  } else {
    // Error: expecting assembler command.
    return nil
  }
  for !p.curTokenIs(token.EOF) {
    ins.Args = append(ins.Args, p.curToken)
    p.next()
  }
  return ins
}

// func (p *Parser) parseLetStatement() *LetStatement {
// 	stmt := &LetStatement{Token: p.curToken}
// 	if !p.expectNext(token.ID) {
// 		return nil
// 	}
// 	stmt.Name = &Identifier{Token: p.curToken, Value: p.curToken.Literal}
// 	if !p.expectNext(token.ASSIGN) {
// 		return nil
// 	}
// 	p.next() // Consume [=].
// 	stmt.Value = p.parseExpression(LOWEST)
// 	if p.nxtTokenIs(token.SCOLON) {
// 		p.next()
// 	}
// 	return stmt
// }
// 
// func (p *Parser) parseReturnStatement() *ReturnStatement {
// 	stmt := &ReturnStatement{Token: p.curToken}
// 	p.next() // Consume [return].
// 	stmt.Value = p.parseExpression(LOWEST)
// 	if p.nxtTokenIs(token.SCOLON) {
// 		p.next()
// 	}
// 	return stmt
// }
// 
// func (p *Parser) parseIfStatement() *IfStatement {
// 	stmt := &IfStatement{Token: p.curToken}
//   if !p.expectNext(token.LPAR) {
//     return nil
//   }
//   p.next() // Consume [(].
//   stmt.Condition = p.parseExpression(LOWEST)
//   if !p.expectNext(token.RPAR) {
//     return nil
//   }  
//   p.next() // Consume [)].
//   stmt.Consequence = p.parseStatement()
//   p.next() // Consume [;|}].
//   //p.debug("cons after")
//   if p.curTokenIs(token.ELSE) {    
//     p.next() // Consume [else].    
//     //p.debug("alt before")
//     stmt.Alternative = p.parseStatement()
//     p.next() // Consume [;|}].
//     //p.debug("alt after")
//   }
// 	return stmt
// }
// 
// func (p *Parser) parseBlockStatement() *BlockStatement {
// 	block := &BlockStatement{Token: p.curToken}
//   block.Statements = []Statement{}
// 	p.next() // Consume [{].
//   i := 10
//   //p.debug("block init")
//   for !p.curTokenIs(token.RBRA) && !p.curTokenIs(token.EOF) && i > 0 {
//     stmt := p.parseStatement()
//     //p.debug("block stmt")
//     if stmt != nil {
//       block.Statements = append(block.Statements, stmt)
//     }
//     p.next() // Consume [;].
//     //p.debug("block stmt 2")
//     i--
//   }
//   //p.next() // Consume [}].
//   //p.debug("block exit")
// 	return block
// }
// 
// func (p *Parser) parseExpressionStatement() *ExpressionStatement {
// 	stmt := &ExpressionStatement{Token: p.curToken}
// 	stmt.Value = p.parseExpression(LOWEST)
// 	if p.nxtTokenIs(token.SCOLON) {
// 		p.next()
// 	}
// 	return stmt
// }
// 
// func (p *Parser) parseExpression(pre int) Expression {
// 	prefix := p.prefixParslets[p.curToken.Typ]
// 	if prefix == nil {
// 		p.error("No prefix parslet found for token [%s].", p.curToken.Literal)
// 		return nil
// 	}
// 	left := prefix()
// 
// 	for !p.nxtTokenIs(token.SCOLON) && pre < p.nxtTokenPrecedence() {
// 		infix := p.infixParslets[p.nxtToken.Typ]
// 		if infix == nil {
// 			return left
// 		}
// 		p.next()
// 		left = infix(left)
// 	}
// 
// 	return left
// }
// 
// func (p *Parser) parseIdentifer() Expression {
// 	expr := &Identifier{Token: p.curToken, Value: p.curToken.Literal}
// 	return expr
// }
// 
// func (p *Parser) parseInteger() Expression {
// 	expr := &Integer{Token: p.curToken}
// 	n, err := strconv.ParseInt(p.curToken.Literal, 10, 64)
// 	if err != nil {
// 		p.errorNext(token.INT)
// 	}
// 	expr.Value = n
// 	return expr
// }
// 
// func (p *Parser) parseBoolean() Expression {
// 	expr := &Boolean{Token: p.curToken}
// 	expr.Value = (p.curToken.Typ == token.TRUE)
// 	return expr
// }
// 
// func (p *Parser) parsePrefix() Expression {
// 	expr := &PrefixExpression{Token: p.curToken}
// 	expr.Operator = p.curToken.Literal
// 	p.next() // Consume operator.
// 	expr.Value = p.parseExpression(PREFIX)
// 	return expr
// }
// 
// func (p *Parser) parseBinary(left Expression) Expression {
// 	expr := &BinaryExpression{Token: p.curToken}
// 	expr.Left = left
// 	expr.Operator = p.curToken.Literal
// 	precedence := p.curTokenPrecedence()
// 	p.next() // Consume operator.
// 	expr.Right = p.parseExpression(precedence)
// 	return expr
// }
// 
// func (p *Parser) parseExpressionGroup() Expression {
// 	p.next() // Consume left parenthesis.
// 	expr := p.parseExpression(LOWEST)
//   if p.nxtTokenIs(token.RPAR) {
//     p.next() // Consume right parenthesis.
//     return expr
//   }
//   p.error("Missing closing parenthesis in [%s].", expr)
// 	return nil
// }
// 
// func (p *Parser) parseFunctionLiteral() Expression {
//   expr := &FunctionLiteral{Token: p.curToken}
//   if !p.expectNext(token.LPAR) {
//     return nil
//   }
// 
// 	return expr
// }
