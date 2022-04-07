package go_hugipipes_musical_notes

import "fmt"

//MOctave describes a musical octave
type MOctave interface {
	//LowerFrequency returns the lowest frequency that belongs to the octave
	LowerFrequency() float64
	//UpperFrequency returns the highest frequency that belongs to the octave
	UpperFrequency() float64
	//String returns the readable name of the octave
	String() string
	//FrequencyBelongsToNote returns, if a tested frequency belongs to the octave or not
	FrequencyBelongsToNote(frequency float64) bool
	//OctaveNumber returns the readable octave name of the note
	OctaveNumber() int
	//AllNotes returns all notes that belong to the octave
	AllNotes() []MNote
}

//MGOctave s a musical generic octave that can be declared to use an octave without specifying a temperament.
//MGOctave can be converted to MOctave through MTemperament
type MGOctave int

const (
	OctaveMinus1 MGOctave = -1
	Octave0               = 0
	Octave1               = 1
	Octave2               = 2
	Octave3               = 3
	Octave4               = 4
	Octave5               = 5
	Octave6               = 6
	Octave7               = 7
	Octave8               = 8
	Octave9               = 9
)

func octaveToNumber(o MGOctave) int {
	switch o {

	case OctaveMinus1:
		return -1
	case Octave0:
		return 0
	case Octave1:
		return 1
	case Octave2:
		return 2
	case Octave3:
		return 3
	case Octave4:
		return 4
	case Octave5:
		return 5
	case Octave6:
		return 6
	case Octave7:
		return 7
	case Octave8:
		return 8
	case Octave9:
		return 9
	}
	panic(fmt.Sprintf("Octave %v not declared for octaveToNumber", o))
}
