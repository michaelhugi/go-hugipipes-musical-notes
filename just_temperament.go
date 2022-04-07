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

//justTempToBaseNote calculates the frequency of the base-note of the equal temperament from a frequency of a given note
//The function stays in the same octave and will *2 in case of an octave overflow
func justTempToBaseNote(baseNote MGNote, note MGNote, freq float64) float64 {
	offs := int(baseNote - note)
	multi := 1.0
	if offs > 0 {
		offs -= 12
		multi = 2.0
	}
	return freq * justTempNoteDiffToFreqFactor(offs) * multi
}

//justTempFromBaseNote calculates the frequency of a given note from the frequency of the base note for just temperament
//The function stays in the same octave and will /2 in case of an octave overflow
func justTempFromBaseNote(baseNote MGNote, note MGNote, baseNoteFreq float64) float64 {
	offs := int(note - baseNote)
	divider := 1.0
	if offs < 0 {
		divider = 2.0
		offs += 12
	}
	return baseNoteFreq * justTempNoteDiffToFreqFactor(offs) / divider
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
