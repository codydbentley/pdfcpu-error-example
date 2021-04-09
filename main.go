package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func main() {
	// external inputs
	var (
		x    = 34
		y    = 42
		refW = 920
		refH = 1190
	)

	// merge page1 and page2
	err := doFileMerge("./input.pdf", "./page1.pdf", "./page2.pdf")

	// open input file to be edited
	inFile, err := os.Open("./input.pdf")
	if err != nil {
		log.Fatalf("unable to open input: %v", err)
	}
	defer inFile.Close()

	// open output file
	outFile, err := os.OpenFile("./output.pdf", os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("unable to create outfile file: %v", err)
	}
	defer outFile.Close()

	// watermark image generated with Image Magick tool: "convert.exe -size 468x68 canvas:#FFFFFF watermark.png"

	// get dimensions of page 1
	pageDimensions, err := api.PageDims(inFile, nil)
	if err != nil {
		log.Fatalf("error retrieving page dimensions: %v", err)
	}
	pageW := pageDimensions[0].Width
	pageH := pageDimensions[0].Height

	// create watermark to be used
	wm := pdfcpu.DefaultWatermarkConfig()
	wm.Mode = pdfcpu.WMImage
	wm.OnTop = true
	wm.Diagonal = 0
	wm.InpUnit = pdfcpu.POINTS
	wm.FileName = "./watermark.png"
	wm.Pos = pdfcpu.TopLeft
	wm.Scale = scale(refW, pageW)
	wm.ScaleEff = float64(refW)
	wm.ScaleAbs = true
	wm.Dx = int(math.Round(calcPos(x, refW, pageW)))
	wm.Dy = int(math.Round(-calcPos(y, refH, pageH)))

	// apply watermark to page 1
	err = api.AddWatermarks(inFile, outFile, []string{"1"}, wm, nil)
	if err != nil {
		log.Fatalf("add watermark error: %v", err)
	}
}

func doFileMerge(destination string, files ...string) error {
	inFiles := make([]*os.File, 0, len(files))
	for _, path := range files {
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("unable to open file %s: %v", path, err)
		}
		inFiles = append(inFiles, file)
	}
	defer func() {
		for _, file := range inFiles {
			file.Close()
		}
	}()
	rs := make([]io.ReadSeeker, len(inFiles))
	for i, file := range inFiles {
		rs[i] = file
	}
	buf := bytes.Buffer{}
	if err := api.Merge(rs, &buf, nil); err != nil {
		return fmt.Errorf("unable to merge documents: %v", err)
	}
	outFile, err := os.OpenFile(destination, os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("unable to create destination file: %v", err)
	}
	defer outFile.Close()
	if _, err = outFile.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("unable to write to destination file: %v", err)
	}
	return nil
}

func scale(ref int, actual float64) float64 {
	return actual / float64(ref)
}

func calcPos(target, refX int, realX float64) float64 {
	return (float64(target) * realX) / float64(refX)
}
