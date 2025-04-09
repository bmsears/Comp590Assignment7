package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Customer struct {
	id int
}

var newCustomerChan = make(chan Customer)
var waitingRoomChan = make(chan Customer, 6)

func receptionist() {
	for {
		customer := <-newCustomerChan
		fmt.Printf("Receptionist greets customer %d\n", customer.id)

		select {
			case waitingRoomChan <- customer:
				fmt.Printf("Customer %d is waiting in the waiting room\n", customer.id)
			default:
				fmt.Printf("Waiting room is full, customer %d is turned away\n", customer.id)
		}


	}
}

func barber() {
	for {
		customer := <-waitingRoomChan
		fmt.Printf("Barber is cutting hair for customer %d\n", customer.id)

		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

		fmt.Printf("Barber finished cutting hair for customer %d\n", customer.id)
	}
}

func handleCustomer(customer Customer) {
    fmt.Printf("Customer %d is arriving at the shop.\n", customer.id)
    newCustomerChan <- customer
}

func generateCustomers() {
	customerID := 1
	for {
		customer := Customer{id: customerID}
		go handleCustomer(customer)
		customerID++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	}
}

func main() {
	fmt.Println("Starting sleeping barber simulation...")

	go receptionist()

	go barber()

	go generateCustomers()

	select {}
}


