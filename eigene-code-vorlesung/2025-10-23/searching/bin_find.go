package searching

// / Liefert die Position von x in der Liste l,
// / oder liefert -1, falls x nicht in l vorkommt.
func BinFind2(l []int, x int) int {
	links := 0 // Hier beginnt der interessante Part.
	rechts := len(l) // Hier endet der interessante Part.
	for rechts > links {
		mitte := (rechts-links)/2 + links

		if l[mitte] == x {
			return mitte
		}
		if x < l[mitte] {
			rechts = mitte
		} else {
			links = mitte + 1
		}
	}

	return -1
}
