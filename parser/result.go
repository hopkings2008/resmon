package parser

import (
	"encoding/json"
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

func (rs *ResStatic) Save(fileName string) {
}

func (rs *ResStatic) Import(fileName string) {
}
