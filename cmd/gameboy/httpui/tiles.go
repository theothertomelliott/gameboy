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
		oam struct {
			X, Y byte
			Tile byte
		}
		page struct {
			Background string
			Window     string
			Screen     string
			Tilesets   []tileset
			OAM        []oam
		}
	)

	var (
		data = page{}
		err  error
	)

	for i := 0; i < 2; i++ {
		lcdcont := gameboy.LCDControl(i * 0xFF)
		tiles := gameboy.GetTilesForRange(s.gb.MMU(), lcdcont.TilePatternTableAddress())
		ts := tileset{
			Index: byte(i),
		}
		for tileIndex, tile := range tiles {
			i, err := renderTileToBase64(tile)
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

	for _, o := range gameboy.GetSpriteData(s.gb.MMU()) {
		if o.X() > 160 || o.Y() > 144 {
			continue
		}
		data.OAM = append(data.OAM, oam{
			X:    o.X(),
			Y:    o.Y(),
			Tile: o.Tile(),
		})
	}

	data.Background, err = renderImageToBase64(s.gb.PPU().RenderBackground())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data.Window, err = renderImageToBase64(s.gb.PPU().RenderWindow())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data.Screen, err = renderImageToBase64(s.gb.PPU().RenderScreen())
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

func renderTileToBase64(tile gameboy.Tile) (string, error) {
	var b bytes.Buffer
	err := gif.Encode(&b, tile.ToImage(), &gif.Options{})
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}

func renderImageToBase64(tile image.Image) (string, error) {
	var b bytes.Buffer
	err := gif.Encode(&b, tile, &gif.Options{})
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}
