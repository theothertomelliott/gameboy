package httpui

import (
	"bytes"
	"encoding/base64"
	"html/template"
	"image"
	"image/gif"
	"net/http"

	"github.com/faiface/pixel"
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
			<div class="tiles">
				{{range .Images}}
					<img class="tileimg" src="data:image/gif;base64,{{.Gif}}" />
				{{end}}
			</div>
		</body>
	</html>`

	type (
		tileData struct {
			Index int
			Gif   string
		}
		table struct {
			Images []tileData
		}
	)

	var data = table{}

	tiles := s.gb.PPU().GetTilesByIndex(0)
	tiles = append(tiles, s.gb.PPU().GetTilesByIndex(1)...)
	for tileIndex, tile := range tiles {
		dst := image.NewRGBA(image.Rect(0, 0, len(tile), len(tile[0])))
		for y, row := range tile {
			for x, value := range row {
				colorVal := 1.0 - (float64(value) / 4)
				color := pixel.RGB(colorVal, colorVal, colorVal)
				color = pixel.RGB(1, 1, 1)
				if value != 0 {
					dst.Set(x, y, color)
				}
			}
		}
		var b bytes.Buffer
		err := gif.Encode(&b, dst, &gif.Options{})
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		data.Images = append(
			data.Images,
			tileData{
				Index: tileIndex,
				Gif:   base64.StdEncoding.EncodeToString(b.Bytes()),
			},
		)
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
