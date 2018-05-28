package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/vgmdj/utils/chars"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func generateFileData(year int, fileName string) (string, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	divisions := "map[string]string{"

	buff := bytes.NewBuffer(content)
	for {
		line, err := buff.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		lineScanner := bufio.NewScanner(strings.NewReader(line))
		lineScanner.Split(bufio.ScanWords)
		lineScanner.Scan()
		source := lineScanner.Text()
		if source == "Source" {
			continue
		}

		//revision
		lineScanner.Scan()

		lineScanner.Scan()
		code := lineScanner.Text()

		lineScanner.Scan()
		name := lineScanner.Text()

		divisions += fmt.Sprintf(`"%s": "%s",`, code, name) + "\n"
	}
	divisions += "}"

	return divisions, nil
}

func generateFile(year int, fileName string) {

	gocode := fmt.Sprintf(`package gbdata

import "github.com/vgmdj/utils/area"

func init() {
s := new(area.Selector)
s.Register("%d", rev%d)
}
	
var rev%d =	`, year, year, year)

	data, err := generateFileData(year, fileName)
	if err != nil {
		log.Fatal(err)
		return
	}

	gocode = fmt.Sprintf("%s %s", gocode, data)

	err = ioutil.WriteFile(fmt.Sprintf("gbdata/gb2260_%d.go", year), []byte(gocode), os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}

}

// before use it, you should use go build and then execute the generate (or generate.exe)
func main() {
	dir := "data/"

	latest := 201802
	years := []int{1986, 1987, 1988, 1989,
		1990, 1991, 1992, 1993, 1994,
		1995, 1996, 1997, 1998, 1999,
		2000, 2001, 2002, 2003, 2004,
		2005, 2006, 2007, 2008, 2009,
		2010, 2011, 2012, 2013, 2014,
		2015, 2016, 2017}

	for _, v := range years {
		fileName := fmt.Sprintf("%s%d12.tsv", dir, v)
		generateFile(v, fileName)
	}

	generateFile(chars.TakeLeftInt(latest, 4), fmt.Sprintf("%s%d.tsv", dir, latest))

	cmd := exec.Command("go", "fmt", "./gbdata/")
	cmd.Start()

	return
}
