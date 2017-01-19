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

	// match regexp, api / regexp
	regElems map[string]*regexp.Regexp

	// matches, api: date / count
	matches map[string]map[string]int64
}

func (p *Parser) Match(str string) []string {
	a := p.reg.FindStringSubmatch(str)
	return a
}

func (p *Parser) ParseFiles(fileName ...string) error {
	for _, name := range fileName {
		err := p.Statistic(name)
		if err != nil {
			log.Errorf("failed to statistic %s, err: %v", name, err)
			continue
		}
	}
	return nil
}

func (p *Parser) Save(filename string) error {
	resStatistic := &ResStatic{}

	for k, v := range p.matches {
		api := ResApi{
			Api: k,
		}

		for kk, vv := range v {
			elem := ResElem{
				Date:  kk,
				Count: vv,
			}
			api.ResElems = append(api.ResElems, elem)
		}
		resStatistic.ResApis = append(resStatistic.ResApis, api)
	}

	err := resStatistic.Save(filename)
	return err
}

func (p *Parser) Import(files ...string) error {
	for _, file := range files {
		rs := &ResStatic{}
		err := rs.Import(file)
		if err != nil {
			log.Errorf("failed to import %s, err: %v", file, err)
			return err
		}
		for _, api := range rs.ResApis {
			s, _ := p.matches[api.Api]
			for _, e := range api.ResElems {
				count, ok := s[e.Date]
				if ok {
					s[e.Date] = count + e.Count
					break
				}
				s[e.Date] = e.Count
			}
		}
	}
	return nil
}

func (p *Parser) Statistic(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		log.Errorf("failed to open %s, err: %v", fileName, err)
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//process the log file.
		line := scanner.Text()
		elems := p.Match(line)
		if len(elems) < 10 {
			log.Debugf("skip %s", line)
			continue
		}
		for k, v := range p.regElems {
			apiName := elems[8]
			if v.MatchString(apiName) {
				calcs, _ := p.matches[k]
				date := elems[1]
				count, ok := calcs[date]
				if ok {
					calcs[date] = count + 1
					break
				}
				calcs[date] = 1
			}
		}
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
