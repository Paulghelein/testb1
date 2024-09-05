package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 48; i < 58; i++ {
		for i2 := 48; i2 < 58; i2++ {
			for i3 := 48; i3 < 58; i3++ {
				for i4 := 48; i4 < 58; i4++ {
					z01.PrintRune(rune(i))
					z01.PrintRune(rune(i2))
					z01.PrintRune(32)
					z01.PrintRune(rune(i3))
					z01.PrintRune(rune(i4))
					z01.PrintRune(44)
				}

			}
		}
	}
}
