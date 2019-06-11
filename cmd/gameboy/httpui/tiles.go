package httpui

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/gif"
	"net/http"

	"github.com/theothertomelliott/gameboy"
)

// HandleTiles renders a page displaying the current tile sets
func (s *Server) HandleTiles(w http.ResponseWriter, r *http.Request) {
	type (
		tileData struct {
			Index int
			Gif   string
		}
		tileset struct {
			Index byte
			Tiles []tileData
		}
		page struct {
			Background string
			Sprites    string
			Screen     string
			Tilesets   []tileset
		}
	)

	var (
		data = page{}
		err  error
	)

	// Bit 3 - BG Tile Map Display Select     (0=9800-9BFF, 1=9C00-9FFF)
	patternMaps := []gameboy.Range{
		gameboy.Range{
			Start: 0x8800,
			End:   0x97FF,
		},
		gameboy.Range{
			Start: 0x8000,
			End:   0x8FFF,
		},
	}

	for i, tileRange := range patternMaps {
		tiles := s.gb.PPU().GetTilesForRange(tileRange)
		ts := tileset{
			Index: byte(i),
		}
		for tileIndex, tile := range tiles {
			i, err := renderImageToBase64(tile)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			ts.Tiles = append(
				ts.Tiles,
				tileData{
					Index: tileIndex,
					Gif:   i,
				},
			)
		}
		data.Tilesets = append(
			data.Tilesets,
			ts,
		)
	}

	data.Background, err = renderImageToBase64(s.gb.PPU().RenderBackground())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data.Screen, err = renderImageToBase64(s.gb.PPU().RenderScreen())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Init 256x256 background
	sprites := image.NewRGBA(image.Rect(0, 0, 256, 256))
	data.Sprites, err = renderImageToBase64(s.gb.PPU().RenderSprites(sprites))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t, err := loadTemplate("tiles.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func renderImageToBase64(tile image.Image) (string, error) {
	var b bytes.Buffer
	err := gif.Encode(&b, tile, &gif.Options{})
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}
