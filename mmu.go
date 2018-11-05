package gameboy

import (
	"sync"
)

type MMU struct {
	RAM            []byte
	CartridgeBanks [][]byte

	mtx sync.RWMutex
}

type Range struct {
	Start uint16
	End   uint16
}

func NewMMU() *MMU {
	return &MMU{
		RAM: make([]byte, 0x10000),
	}
}

// LoadCartridge loads a Cartridge ROM into memory
func (m *MMU) LoadCartridge(data []byte) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	for index := 0x000; index < len(data); index++ {
		m.RAM[index] = data[index]
	}
	// // Write cartridge header to RAM
	// var index = 0
	// for index = 0x100; index < 0x150; index++ {
	// 	m.RAM[index] = data[index]
	// }
	// // Write remainder of bank 0
	// for ; index < 0x4000; index++ {
	// 	m.RAM[index] = data[index]
	// }
	// // Fill banks
	// for ; index < len(data); index += 0x3FFF {
	// 	end := index + 0x3FFF
	// 	if end > len(data)-1 {
	// 		end = len(data) - 1
	// 	}
	// 	m.CartridgeBanks = append(m.CartridgeBanks, data[index:end])
	// }

	// // Add the first bank to RAM
	// m.switchBank(0)
}

func (m *MMU) switchBank(bank byte) {
	if len(m.CartridgeBanks) > int(bank-1) {
		for i := 0; i < len(m.CartridgeBanks[bank]); i++ {
			m.RAM[0x4000+i] = m.CartridgeBanks[bank][i]
		}
	}
}

// ReadRange will return a range in RAM
func (m *MMU) ReadRange(r Range) []byte {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	return m.RAM[r.Start : r.End+1]
}

func (m *MMU) Read8(pos uint16) byte {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	return m.RAM[pos]
}

func (m *MMU) Write8(pos uint16, values ...byte) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	for _, value := range values {
		// Check for write to ROM area for bank switching
		if pos >= 0x150 && pos < 0x8000 {
			m.switchBank(value - 1)
			return
		}
		m.RAM[pos] = value
		pos++
	}
}

func (m *MMU) Read16(pos uint16) uint16 {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	low := uint16(m.RAM[pos])
	high := uint16(m.RAM[pos+1])
	return low | high<<8
}

func (m *MMU) Write16(pos uint16, value uint16) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	low := byte(value & 0xFF)
	high := byte(value >> 8)
	m.RAM[pos] = low
	m.RAM[pos+1] = high
}

// Clear resets the RAM to 0
func (m *MMU) Clear() {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	m.RAM = make([]byte, 0x10000)
}
