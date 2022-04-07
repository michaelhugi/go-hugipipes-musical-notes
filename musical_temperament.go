package go_hugipipes_musical_notes

//MTemperament declares a MTemperament in music theory and it's octaves and notes with their frequencies
//MTemperament can be equal- or just scaled
type MTemperament interface {
	//AllOctaves returns all available octaves within the midi range
	AllOctaves() []MOctave
	//OctaveFromFrequency returns the octave a Frequency belongs to
	OctaveFromFrequency(frequency float64) (*MOctave, error)
	//NoteFromFrequency returns the musical note a Frequency belongs to
	NoteFromFrequency(frequency float64) (*MNote, error)
	//MGNoteToNote is used to change a musical generic note to a musical note of the equivalent temperament in the corresponding octave
	MGNoteToNote(n *MGNote, o *MGOctave) *MNote
}
