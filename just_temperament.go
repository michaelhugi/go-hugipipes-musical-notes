package go_hugipipes_musical_notes

const justTempF1 = 16.0 / 15.0
const justTempF2 = 9.0 / 8.0
const justTempF3 = 6.0 / 5.0
const justTempF4 = 5.0 / 4.0
const justTempF5 = 4.0 / 3.0
const justTempF6 = 45.0 / 32.0
const justTempF7 = 3.0 / 2.0
const justTempF8 = 8.0 / 5.0
const justTempF9 = 5.0 / 3.0
const justTempF10 = 9.0 / 5.0
const justTempF11 = 15.0 / 8.0
const justTempF12 = 2.0

//FrequencyOfBaseNoteInJustScale calculates the frequency of the base-note of the equal temperament from a frequency of a given note
//The function stays in the same octave and will *2 in case of an octave overflow
func FrequencyOfBaseNoteInJustScale(baseNote MGNote, note SimpleNote) float64 {
	offs := int(baseNote - note.note)
	multi := 1.0
	if offs > 0 {
		offs -= 12
		multi = 2.0
	}
	return note.frequency * justTempNoteDiffToFreqFactor(offs) * multi
}

func justTempNoteDiffToFreqFactor(noteDiff int) float64 {
	switch noteDiff {
	case -12:
		return 1 / justTempF12
	case -11:
		return 1 / justTempF11
	case -10:
		return 1 / justTempF10
	case -9:
		return 1 / justTempF9
	case -8:
		return 1 / justTempF8
	case -7:
		return 1 / justTempF7
	case -6:
		return 1 / justTempF6
	case -5:
		return 1 / justTempF5
	case -4:
		return 1 / justTempF4
	case -3:
		return 1 / justTempF3
	case -2:
		return 1 / justTempF2
	case -1:
		return 1 / justTempF1
	case 0:
		return 1
	case 1:
		return 1 * justTempF1
	case 2:
		return 1 * justTempF2
	case 3:
		return 1 * justTempF3
	case 4:
		return 1 * justTempF4
	case 5:
		return 1 * justTempF5
	case 6:
		return 1 * justTempF6
	case 7:
		return 1 * justTempF7
	case 8:
		return 1 * justTempF8
	case 9:
		return 1 * justTempF9
	case 10:
		return 1 * justTempF10
	case 11:
		return 1 * justTempF11
	case 12:
		return 1 * justTempF12
	}
	panic("Invalid notes for pitching in just tuning")
}

//JustScaleFromBaseNote calculates the scale of all notes in just temperament from the base-note
func JustScaleFromBaseNote(baseNote float64) []float64 {
	n := make([]float64, 12)
	//C starting from C
	n[0] = baseNote
	//Cis starting from C
	n[1] = baseNote * 16.0 / 15.0
	//D starting from C
	n[2] = baseNote * 9.0 / 8.0
	//Dis starting from C
	n[3] = baseNote * 6.0 / 5.0
	//E starting from C
	n[4] = baseNote * 5.0 / 4.0
	//F starting from C
	n[5] = baseNote * 4.0 / 3.0
	//Fis starting from C
	n[6] = baseNote * 45.0 / 32.0
	//G starting from C
	n[7] = baseNote * 3.0 / 2.0
	//Gis starting from C
	n[8] = baseNote * 8.0 / 5.0
	//A starting from C
	n[9] = baseNote * 5.0 / 3.0
	//Ais starting from C
	n[10] = baseNote * 9.0 / 5.0
	//B starting from C
	n[11] = baseNote * 15.0 / 8.0

	return n
}

//JustScaleFromBaseNoteWithinOctave calculates the scale of all notes in just temperament from the base-note
//but overflows when octave changes at C
func JustScaleFromBaseNoteWithinOctave(baseNote SimpleNote) []SimpleNote {
	fs := JustScaleFromBaseNote(baseNote.frequency)

	n := make([]SimpleNote, len(fs))
	cPassed := false
	for i, freq := range fs {
		no := AddHalftones(baseNote.note, i)
		if i > 0 {
			if no == C {
				cPassed = true
			}
		}
		if cPassed {
			freq = freq / 2.0
		}
		n[i] = NewSimpleNote(no, freq)
	}
	return n
}

//SimpleNote contains a frequency ant what not it is. It's used for calculations of complex tunings
type SimpleNote struct {
	frequency float64
	note      MGNote
}

func NewSimpleNote(note MGNote, frequency float64) SimpleNote {
	return SimpleNote{
		frequency: frequency,
		note:      note,
	}
}
