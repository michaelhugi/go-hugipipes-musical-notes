# go-hugipipes-musical-notes

This is a go-library to handle musical notes theory for different temperaments.

Currently the library supports equal temperament and just temperament

To use the library see the following examples to get started

### Equal temperament

Equal temperament is the most common used temperament, especially in western music. All half-notes have the same
interval so all scales sound the same

```go
package main

import (
	mn "go-hugipipes-musical-notes"
)

func main() {
	//This will create a equal temperament where A4 will be at 440Hz
	temp := mn.NewMTemperamentEqual(440.0)
	println("%+v", temp)
	//temp implements the interface MTemperament where you have a lot of possibilities to work with octaves and notes
}

```

### Just temperament

Just temperament is scaled around a base-note. All other notes have a rational factor to the base notes. This makes this
temperament sound very clean for scales with the base notes but other base notes sound weird.This temperament is used by
bagpipes for example, where all notes must be in perfect harmony to the drones to avoid beat
frequencies

```go
package main

import (
	mn "go-hugipipes-musical-notes"
)

func main() {
	baseNote := mn.B
	pitchNote := mn.A
	pitchOctave := mn.Octave4
	pitch := 440.0
	//This will create a temperament with all notes harmonic to C, where A4 will be 440Hz
	temp := mn.NewMTemperamentJust(baseNote, pitchNote, pitchOctave, pitch)
	println("%+v", temp)
	//temp implements the interface MTemperament where you have a lot of possibilities to work with octaves and notes
}

```