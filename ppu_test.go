package gameboy

import (
	"testing"
)

func TestRenderTile(t *testing.T) {
	result := renderTile(
		[]byte{
			0x7C, 0x7C,
			0x00, 0xC6,
			0xC6, 0x00,
			0x00, 0xFE,
			0xC6, 0xC6,
			0x00, 0xC6,
			0xC6, 0x00,
			0x00, 0x00,
		},
		0xE4,
	)
	expected := [][]byte{
		{0, 3, 3, 3, 3, 3, 0, 0},
		{2, 2, 0, 0, 0, 2, 2, 0},
		{1, 1, 0, 0, 0, 1, 1, 0},
		{2, 2, 2, 2, 2, 2, 2, 0},
		{3, 3, 0, 0, 0, 3, 3, 0},
		{2, 2, 0, 0, 0, 2, 2, 0},
		{1, 1, 0, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	for line, lineValues := range result {
		for col, value := range lineValues {
			if value != expected[line][col] {
				t.Errorf("expected %v, got %v", expected, result)
				return
			}
		}
	}
}
