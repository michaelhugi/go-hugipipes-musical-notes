package go_hugipipes_musical_notes

//MTemperament declares a MTemperament in music theory and it's octaves and notes with their frequencies
//MTemperament can be equal- or just scaled
type MTemperament interface {
	//AllOctaves returns all available octaves within the midi range
	AllOctaves() []MOctave
	//OctaveFromFrequency returns the octave a Frequency belongs to
	OctaveFromFrequency(frequency float64) (MOctave, error)
	//NoteFromFrequency returns the musical note a Frequency belongs to
	NoteFromFrequency(frequency float64) (MNote, error)
	//Octave returns the MOctave for the temperament from MGOctave
	Octave(octave MGOctave) MOctave
}
