package gameboy

import (
	"fmt"
	"testing"
)

func TestNewTile(t *testing.T) {
	var tests = []struct {
		name     string
		input    []byte
		expected [][]byte
	}{
		{
			name: "empty",
		},
		{
			name: "blank",
			input: []byte{
				0, 0,
				0, 0,
				0, 0,
				0, 0,
				0, 0,
				0, 0,
				0, 0,
				0, 0,
			},
			expected: [][]byte{
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			name: "A",
			input: []byte{
				0x7C, 0x7C,
				0x0, 0xC6,
				0xC6, 0x00,
				0, 0xFE,
				0xC6, 0xC6,
				0, 0xC6,
				0xC6, 0,
				0, 0,
			},
			expected: [][]byte{
				{0, 3, 3, 3, 3, 3, 0, 0},
				{2, 2, 0, 0, 0, 2, 2, 0},
				{1, 1, 0, 0, 0, 1, 1, 0},
				{2, 2, 2, 2, 2, 2, 2, 0},
				{3, 3, 0, 0, 0, 3, 3, 0},
				{2, 2, 0, 0, 0, 2, 2, 0},
				{1, 1, 0, 0, 0, 1, 1, 0},
			},
		},
	}
	for _, test := range tests {
		tile := NewTile(test.input)
		var match = true
		var got string
		var expected string
		for j, row := range test.expected {
			for i, val := range row {
				if tile.At(i, j) != val {
					match = false
				}
				got += fmt.Sprintf("%X ", tile.At(i, j))
				expected += fmt.Sprintf("%X ", val)
			}
			got += "\n"
			expected += "\n"
		}
		if !match {
			t.Errorf("Expected:\n%v\nGot:\n%v", expected, got)
		}
	}
}
