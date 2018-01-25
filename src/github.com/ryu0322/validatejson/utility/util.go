package utility

import (
	"io/ioutil"
	"strings"
	"log"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// SJIS→UTF-8
func StoU(encstr string) string {

	iostr := strings.NewReader(encstr)
	reader := transform.NewReader(iostr, japanese.ShiftJIS.NewDecoder())
	ret, err := ioutil.ReadAll(reader)

	if err != nil {
		log.Fatal(err)
	}

	return string(ret)
}

// UTF-8→SJIS
func UtoS(encstr string) string {
	iostr := strings.NewReader(encstr)
	reader := transform.NewReader(iostr, japanese.ShiftJIS.NewEncoder())
	ret, err := ioutil.ReadAll(reader)

	if err != nil {
		log.Fatal(err)
	}

	return string(ret)
}