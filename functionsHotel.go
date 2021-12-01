package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const hotelName = "Gopher Paris Inn"
	const totalRooms = 134
	const firstRoomNumber = 110

	rand.Seed(time.Now().UTC().UnixNano())
	roomsOccupied := rand.Intn(totalRooms)
	roomsAvailable := totalRooms - roomsOccupied

	occupancyRate := (float32(roomsOccupied) / float32(totalRooms)) * 100
	occupancyLevel := occupancyLevel(occupancyRate)

	fmt.Println("Hotel:", hotelName)
	fmt.Println("Number of rooms", totalRooms)
	fmt.Println("Rooms available", roomsAvailable)
	fmt.Println("                  Occupancy Level:", occupancyLevel)
	fmt.Printf("                  Occupancy Rate: %0.2f %%\n", occupancyRate)

	if roomsAvailable > 0 {
		fmt.Println("Rooms:")
		for i := 0; roomsAvailable > i; i++ {
			roomNumber := firstRoomNumber + i
			size := rand.Intn(6) + 1
			nights := rand.Intn(10) + 1
			printRoomDetails(roomNumber, size, nights)
		}
	} else {
		fmt.Println("No rooms available for tonight")
	}

}

// displey information about a room
func printRoomDetails(roomNumber, size, nights int) {
	fmt.Println(roomNumber, ":", size, "people /", nights, " nights ")
}

// retrieve occupancyLevel from an occupancyRate
// from 0% to 30% occupancy rate return low
// from 31% to 69% occupancy rate return medium
// from 70% to 100% occupancy rate return high
func occupancyLevel(occupancyRate float32) string {
	if occupancyRate > 70 {
		return "High"
	} else if occupancyRate > 30 {
		return "Medium"
	} else {
		return "Low"
	}
}
