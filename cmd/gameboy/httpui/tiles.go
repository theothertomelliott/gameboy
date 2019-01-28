package httpui

import (
	"bytes"
	"encoding/base64"
	"html/template"
	"image"
	"image/color"
	"image/gif"
	"net/http"
)

// HandleTiles renders a page displaying the current tile sets
func (s *Server) HandleTiles(w http.ResponseWriter, r *http.Request) {
	const tpl = `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>Gameboy - Tiles</title>
			<style>
				body {font-family: "Courier New", Courier, serif;}
				.tiles {
					font-size: 0;
				}
				.tileimg {
					display: inline-block;
					width: 32px;
					height: 32px;
				}
			</style>
		</head>
		<body>
			<h1>Tiles</h1>
			{{range .Tilesets}}
			<h2>Tile set {{ .Index }}</h2>
			<div class="tiles">
				{{range .Tiles}}
					<img class="tileimg" src="data:image/gif;base64,{{.Gif}}" />
				{{end}}
			</div>
			{{end}}
			<h1>Background</h1>
			<img src="data:image/gif;base64,{{.Background}}" />
			<h1>Sprites</h1>
			<img src="data:image/gif;base64,{{.Sprites}}" />
		</body>
	</html>`

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
			Tilesets   []tileset
		}
	)

	var (
		data = page{}
		err  error
	)
	for i := byte(0); i < 2; i++ {
		tiles := s.gb.PPU().GetTilesByIndex(i)
		ts := tileset{
			Index: i,
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

	// Init 256x256 background
	var sprites = make([][]byte, 256)
	for i := range sprites {
		sprites[i] = make([]byte, 256)
	}
	data.Sprites, err = renderImageToBase64(s.gb.PPU().RenderSprites(sprites))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t, err := template.New("tiles").Parse(tpl)
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

func renderImageToBase64(tile [][]byte) (string, error) {
	dst := image.NewRGBA(image.Rect(0, 0, len(tile), len(tile[0])))
	for y, row := range tile {
		for x, value := range row {
			colorVal := 255 - ((256 / 4) * uint8(value))
			dst.Set(x, y, color.RGBA{colorVal, colorVal, colorVal, 255})
		}
	}
	var b bytes.Buffer
	err := gif.Encode(&b, dst, &gif.Options{})
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}
