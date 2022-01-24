package mmu

import (
	"fmt"
	"log"

	"github.com/theothertomelliott/gameboy/ioports"
)

type MMU struct {
	RAM            []byte
	CartridgeData  []byte
	CartridgeBanks [][]byte
	ROM            []byte

	inROM bool

	tracer MMUTracer

	testOutput []byte
}

type MMUTracer interface {
	AddMMU(pos uint16, values ...byte)
}

type Range struct {
	Start uint16
	End   uint16
}

func New(tracer MMUTracer) *MMU {
	return &MMU{
		RAM:    make([]byte, 0x10000),
		tracer: tracer,
	}
}

func (m *MMU) LoadROM(data []byte) {
	m.ROM = nil
	for _, b := range data {
		m.ROM = append(m.ROM, b)
	}
	m.inROM = true
}

// LoadCartridge loads a Cartridge ROM into memory
func (m *MMU) LoadCartridge(data []byte) error {
	// This prevents use of the combined blarg test roms
	cartridgeType := data[0x0147]
	log.Printf("Cartridge type: 0x%x\n", cartridgeType)
	// if cartridgeType != 0x00 && cartridgeType != 0x01 {
	// 	panic("Cartridge type unsupported")
	// }

	// Check DMG/CGB header
	needCGB := data[0x143] == 0xC0
	if needCGB {
		return fmt.Errorf("Cartridge expects Color Game Boy support")
	}

	m.CartridgeData = data
	m.testOutput = nil
	m.RAM = make([]byte, 0x10000)
	for index := 0x000; index < len(data); index++ {
		m.RAM[index] = data[index]
	}
	// Write cartridge header to RAM
	var index = 0
	for index = 0x100; index < 0x150; index++ {
		m.RAM[index] = data[index]
	}
	// Write remainder of bank 0
	for ; index < 0x4000; index++ {
		m.RAM[index] = data[index]
	}
	// Fill banks
	for ; index < len(data); index += 0x3FFF {
		end := index + 0x3FFF
		if end > len(data)-1 {
			end = len(data) - 1
		}
		m.CartridgeBanks = append(m.CartridgeBanks, data[index:end])
	}
	log.Printf("Loaded %d cartridge banks\n", len(m.CartridgeBanks))

	// Add the first bank to RAM
	m.switchBank(0)

	return nil
}

// ResetCartridge resets the content of memory to the cartridge
func (m *MMU) ResetCartridge() {
	m.LoadCartridge(m.CartridgeData)
}

func (m *MMU) switchBank(bank byte) {
	log.Println("Switching to bank:", bank)
	for i := 0; i < len(m.CartridgeBanks[bank]); i++ {
		m.RAM[0x4000+i] = m.CartridgeBanks[bank][i]
	}
}

// ReadRange will return a range in RAM
func (m *MMU) ReadRange(r Range) []byte {
	return m.RAM[r.Start : r.End+1]
}

func (m *MMU) Read8(pos uint16) byte {
	if pos <= 0xFF && m.inROM {
		return m.ROM[pos]
	}

	// Unused area
	if pos >= 0xFEA0 && pos <= 0xFEFF {
		return 0xFF
	}

	return m.RAM[pos]
}

func (m *MMU) Write8(pos uint16, values ...byte) {
	if pos == 0x2000 {
		m.switchBank(values[0])
	}
	if pos == ioports.DMACONT {
		start := uint16(values[0]) * 0x100
		end := start + 0x100
		m.Write8(0xFE00, m.ReadRange(Range{
			Start: start,
			End:   end,
		})...)
		return
	}

	// Unused area
	if pos >= 0xFEA0 && pos <= 0xFEFF {
		return
	}

	if m.tracer != nil {
		m.tracer.AddMMU(pos, values...)
	}

	// Turn off ROM
	if pos == 0xFFFF && m.inROM {
		m.inROM = false
	}

	if pos == 0xFF02 && values[0] == 0x81 {
		c := m.Read8(0xFF01)
		m.testOutput = append(m.testOutput, c)
	}

	for _, value := range values {
		m.RAM[pos] = value
		pos++
	}
}

func (m *MMU) Read16(pos uint16) uint16 {
	// Unused area
	if pos >= 0xFEA0 && pos <= 0xFEFF {
		return 0xFFFF
	}

	if pos+1 <= 0xFF && m.inROM {
		low := uint16(m.ROM[pos])
		high := uint16(m.ROM[pos+1])
		return low | high<<8
	}
	low := uint16(m.RAM[pos])
	high := uint16(m.RAM[pos+1])
	return low | high<<8
}

func (m *MMU) Write16(pos uint16, value uint16) {
	// Unused area
	if pos >= 0xFEA0 && pos <= 0xFEFF {
		return
	}

	low := byte(value & 0xFF)
	high := byte(value >> 8)
	m.RAM[pos] = low
	m.RAM[pos+1] = high
	if m.tracer != nil {
		m.tracer.AddMMU(pos, low, high)
	}
}

// Clear resets the RAM to 0
func (m *MMU) Clear() {
	m.RAM = make([]byte, 0x10000)
}

func (m *MMU) TestOutput() string {
	return string(m.testOutput)
}
