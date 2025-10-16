package fizzbuzz

import "fmt"

// Fizz gibt die Zahlen von 1 bis n aus und ersetzt dabei
// jede durch x teilbare Zahl durch "fizz" und jede durch
// y teilbare Zahl durch "buzz".
// Bei Zahlen, die sowohl durch x als auch durch y teilbar sind,
// wird "fizzbuzz" ausgegeben.

func FizzCompact(x, y, n int) {
	for i := 1; i <= n; i++ {
		// Wenn i weder durch x noch durch y teilbar ist
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
			continue
		}
		if i%x == 0 {
			fmt.Print("fizz")
		}
		if i%y == 0 {
			fmt.Print("buzz")
		}
		fmt.Println()
	}
}
