package parser

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sort"

	log "github.com/Sirupsen/logrus"
)

type ResElem struct {
	Date  string `json:"date"`
	Count int64  `json:"count"`
}

type ResApi struct {
	Api      string    `json:"api"`
	ResElems []ResElem `json:"static"`
}

type ResStatic struct {
	ResApis []ResApi
}

func (rs *ResStatic) Save(fileName string) error {
	for _, api := range rs.ResApis {
		sort.Sort(ResElemSorter(api.ResElems))
	}

	b, err := json.Marshal(rs.ResApis)
	if err != nil {
		log.Errorf("failed to convert to json for file %s, err: %v", fileName, err)
		return err
	}

	err = ioutil.WriteFile(fileName, b, os.ModePerm)
	if err != nil {
		log.Errorf("failed to write file %s, err: %v", fileName, err)
		return err
	}
	return nil
}

func (rs *ResStatic) Import(fileName string) error {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Errorf("failed to read %s, err: %v", fileName, err)
		return err
	}
	err = json.Unmarshal(b, &rs.ResApis)
	if err != nil {
		log.Errorf("failed to parse json %s, err: %v", fileName, err)
		return err
	}

	return nil
}

type ResElemSorter []ResElem

func (res ResElemSorter) Len() int {
	return len(res)
}

func (res ResElemSorter) Swap(i, j int) {
	res[i], res[j] = res[j], res[i]
}

func (res ResElemSorter) Less(i, j int) bool {
	return res[i].Date < res[j].Date
}
