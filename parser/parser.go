package parser

import (
	"bytes"
	"strconv"
)

type Parser struct {
	Buffer   *bytes.Buffer
	Byte     byte
	Filename string
}

// Skip comments and whitespace, advancing buffer
func (p *Parser) Skip() {
	for {
		b, _ := p.Buffer.ReadByte()
		if b == '#' {
			p.Buffer.ReadBytes('\n')
			continue
		}
		if b == ' ' || b == '\t' || b == '\r' || b == '\n' {
			continue
		}
		break
	}
	p.Buffer.UnreadByte()
}

// ReadInt at current buffer position
func (p *Parser) ReadInt() int {
	var i int
	var str string
	for {
		b, _ := p.Buffer.ReadByte()
		if b >= '0' && b <= '9' {
			str += (string)(b)
		} else {
			p.Buffer.UnreadByte()
			break
		}
	}
	i, _ = strconv.Atoi(str)
	return i
}

// NextToken skips to next token and returns it, advancing buffer
func (p *Parser) NextToken() interface{} {
	p.Skip()
	b, _ := p.Buffer.ReadByte()
	if b >= '0' && b <= '9' {
		p.Buffer.UnreadByte()
		return p.ReadInt()
	}
	if b == '{' {
		var ints []int
		for {
			v := p.NextToken()
			switch v := v.(type) {
			case rune:
				if v == '}' {
					return ints
				}
			case int:
				ints = append(ints, v)
			}
		}
	}
	if b == '}' {
		return (rune)(b)
	}
	return p.NextToken()
}
