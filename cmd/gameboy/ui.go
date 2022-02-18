package main

import "github.com/theothertomelliott/gameboy"

type UI interface {
	Run(gb *gameboy.DMG)
}
