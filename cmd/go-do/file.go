package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	godo "github.com/FryDay/go-do"
)

func filename(td *godo.ToDo) string {
	return filepath.Join(configDir, fmt.Sprintf("go-do_%s.json", td.TimeCreated.Format("060102150405.999")))
}

func writeToDoFile(td *godo.ToDo) {
	go ioutil.WriteFile(filename(td), td.Bytes(), 0600)
}

func deleteToDoFile(td *godo.ToDo) {
	go os.Remove(filename(td))
}
