package go_hugipipes_musical_notes

//MOctave describes a musical octave
type MOctave interface {
	//LowerFrequency returns the lowest frequency that belongs to the octave
	LowerFrequency() float64
	//UpperFrequency returns the highest frequency that belongs to the octave
	UpperFrequency() float64
	//String returns the readable name of the octave
	String() string
	//FrequencyBelongsToOctave returns, if a tested frequency belongs to the octave or not
	FrequencyBelongsToOctave(frequency float64) bool
	//Octave returns the readable octave name of the note
	Octave() MGOctave
	//AllNotes returns all notes that belong to the octave
	AllNotes() []MNote
	//Note returns the musical note for a note in the concert pitch
	Note(note MGNote) MNote
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
