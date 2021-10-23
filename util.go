package main

import (
	"io/ioutil"
	"path"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func read(file string) string {
	b, err := ioutil.ReadFile(path.Join(thisDir, file))
	check(err)
	return string(b)
}
