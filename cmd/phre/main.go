package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/brunetto/goutils/text"
	"github.com/pkg/errors"
)

type FnData struct {
	BaseName string
	Ext      StringSet
}

type FnDataSlice []*FnData

func (a FnDataSlice) Len() int           { return len(a) }
func (a FnDataSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a FnDataSlice) Less(i, j int) bool { return a[i].BaseName < a[j].BaseName }

type FnDataSet map[string]*FnData

func NewFnDataSet() FnDataSet {
	s := FnDataSet{}
	return s
}
func (s FnDataSet) Add(bn, ext string) {
	if s.Contains(bn) {
		s[bn].Ext.Add(ext)
	} else {
		s[bn] = &FnData{BaseName: bn, Ext: NewStringSet(ext)}
	}
}

func (s FnDataSet) Contains(str string) bool {
	_, exists := s[str]
	return exists
}

func (s FnDataSet) Remove(str string) {
	delete(s, str)
}

func (s FnDataSet) ToSlice() FnDataSlice {
	sl := FnDataSlice{}
	for _, v := range s {
		sl = append(sl, v)
	}
	return sl
}

type StringSet map[string]struct{}

func NewStringSet(strs ...string) StringSet {
	s := StringSet{}
	s.Add(strs...)
	return s
}

func (s StringSet) Add(strs ...string) {
	for _, str := range strs {
		s[str] = struct{}{}
	}
}

func (s StringSet) Contains(str string) bool {
	_, exists := s[str]
	return exists
}

func (s StringSet) Remove(strs ...string) {
	for _, str := range strs {
		delete(s, str)
	}
}

func SplitOnExtension(fn string) (fbn string, ext string) {
	ext = filepath.Ext(fn)
	fbn = strings.TrimSuffix(fn, ext)
	return fbn, ext
}

func main() {
	// file type extensions to consider
	extensions := NewStringSet(".cr2", ".CR2", ".jpeg", ".jpg", ".JPEG", ".JPG", ".tiff", ".TIFF",
		".png", ".PNG", ".NEF", ".nef", ".dng", ".DNG", ".MOV", ".mov", ".mpeg", ".MPEG", ".mpeg4", ".MPEG4", ".mp4", ".MP4")

	// get current directory name to be used as new base-name
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(errors.Wrap(err, "can't detect current directory"))
	}
	cwd = filepath.Base(cwd)

	// get all the file names in the current directory
	fns, err := filepath.Glob("./*")
	if err != nil {
		log.Fatal(errors.Wrap(err, "can't list files"))
	}

	// create a map (no repetition) with the old base-names and the related existing extensions
	fnsMap := NewFnDataSet()
	for _, fn := range fns {
		fbn, ext := SplitOnExtension(fn)
		if !extensions.Contains(ext) {
			continue
		}
		fnsMap.Add(fbn, ext)
	}

	// convert map to slice and sort it
	fnSlice := fnsMap.ToSlice()
	sort.Sort(fnSlice)

	// for each old base-name and extension, rename
	padlenght := len(strconv.Itoa(len(fnSlice))) + 1
	for i, fnd := range fnSlice {
		for ext, _ := range fnd.Ext {
			oldName := fnd.BaseName + ext
			newName := cwd + "-" + text.LeftPad(strconv.Itoa(i), "0", padlenght) + ext
			fmt.Printf("Renaming %v to %v\n", oldName, newName)
			err = os.Rename(oldName, newName)
			if err != nil {
				log.Fatal(errors.Wrap(err, "can't rename file"))
			}
		}
	}
}
