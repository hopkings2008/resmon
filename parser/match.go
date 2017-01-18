package parser

import (
	"bufio"
	"os"
	"regexp"

	log "github.com/Sirupsen/logrus"
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

	regElems map[string]*regexp.Regexp

	matches map[string]map[string]int64
}

func (p *Parser) Match(str string) []string {
	a := p.reg.FindStringSubmatch(str)
	return a
}

func (p *Parser) Static(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		log.Errorf("failed to open %s, err: %v", fileName, err)
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//process the log file.
		log.Debugf("%s", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Errorf("failed to read %s, err: %v", fileName, err)
		return err
	}
	return nil
}

func CreateParser(regStr string, want int, suffix string) (*Parser, error) {
	p := &Parser{
		RegStr:   regStr,
		Want:     want,
		Suffix:   suffix,
		matches:  make(map[string]map[string]int64),
		regElems: make(map[string]*regexp.Regexp),
	}
	p.matches["image_get"] = make(map[string]int64)
	p.matches["image_add"] = make(map[string]int64)
	p.reg = regexp.MustCompile(regStr)
	p.regElems["image_get"] = regexp.MustCompile("^https://res.shiqichuban.com/v1/image/get/")
	p.regElems["image_add"] = regexp.MustCompile("^https://res.shiqichuban.com/v1/image/add")
	return p, nil
}
