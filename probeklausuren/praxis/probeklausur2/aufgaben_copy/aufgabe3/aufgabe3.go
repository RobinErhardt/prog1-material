package aufgabe3

import "math"

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * RANDBEDINGUNG: Die Funktion muss rekursiv sein.
 * ERREICHBARE PUNKTE: 10
 */

// CountSquares erwartet eine Liste von Zahlen.
// CountSquares liefert die Anzahl der QuadratzahlenZahlen in der Liste.
func CountSquares(list []int) int {
	// TODO
	if len(list) == 0 {
		return 0
	}
	head := list[0]
	result := CountSquares(list[1:])

	if IsSquare(head) {
		result++
	}
	return result
}

func IsSquare(n int) bool {
	r := int(math.Sqrt(float64(n)))
	return r*r == n
}