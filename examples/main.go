package main

import (
	"fmt"
	"io"
	"os"

	"github.com/mebyus/er"
)

const (
	ErCReadSource er.Code = er.CodeGap + 1 + iota
)

func readSource(path string) (src io.ReadCloser, e er.Er) {
	src, err := os.Open(path)
	if err != nil {
		e = er.From(er.COpenFile, err)
		return
	}
	return
}

func main() {
	src, e := readSource("main1.go")
	if e != nil {
		fmt.Println(e.Up(ErCReadSource, "read source"))
		os.Exit(1)
	}
	defer src.Close()
	_, err := io.Copy(os.Stdout, src)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
