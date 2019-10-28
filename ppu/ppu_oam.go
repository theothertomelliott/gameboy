package ppu

import "github.com/theothertomelliott/gameboy/binary"

type OAM []byte

func (o OAM) X() byte {
	return o[1] - 8
}

func (o OAM) Y() byte {
	return o[0] - 16
}

func (o OAM) Tile() byte {
	return o[2]
}

func (o OAM) flags() byte {
	if len(o) < 4 {
		return 0
	}
	return o[3]
}

func (o OAM) Priority() bool {
	return binary.Bit(7, o.flags()) != 0
}

func (o OAM) YFlip() bool {
	return binary.Bit(6, o.flags()) != 0
}

func (o OAM) XFlip() bool {
	return binary.Bit(5, o.flags()) != 0
}

func (o OAM) Palette() byte {
	return binary.Bit(4, o.flags())
}
