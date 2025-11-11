package length

type Lenght int

// FromCentimeters konstruiert eine Länge aus einer Zentimeter-Angabe.
func FromCentimeters(C int) Lenght {
	return Lenght(C)
}

// Centimeters liefert die Länge in Zentimetern.
func (l Lenght) Centimeters() int {
	return int(l)
}

// Meters liefert die Länge in Metern.
func (l Lenght) Meters() int {
	return int(l)
}