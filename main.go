package main

import (
	"errors"
	"fmt"
	"os"
)

type menuAction int
type seat byte

const (
	minPlaces = 60
	priceHigh = 10
	priceLow  = 8
)

const (
	freeSeat seat = 'S'
	soldSeat seat = 'B'
)

const (
	menuExit menuAction = iota
	menuSeats
	menuBuy
	menuStatistics
)

var (
	places           [][]seat
	rows, seats      int
	income           int
	totalIncome      int
	purchasedTickets int
	totalTickets     int
)

func init() {
	fmt.Println("Enter the number of rows:")
	fmt.Scan(&rows)
	fmt.Println("Enter the number of seats in each row:")
	fmt.Scan(&seats)
	fmt.Println()

	places = make([][]seat, rows)
	for i := 0; i < rows; i++ {
		places[i] = make([]seat, seats)
		for j := 0; j < seats; j++ {
			places[i][j] = 'S'
		}
		totalIncome += getPrice(i+1) * seats
	}

	totalTickets = rows * seats
}

func askPlace() (int, int) {
	var row, seat int

	fmt.Println("Enter a row number:")
	fmt.Scan(&row)
	fmt.Println("Enter a seat number in that row:")
	fmt.Scan(&seat)
	fmt.Println()

	return row, seat
}

func markSold(row, seat int) error {
	if row < 1 || row > rows || seat < 1 || seat > seats {
		return errors.New("Wrong input!")
	}

	selectedSeat := &places[row-1][seat-1]
	if *selectedSeat == soldSeat {
		return errors.New("That ticket has already been purchased!")
	}
	*selectedSeat = soldSeat

	return nil
}

func getPrice(row int) int {
	var price int

	if rows*seats < minPlaces || row <= rows/2 {
		price = priceHigh
	} else {
		price = priceLow
	}

	return price
}

func buyTicket() {
	var row, seat int

	for {
		row, seat = askPlace()
		if err := markSold(row, seat); err != nil {
			fmt.Println(err)
			fmt.Println()
		} else {
			price := getPrice(row)
			purchasedTickets++
			income += price

			fmt.Printf("Ticket price: $%d\n\n", price)
			break
		}
	}
}

func displayPlaces() {
	fmt.Printf("Cinema:\n ")

	// display seat numbers
	for i := 1; i <= seats; i++ {
		fmt.Printf(" %d", i)
	}

	// display row number with seat value
	for i := 0; i < rows; i++ {
		fmt.Printf("\n%d", i+1)
		for j := 0; j < seats; j++ {
			fmt.Printf(" %c", places[i][j])
		}
	}

	fmt.Print("\n\n")
}

func displayStatistics() {
	fmt.Printf("Number of purchased tickets: %d\n", purchasedTickets)
	fmt.Printf("Percentage: %.2f%%\n", float32(purchasedTickets)/float32(totalTickets)*100)
	fmt.Printf("Current income: $%d\n", income)
	fmt.Printf("Total income: $%d\n", totalIncome)
	fmt.Println()
}

func chooseAction() menuAction {
	var action menuAction

	fmt.Println("1. Show the seats")
	fmt.Println("2. Buy a ticket")
	fmt.Println("3. Statistics")
	fmt.Println("0. Exit")
	fmt.Scan(&action)
	fmt.Println()

	return action
}

func doAction(action menuAction) {
	switch action {
	case menuExit:
		os.Exit(0)
	case menuSeats:
		displayPlaces()
	case menuBuy:
		buyTicket()
	case menuStatistics:
		displayStatistics()
	}
}

func main() {
	for {
		action := chooseAction()
		doAction(action)
	}
}
