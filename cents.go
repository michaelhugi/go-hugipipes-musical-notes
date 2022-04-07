package go_hugipipes_musical_notes

import "math"

//AddCents adds musical cents to a frequency. 100 Cents is a half-tone
func AddCents(frequency float64, cents float64) float64 {
	return frequency * math.Pow(2, 1/(1200/cents))

}

var log2 = math.Log(2)

//Cents returns how many cents offsetFrequency is off to baseFrequency
func Cents(baseFrequency float64, offsetFrequency float64) float64 {
	return (1200 * math.Log(offsetFrequency/baseFrequency)) / log2
}
