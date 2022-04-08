package go_hugipipes_musical_notes

import "fmt"

//MNote describes a musical note
type MNote interface {
	//LowerFrequency returns the lowest frequency that belongs to the note
	LowerFrequency() float64
	//UpperFrequency returns the highest frequency that belongs to the note
	UpperFrequency() float64
	//ExactFrequency returns the exact frequency this note should have
	ExactFrequency() float64
	//OffsetInCents returns the offset in cents from the compareFrequency
	OffsetInCents(compareFrequency float64) float64
	//Octave returns the musical octave the note belongs to
	Octave() MOctave
	//MidiNoteNumber returns the midi note number
	MidiNoteNumber() uint8
	//String returns the readable name of the note
	String() string
	//FrequencyBelongsToNote returns, if a tested frequency belongs to the note or not
	FrequencyBelongsToNote(frequency float64) bool
	//Note returns the generic type of the note
	Note() MGNote
}

//MGNote s a musical generic note that can be declared to use a note without specifying a temperament.
//MGNote can be converted to MNote through MTemperament
type MGNote int

const (
	C   MGNote = 0
	Cis        = 1
	D          = 2
	Dis        = 3
	E          = 4
	F          = 5
	Fis        = 6
	G          = 7
	Gis        = 8
	A          = 9
	Ais        = 10
	B          = 11
)

func AddHalftones(n MGNote, i int) MGNote {
	nn := -1

	switch n {
	case C:
		nn = 0
	case Cis:
		nn = 1
	case D:
		nn = 2
	case Dis:
		nn = 3
	case E:
		nn = 4
	case F:
		nn = 5
	case Fis:
		nn = 6
	case G:
		nn = 7
	case Gis:
		nn = 8
	case A:
		nn = 9
	case Ais:
		nn = 10
	case B:
		nn = 11
	default:
		panic(fmt.Sprintf("Note %d is not a valid note number", n))
	}
	nn += i
	for nn > B {
		nn -= 12
	}
	switch nn {
	case 0:
		return C
	case 1:
		return Cis
	case 2:
		return D
	case 3:
		return Dis
	case 4:
		return E
	case 5:
		return F
	case 6:
		return Fis
	case 7:
		return G
	case 8:
		return Gis
	case 9:
		return A
	case 10:
		return Ais
	case 11:
		return B
	default:
		panic(fmt.Sprintf("Note %d is not a valid note number", nn))
	}
}
