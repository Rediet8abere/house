package main

import (
	"fmt"
	"strconv"
)

// Declares a public struct that represents a house.
// The struct should contain variables numberOfRooms,
// city, address, and price, each assigned their
// appropriate data type (string, int, etc)

// When your program runs, allow users
// to enter information for one to many houses.

// After the user submits all their information,
// print a list of the houses for sale in the following format:
type house struct {
	numberOfRooms int
	city          string
	address       string
	price         int
}

var HouseList []house

func main() {

	fmt.Print("How many houses would you like to enter? ")
	var houseNum string
	fmt.Scanln(&houseNum)
	i := 0
	numHouse, _ := strconv.Atoi(houseNum)

	// var ahouse House
	// technically go doesnt have while, but
	// for can be used while in go.
	for i < numHouse {
		fmt.Println("House", i+1)
		fmt.Print("How many rooms? ")
		var numRooms int
		fmt.Scanln(&numRooms)
		fmt.Print("City? ")
		var cityAt string
		fmt.Scanln(&cityAt)
		fmt.Print("address? ")
		var address string
		fmt.Scanln(&address)
		fmt.Print("price? ")
		var price int
		fmt.Scanln(&price)

		HouseList = append(HouseList, house{
			numberOfRooms: numRooms, city: cityAt,
			address: address, price: price,
		})

		i += 1
	}

	fmt.Println("\n")
	for i := 0; i < len(HouseList); i++ {
		fmt.Printf("%s %s %d %d\n", HouseList[i].address, HouseList[i].city, HouseList[i].numberOfRooms, HouseList[i].price)
	}

}
