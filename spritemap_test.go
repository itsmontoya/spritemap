package spritemap

// Credits for the test sprite:
// https://kenney.nl/assets/pixel-platformer

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestSpritemap_NewFromFile(t *testing.T) {
	type args struct {
		filename    string
		tileSize    int
		tileSpacing int
	}

	tests := []struct {
		name string
		args args

		wantTiles   int
		wantRows    int
		wantColumns int
		wantErr     bool
	}{
		{
			name: "Basic",
			args: args{
				filename:    "test.png",
				tileSize:    18,
				tileSpacing: 1,
			},
			wantTiles:   180,
			wantRows:    9,
			wantColumns: 20,
			wantErr:     false,
		},
		{
			name: "Invalid file path",
			args: args{
				filename:    "missing.png",
				tileSize:    18,
				tileSpacing: 1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := NewFromFile(tt.args.filename, tt.args.tileSize, tt.args.tileSpacing)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected failure, got success")
				}

				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if s == nil {
				t.Fatal("NewFromFile() returned nil")
			}

			if s.rows != tt.wantRows {
				t.Errorf("expected %d rows, got %d", tt.wantRows, s.rows)
			}

			if s.columns != tt.wantColumns {
				t.Errorf("expected %d columns, got %d", tt.wantColumns, s.columns)
			}

			if len(s.tiles) != tt.wantTiles {
				t.Errorf("expected %d tiles, got %d", tt.wantTiles, len(s.tiles))
			}
		})
	}
}

func TestSpritemap_GetByIndex(t *testing.T) {
	s := &Spritemap{
		tiles: []*ebiten.Image{makeFakeTile(), makeFakeTile()},
	}

	tests := []struct {
		index     int
		expectErr bool
	}{
		{0, false},
		{1, false},
		{2, true},  // out of bounds
		{-1, true}, // negative index
	}

	for _, tt := range tests {
		_, err := s.GetByIndex(tt.index)
		if tt.expectErr && err == nil {
			t.Errorf("expected error for index %d but got none", tt.index)
		}
		if !tt.expectErr && err != nil {
			t.Errorf("unexpected error for index %d: %v", tt.index, err)
		}
	}
}

func TestSpritemap_GetByRowAndColumn(t *testing.T) {
	s := &Spritemap{
		columns: 3,
		tiles:   []*ebiten.Image{makeFakeTile(), makeFakeTile(), makeFakeTile(), makeFakeTile()},
	}

	tests := []struct {
		row, col  int
		expectErr bool
		expectIdx int
	}{
		{0, 0, false, 0},
		{0, 2, false, 2},
		{1, 0, false, 3},
		{1, 2, true, 5}, // out of bounds (only 4 tiles)
	}

	for _, tt := range tests {
		got, err := s.GetByRowAndColumn(tt.row, tt.col)
		if tt.expectErr && err == nil {
			t.Errorf("expected error for row=%d, col=%d", tt.row, tt.col)
		}
		if !tt.expectErr && err != nil {
			t.Errorf("unexpected error for row=%d, col=%d: %v", tt.row, tt.col, err)
		}
		if !tt.expectErr && got != s.tiles[tt.expectIdx] {
			t.Errorf("unexpected tile for row=%d, col=%d", tt.row, tt.col)
		}
	}
}

func TestSpritemap_getIndex(t *testing.T) {
	s := &Spritemap{columns: 10}

	tests := []struct {
		row, col int
		want     int
	}{
		{0, 0, 0},
		{0, 1, 1},
		{1, 0, 10},
		{2, 5, 25},
	}

	for _, tt := range tests {
		got := s.getIndex(tt.row, tt.col)
		if got != tt.want {
			t.Errorf("getIndex(%d, %d) = %d, want %d", tt.row, tt.col, got, tt.want)
		}
	}
}

func makeFakeTile() *ebiten.Image {
	return ebiten.NewImage(1, 1) // dummy tile
}
