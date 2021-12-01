package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	johnPrice := computePrice(145.90, 3)
	paulPrice := computePrice(26.32, 10)
	robPrice := computePrice(189.21, 2)

	total := johnPrice + paulPrice + robPrice
	fmt.Printf("TOTAL: %0.2f $\n", total)

	vacant := vacantRooms()
	fmt.Println("Vacant rooms: ", vacant)

	printHeader()
}

// compute the price of a room based on its rate per night and the numer of night
func computePrice(rate float32, nights int) float32 {
	return rate * float32(nights)
}

func vacantRooms() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(100)
}

func printHeader() {
	fmt.Println("Hotel Golang")
	fmt.Println("San Francisco, CA")
}
