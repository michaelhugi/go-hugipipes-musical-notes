package go_hugipipes_musical_notes

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
