package spritemap

import (
	"fmt"
	"image"
	"os"
)

func newImageFromFile(filename string) (img image.Image, err error) {
	var f *os.File
	if f, err = os.Open(filename); err != nil {
		err = fmt.Errorf("error opening source file: %v", err)
		return
	}
	defer f.Close()

	if img, _, err = image.Decode(f); err != nil {
		err = fmt.Errorf("error decoding image: %v", err)
		return
	}

	return
}

func getRowsAndColumns(img image.Image, tileSize, tileSpacing int) (rows, columns int) {
	b := img.Bounds()
	w := b.Max.X
	h := b.Max.Y
	columns = (w + tileSpacing) / (tileSize + tileSpacing)
	rows = (h + tileSpacing) / (tileSize + tileSpacing)
	return
}
