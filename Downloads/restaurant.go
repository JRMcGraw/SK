package main

import (
	"fmt"
//	"time"
)

var menu =[]string{
"eggs",
"cheese",
"grape juice",
"ice cream",
"frozen strawberries",
"wweet potato soup",
"hamburger",
"grapes",
"cooked rice",
"rebel salted caramel ice cream",
"banana",
"macadamia nuts",
"white chocolate",
"mandarine",
}

var servers = []string{
//"Bob",
//"Katie",
"Kelly",
"Kenna",
//"Bonnie",
//"Gram",
//"Grandad",
"JR",
}

type who_served_what struct {
	name	string
	meal	string
}

var all_meals_served []string

func main() {


	menu_items := make(chan string, len(menu)+1)	// Filled with Menu Items (Meals)
	served := make(chan who_served_what)		// Tells Chef a specific Meal was servered
	done := make(chan bool, len(servers))		// Tell each server to go home
	closed := make(chan bool)			// All servers are gone
	//clock := make(chan string)			// All servers are gone

	fmt.Println("Hired Waiters:", servers)
	fmt.Println("Hired a Chef.")

	for _, server_name := range servers {

		//go server(menu_items, served, done, server_name, clock)
		go server(menu_items, served, done, server_name )
	}


	fmt.Println("Restaurant Open.")

	go chef(menu_items, served, done, closed, len(servers))


	<- closed


	fmt.Println("Restaurant is closed.  Waiting for", len(servers), "servers to clock uot.")
	
	for i, waiter := range servers {


		//next_name := <- clock

		//fmt.Println("Server:", i, next_name, "went to the bar.")
		fmt.Println("Server:", i, waiter, "went to the bar.")
	}

	fmt.Println("All Servers clocked out.")

}

func chef(menu_items chan string, served chan who_served_what, done, closing chan bool, waiters int) {

	fmt.Println("Chef starting.")

	for _, meal := range menu {

		menu_items <- meal
		fmt.Println("Added menu_item:", meal)
	}
	fmt.Println("Menu created.", menu)

	close(menu_items)			// Menu created

	meals_served := 0

	for {

		who_served_what :=<- served		// A meal is served

		all_meals_served = append(all_meals_served, who_served_what.meal)

		meals_served++

		fmt.Println(who_served_what.name, "served meal", meals_served, ":", who_served_what.meal)

		if meals_served == len(menu) { break }
	}

	for i := 1; i <= waiters; i++ {

		done <- true		// Tell all servers to stop
	}

	fmt.Println(meals_served, "Meals were Served by", waiters, "Servers:", servers)

	closing <- true
}

//func server(menu_items chan string, served chan who_served_what, done chan bool, server_name string, clock chan string) {
func server(menu_items chan string, served chan who_served_what, done chan bool, server_name string) {

	fmt.Println(server_name, "is ready to wait tables:")

	my_meals := []string{}

	what_I_served := who_served_what{name: server_name}

	do_break := false

	for {

		fmt.Println(server_name, "Selecting next menu uitem.")

		select {
	
		case meal, ok := <- menu_items:

			if !ok { 
			
				fmt.Println("No more meals for", server_name, "to serve.")
				do_break = true
			}

			my_meals = append( my_meals , meal)

			fmt.Println(server_name, "served:", meal,". That makes", len(my_meals), my_meals)
	
			what_I_served.meal = meal

			served <- what_I_served		// Tells chef a meal was served

	
		case <- done:
		
			break 
		}

		fmt.Println(server_name, "Meal served.")

		if do_break { break }
	}
	
	fmt.Println(server_name, "served:", len(my_meals), "meals:", my_meals)
	fmt.Println(server_name, "clocking out.")

	//clock <- server_name

}


