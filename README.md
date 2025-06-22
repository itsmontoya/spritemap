# Spritemap

A lightweight Go library for handling sprite sheets and tiled maps, built for 2D games using Ebiten. Spritemap makes it easy to extract and render individual tiles from a sprite sheet with configurable tile size and spacing.

## 🚀 Features

- Simple API for working with tile-based sprite sheets
- Get individual tiles by index or row/column
- Automatically handles tile spacing
- Built to work seamlessly with [Ebiten](https://ebiten.org/)
- Pure Go with no external dependencies

## 📦 Installation

```bash
go get github.com/itsmontoya/spritemap@latest
```

## ⚙️ Quick Start

```go
import (
    "embed"
    "image"
    _ "image/png"
    "log"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/itsmontoya/spritemap"
)

//go:embed tiles.png
var tileFS embed.FS

var sm *spritemap.Spritemap

func init() {
    f, err := tileFS.Open("tiles.png")
    if err != nil {
        log.Fatal(err)
    }
    img, _, err := image.Decode(f)
    if err != nil {
        log.Fatal(err)
    }

    sm, err = spritemap.New(img, 16, 0) // 16x16 tiles, 0px spacing
    if err != nil {
        log.Fatal(err)
    }
}
```

Inside your Ebiten `Draw` method:

```go
tile, err := sm.GetByRowAndColumn(0, 1)
if err == nil {
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Translate(50, 50)
    screen.DrawImage(tile, op)
}
```

## 📚 Usage

### `spritemap.New`

Creates a `Spritemap` from a decoded image.

```go
sm, err := spritemap.New(img, tileSize, tileSpacing)
```

- `img` – a decoded `image.Image` (e.g. from `image.Decode`)
- `tileSize` – width and height of each tile (e.g. `16`)
- `tileSpacing` – spacing between tiles in pixels (can be `0`)

### `spritemap.NewFromFile`

Creates a `Spritemap` from a PNG file.

```go
sm, err := spritemap.NewFromFile("tiles.png", 16, 0)
```

### `(*Spritemap) GetByIndex`

Gets a tile by its flat index.

```go
tile, err := sm.GetByIndex(5)
```

### `(*Spritemap) GetByRowAndColumn`

Gets a tile by row and column.

```go
tile, err := sm.GetByRowAndColumn(1, 3)
```

## 📚 Example Use Case

To draw a level from a grid of tile indices:

```go
for y, row := range level {
    for x, tileIndex := range row {
        tile, err := sm.GetByIndex(tileIndex)
        if err != nil {
            continue
        }

        op := &ebiten.DrawImageOptions{}
        op.GeoM.Translate(float64(x*16), float64(y*16))
        screen.DrawImage(tile, op)
    }
}
```

## 💡 Tips & Notes

- Tiles are read left-to-right, top-to-bottom
- `embed` makes it easy to bundle sprite sheets into your Go binary

## 📚 Examples

Check out the included `examples/tiles.go`, which shows:
- Loading a sprite sheet
- Rendering a tile in Ebiten

## 💡 Tips & Tricks

- Use `embed` to bundle your sprite sheet for easy distribution.
- Tile indexes are read **left-to-right, top-to-bottom** starting at 0.

## 🧑‍💻 Contributing

Contributions welcome! Please open PRs for bug fixes, new features, or examples.

## 📄 License

Licensed under the [MIT License](https://github.com/itsmontoya/spritemap/blob/main/LICENSE). Have fun building your games! 🎮
