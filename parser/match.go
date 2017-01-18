package parser

import (
	//"fmt"
	"regexp"
)

type Parser struct {
	// regex string for parser.
	RegStr string
	// the segment to calculate.
	Want int
	// the suffix for the match.
	Suffix string
	// begin time
	Begin string

	reg *regexp.Regexp

	matches map[string]map[string]int64
}

func (p *Parser) Match(str string) []string {
	a := p.reg.FindStringSubmatch(str)
	return a
}

func CreateParser(regStr string, want int, suffix string) (*Parser, error) {
	p := &Parser{
		RegStr:  regStr,
		Want:    want,
		Suffix:  suffix,
		matches: make(map[string]map[string]int64),
	}
	p.reg = regexp.MustCompile(regStr)
	return p, nil
}
