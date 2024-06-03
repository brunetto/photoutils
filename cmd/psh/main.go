package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/barasher/go-exiftool"
	"github.com/h2non/bimg"
	"github.com/pkg/errors"
)

func dieIf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	t0 := time.Now()
	defer func() { log.Print("done in ", time.Since(t0)) }()

	dieIf(os.MkdirAll("./sherwood", 0700))

	files, err := filepath.Glob("*.jpg")
	dieIf(err)

	for i, fn := range files {
		newName := fmt.Sprintf("./sherwood/Ph.biemmezeta%03d.jpg", i)

		original, err := bimg.Read(fn)
		dieIf(err)

		img := bimg.NewImage(original)

		opt := bimg.Options{}
		size, err := img.Size()
		dieIf(err)

		// major dimension resize to 1920
		if size.Height > size.Width {
			opt.Height = 1920
		} else {
			opt.Width = 1920
		}

		newBuf, err := bimg.Resize(original, opt)
		dieIf(err)

		fmt.Println(fn, "->", newName)

		err = bimg.Write(newName, newBuf)
		dieIf(err)
	}
}

func CopyrightMetadata(fn string) error {
	et, err := exiftool.NewExiftool()
	if err != nil {
		return errors.Wrapf(err, "can't init exiftools for %v", fn)
	}

	fileInfos := et.ExtractMetadata(fn)
	fileInfos[0].SetString("Author", "biemmezeta")
	fileInfos[0].SetString("Creator", "biemmezeta")
	fileInfos[0].SetString("Copyright", "biemmezeta")
	fileInfos[0].SetString("Copyright Notice", "biemmezeta")
	fileInfos[0].SetString("Comments", "biemmezeta")

	// for _, fileInfo := range fileInfos {
	// 	if fileInfo.Err != nil {
	// 		fmt.Printf("Error concerning %v: %v\n", fileInfo.File, fileInfo.Err)
	// 		continue
	// 	}

	// 	for k, v := range fileInfo.Fields {
	// 		fmt.Printf("[%v] %v\n", k, v)
	// 	}
	// }

	et.WriteMetadata(fileInfos)
	err = et.Close()
	if err != nil {
		return errors.Wrapf(err, "error while processing metadata for %v", fn)
	}

	return nil
}
