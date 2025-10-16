package fizzbuzz

import "fmt"

// Fizz gibt die Zahlen von 1 bis n aus und ersetzt dabei
// jede durch x teilbare Zahl durch "fizz" und jede durch
// y teilbare Zahl durch "buzz".
// Bei Zahlen, die sowohl durch x als auch durch y teilbar sind,
// wird "fizzbuzz" ausgegeben.

func FizzImproved(x, y, n int) {
	for i := 1; i <= n; i++ {
		// wenn i durch x und y teilbar ist, gib "fizzbuzz" aus
		if i%x == 0 && i%y == 0 {
			fmt.Println("fizzbuzz")
		} else
		// wenn i durch x teilbar ist, gib "fizz" aus
		if i%x == 0 {
			fmt.Println("fizz")
		} else
		// wenn i durch y teilbar ist, gib "buzz" aus
		if i%y == 0 {
			fmt.Println("buzz")
		} else {
			// sonst gib i aus
			fmt.Println(i)
		}
	}
}
