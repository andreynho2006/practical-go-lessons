package main

import (
	"fmt"
)

func main() {
	// now := time.Now()
	// fmt.Println(now)

	// hexadecimal octal acii unicode
	num := 2548
	n2 := 0x9F4
	fmt.Printf("Num: %X\n", num)
	fmt.Printf("%X\n", n2)
	fmt.Printf("Decimal: %d\n", n2)

	b := make([]byte, 0)
	b = append(b, 255)
	b = append(b, 10)
	fmt.Println(b)

	fmt.Printf("Binary: %b\n", n2)

	raw := `spring rain:
	browsing under an umbrella
	at the picture-book store`
	fmt.Println(raw)

	interpreted := "i love spring"
	fmt.Println(interpreted)

	var aRune rune = 'Z'
	fmt.Printf("Unicode Code point of &#39;%c&#39;: %U\n", aRune, aRune)

	const hotelName string = "Gopher Hotel"

	const longitude = 24.806078
	const latitude = -78.233027

	var occupancy int = 12

	fmt.Println(hotelName)
	fmt.Println(longitude)
	fmt.Println(latitude)
	fmt.Println(occupancy)
}
