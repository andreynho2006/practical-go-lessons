package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	//var rooms, roomsOccupied uint8 = 100, 10
	var rooms, roomsOccupied int = 100, rand.Intn(100)

	fmt.Println("rooms: ", rooms, " roomsOccupied: ", roomsOccupied)

	// Example 1: is there more rooms than roomsOccupied
	fmt.Println("boolean expression: 'rooms > roomsOccupied' is equal to: ")

	fmt.Println(rooms > roomsOccupied) //*\label{condExpBool1}

	//Example 2: is there more roomsOccupied than rooms
	fmt.Println("boolean expression: 'roomsOccupied > rooms' is equal to: ")
	fmt.Println(roomsOccupied > rooms) //*\label{condExpBool2}

	//Example 3: is roomOccupied equals to rooms
	fmt.Println("boolean expression: 'roomsOccupied == rooms' is equal to: ")
	fmt.Println(roomsOccupied == rooms) //*\label{condExpBool3}

	// hotelName := "The Gopher Hotel"
	// fmt.Println("Hotel " + hotelName)
	// //var roomsAvailable uint8
	// var roomsAvailable int

	// roomsAvailable = rooms - roomsOccupied //*\label{cond1diff}
	// fmt.Println(roomsAvailable, " rooms available")
}
