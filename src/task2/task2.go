package task2

import "math"

type Envelope struct {
	Width  float64
	Height float64
}

func (e *Envelope) Validate() bool {
	if e.Width <= 0 || e.Height <= 0 {
		return false
	}
	return true
}

func ValidateParams(e1, e2 Envelope) (ok bool, reason string) {

	if !(e1.Validate() && e2.Validate()) {
		return false, "Envelopes must have positive width and height"
	}
	return true, ""
}

func CanEncloseEnvelopes(e1, e2 Envelope) (res bool, minEnvelope uint8) {

	// Если у конвертов минимальная или максимальная стороны равны то конверты нельзя влаживать друг в друга
	minE1 := math.Min(e1.Width, e1.Height)
	minE2 := math.Min(e2.Width, e2.Height)
	maxE1 := math.Max(e1.Width, e1.Height)
	maxE2 := math.Max(e2.Width, e2.Height)

	switch {
	case minE1 == minE2 || maxE1 == maxE2:
		return
	case minE1 < minE2 && maxE1 < maxE2:
		return true, 1
	case minE1 > minE2 && maxE1 > maxE2:
		return true, 2
	}
	return
}
