# Missing from the Gameboy CPU Manual

# LHDL?

Loads the value of hl, not the value at the memory location of HL.

TODO: Confirm in commit history

# POP AF

Only write the top 4 bits.

# DAA

Has a very vague description

Better outline
https://ehaskins.com/2018-01-30%20Z80%20DAA/
Table:
http://www.z80.info/z80syntx.htm#DAA

A great resource:
http://imrannazar.com/GameBoy-Emulation-in-JavaScript:-Input

# CP

The carry and half carry behavior are reversed:
https://forums.nesdev.com/viewtopic.php?f=20&t=12861&p=234005#p234005

# SUB, SBC, DEC, etc

"no borrow" - incorrect