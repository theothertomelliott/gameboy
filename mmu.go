package gameboy

type MMU struct {
	RAM            []byte
	CartridgeBanks [][]byte
}

func NewMMU() *MMU {
	return &MMU{
		RAM: make([]byte, 0x10000),
	}
}

// LoadCartridge loads a Cartridge ROM into memory
func (m *MMU) LoadCartridge(data []byte) {
}

func (m *MMU) Read8(pos uint16) byte {
	return m.RAM[pos]
}

func (m *MMU) Write8(pos uint16, value byte) {
	// TODO: Check for write to ROM area for bank switching
	m.RAM[pos] = value
}

func (m *MMU) Read16(pos uint16) uint16 {
	low := uint16(m.RAM[pos])
	high := uint16(m.RAM[pos+1])
	return low | high<<8
}

func (m *MMU) Write16(pos uint16, value uint16) {
	low := byte(value & 0xFF)
	high := byte(value >> 8)
	m.RAM[pos] = low
	m.RAM[pos+1] = high
}

// Clear resets the RAM to 0
func (m *MMU) Clear() {
	m.RAM = make([]byte, 0x10000)
}
