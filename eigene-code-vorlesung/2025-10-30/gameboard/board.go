package gameboard

// Verwende eine String-Matrix für das Spielfeld.

// NewBoard liefert eine 3x3-Matrix aus Strings.
func NewBoard() [][]string {
	board := [][]string{}

	// Füge drei Zeilen zu board hinzu.
	for row := 0; row < 3; row++ {
		board = append(board, []string{" ", " ", " "})
	}

	return board
}
