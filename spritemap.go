package spritemap

import (
	"fmt"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

func New(src image.Image, tileSize, tileSpacing int) (out *Spritemap, err error) {
	var s Spritemap
	s.tileSize = tileSize
	s.tileSpacing = tileSpacing
	s.sprite = ebiten.NewImageFromImage(src)
	s.rows, s.columns = getRowsAndColumns(src, tileSize, tileSpacing)
	s.forEach(s.appendTile)
	out = &s
	return
}

func NewFromFile(filename string, tileSize, tileSpacing int) (out *Spritemap, err error) {
	var src image.Image
	if src, err = newImageFromFile(filename); err != nil {
		return
	}

	return New(src, tileSize, tileSpacing)
}

type Spritemap struct {
	sprite *ebiten.Image
	tiles  []*ebiten.Image

	tileSize    int
	tileSpacing int

	rows    int
	columns int
}

func (s *Spritemap) GetByIndex(index int) (out *ebiten.Image, err error) {
	switch {
	case index >= len(s.tiles):
		err = fmt.Errorf("index of <%d> is out of bounds, total length is <%d>", index, len(s.tiles))
		return
	case index < 0:
		err = fmt.Errorf("index of <%d> is less than 0", index)
		return
	default:
		return s.tiles[index], nil
	}
}

func (s *Spritemap) GetByRowAndColumn(row, column int) (out *ebiten.Image, err error) {
	index := s.getIndex(row, column)
	return s.GetByIndex(index)
}

func (s *Spritemap) getIndex(row, column int) (index int) {
	index = s.columns * row
	index += column
	return
}

func (s *Spritemap) appendTile(row, column int) {
	sx := column * (s.tileSize + s.tileSpacing)
	sy := row * (s.tileSize + s.tileSpacing)
	subImage := s.sprite.SubImage(image.Rect(sx, sy, sx+s.tileSize, sy+s.tileSize))
	tile := ebiten.NewImageFromImage(subImage)
	s.tiles = append(s.tiles, tile)

}

func (s *Spritemap) forEach(fn func(row, column int)) {
	for row := range s.rows {
		for column := range s.columns {
			fn(row, column)
		}
	}
}
