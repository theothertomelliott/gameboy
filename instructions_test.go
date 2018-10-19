package gameboy_test

import (
	"testing"

	"github.com/theothertomelliott/gameboy"
)

func TestLD(t *testing.T) {
	cpu := gameboy.NewCPU()

	src := &gameboy.Register{}
	src.Write(100)
	dst := &gameboy.Register{}

	cpu.LD(dst, src)

	if dst.Read() != 100 {
		t.Errorf("dst: expected %d, got %d", 100, dst.Read())
	}
}

func TestLD16(t *testing.T) {
	cpu := gameboy.NewCPU()

	src := &gameboy.RegisterPair{}
	src.Write(100)
	dst := &gameboy.RegisterPair{}

	cpu.LD(dst, src)

	if dst.Read() != 100 {
		t.Errorf("dst: expected %d, got %d", 100, dst.Read())
	}
}

func TestLDI(t *testing.T) {
	cpu := gameboy.NewCPU()

	src := &gameboy.Register{}
	src.Write(10)
	dst := &gameboy.Register{}

	cpu.LDI(dst, src)

	if dst.Read() != 11 {
		t.Errorf("dst: expected %d, got %d", 11, dst.Read())
	}
}

func TestLDD(t *testing.T) {
	cpu := gameboy.NewCPU()

	src := &gameboy.Register{}
	src.Write(10)
	dst := &gameboy.Register{}

	cpu.LDD(dst, src)

	if dst.Read() != 9 {
		t.Errorf("dst: expected %d, got %d", 9, dst.Read())
	}
}
