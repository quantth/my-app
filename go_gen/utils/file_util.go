package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func GetFile(f string) []byte {
	b, err := ioutil.ReadFile(f)

	if err != nil {
		panic(err)
	}

	return b
}

func GetDir(dir string) []os.FileInfo {
	p, err := ioutil.ReadDir(dir)

	if err != nil {
		panic(err)
	}

	return p
}

func WriteFile(fileName string, b bytes.Buffer) {
	err := ioutil.WriteFile(fileName, b.Bytes(), 0644)

	if err != nil {
		panic(err)
	}
}

func ConvertDate(date string) string {
	myDate, _ := time.Parse("2006-01-02", date)
	return fmt.Sprintf("%s", myDate.Format("January 02, 2006"))
}
