package aufgabe4

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// ElementSums erwartet zwei int-Listen l1 und l2.
// Sie liefert eine Liste, die an jeder Position
// jeweils die Summe der beiden Elemente enthält.
//
// Annahmen für die Berechnung:
// Falls eine Liste kürzer ist als die andere, soll für die Berechnung der
// hinteren Werte ihr letztes Element verwendet werden.
// Für leere Listen soll für die Berechnung ggf. 0 verwendet werden.
func ElementSums(l1, l2 []int) []int {
	result := []int{}
	// TODO
	max := max(len(l1), len(l2))
	c1, c2 := 0, 0
	for i := 0; i < max; i++ {
		if i < len(l1) {
			c1 = l1[i]
		}
		if i < len(l2) {
			c2 = l2[i]
		}
		result = append(result, c1 + c2)
	}
	return result
}
