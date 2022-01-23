# Gameboy

A Gameboy (DMG-01) emulator, implemented in [Go](https://golang.org/).

References used include:
* [The Gameboy CPU Manual](http://marc.rawer.de/Gameboy/Docs/GBCPUman.pdf)
* [The Nesdev Forum](https://forums.nesdev.com/)
* [Gameboy Opcodes](http://www.pastraiser.com/cpu/gameboy/gameboy_opcodes.html)
* [Gameboy Emulation in Javascript](http://imrannazar.com/GameBoy-Emulation-in-JavaScript:-Input)
* [Blargg Gameboy Test Roms](https://github.com/retrio/gb-test-roms)

Graphics and keyboard input are implemented using [Pixel](https://github.com/faiface/pixel).

## Status

This project is very much a work in progress. Current milestones reached:

* Successfully runs the DMG-01 boot rom
* Passes tests in the cpu instruction roms from the [Blargg test set](https://github.com/retrio/gb-test-roms).
* Tetris is playable

Key remaining work:

* Sound support
* Cartridge bank switching

## Dependencies

Some prerequisite libraries are needed for Pixel, see [Pixel requirements](https://github.com/faiface/pixel#requirements) for more details.

## Installation

You can build and install this emulator using `go get`:

    $ go get -u github.com/theothertomelliott/gameboy/cmd/gameboy

## Usage

Assuming $GOPATH/bin is on your PATH, once installed, you can run the `gameboy` binary at the commandline, specifying a cartridge ROM filename as argument:

    $ gameboy path/to/rom.gb

You may also optionally provide a path to the Gameboy boot ROM.

## Debugging

By default, the emulator serves a debugger web app on port 8080. Visit http://localhost:8080 once the emulator is running to debug.