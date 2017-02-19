package main

import (
	"path/filepath"
	"log"
	"fmt"
	"os"
	"github.com/brunetto/goutils/text"
	"strconv"
)

func main () {
	var (
		fileNames  []string
		err        error
		fileName   string
		extensions []string
		ext        string
		idx        int
		cwd        string
		newName    string
	)

	extensions = []string{"cr2", "CR2", "jpeg", "jpg", "JPEG", "JPG", "tiff", "TIFF", "png", "PNG", "NEF", "nef", "MOV", "mov"}

	cwd, err = os.Getwd()
	if err != nil {
		log.Fatal("Error detecting current directory: ", err.Error())
	}
	cwd = filepath.Base(cwd)

	for _, ext = range extensions {
		fileNames, err = filepath.Glob("*." + ext)
		for idx, fileName = range fileNames {
			if err != nil{
				log.Fatal("Error listing directory files: " + err.Error())
			}
			newName = cwd + "-" + text.LeftPad(strconv.Itoa(idx), "0", 4) + "." + ext
			fmt.Printf("Renaming %v into %v\n", fileName, newName)
			err = os.Rename(fileName, newName)
			if err != nil {
				log.Fatal("Error renaming file: " + err.Error())
			}
		}
	}
}


