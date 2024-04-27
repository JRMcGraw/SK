package main 

import ( 

	"fmt" 
	"time" 
	"os" 
	"strings" 
	"encoding/json"
) 

type Person struct { 
    First	string 
    Last	string 
    ID		int 
    Home	Address 
    Phone	string 
    Email	string 
    Cards	[]string 
} 

type Address struct { 
    ID		int 
    Street	string 
    AptSuite	string 
    City	string 
    State	string 
    Zip		string 
    Country	string 
} 

type Card struct { 
    ID		int 
    Type	string
    Number	string
    ExpDate	string
    CCV		string
}

type Item struct {
    Name	string
    Price	int
    Quantity	int
}

type Order struct {
    Name	string
    Items	[]Item
    OrderDate	time.Time
    ShipDate	time.Time
    ShipMethod	string
}

var people []Person

var addresses []Address
var customerAddress map[int]Address

var allCards []Card
var cardSet map[int][]Card

var items []Item

var peopleFile = "people.json"
var adressFile = "address.json"
var cardsFile = "cards.json"
var itemsFile = "items.json"

// Each Customer's Address

func loadAddresses(runningTests bool) {

	addressJSON, _ := os.ReadFile(adressFile)
	json.Unmarshal([]byte(addressJSON),&addresses)

	customerAddress = make(map[int]Address)

	for _, address := range addresses {

		customerAddress[address.ID] = address
	}

	//fmt.Println(fmt.Sprintf("\n       %s       \n", "Customer Addresses"))

	//for i, address := range customerAddress {

	//	if !runningTests { fmt.Println(i, address) }
	//}
}

// Each Customer's Credit Card(s)

func loadCards(runningTest bool) {

	cardJSON, _ := os.ReadFile(cardsFile)
	json.Unmarshal([]byte(cardJSON),&allCards)

	cardSet = make(map [int][]Card)

	for _, card := range allCards {

		value, _ := cardSet[card.ID]

		value = append(value, card)

		cardSet[card.ID] = value
	}

	//fmt.Println(fmt.Sprintf("\n       %s       \n", "Customer Credit Cards"))

	//for i, set := range cardSet {
		//if !runningTest { fmt.Println(i, set) }
	//}
}

// Customers

func loadPeople(runningTests bool) {

	peopleJSON, _ := os.ReadFile(peopleFile)
	json.Unmarshal([]byte(peopleJSON),&people)

	for p, person := range people {

		person.Home = customerAddress[person.ID]

		people[p].Home = customerAddress[person.ID]

		if runningTests { continue }

		//fmt.Println("Person:", person)
		//fmt.Println("Person ID:", person.ID)

		//fmt.Println("Address:", customerAddress[person.ID])
		//for i, address := range customerAddress {
			//fmt.Println("All Adresses:", i, address)
		//}
	}
}

// Product Catalog

func loadItems(runningTests bool) {

	itemJSON, _ := os.ReadFile(itemsFile)
	json.Unmarshal([]byte(itemJSON),&items)
}

func main() {

	loadAddresses(false)
	loadCards(false)
	loadPeople(false)
	loadItems(false)

	inDevelopment := true

	if inDevelopment {

		fmt.Println(fmt.Sprintf("\n       %s       \n", "Customers"))

		for _, c := range people { fmt.Println(c) }

		fmt.Println(fmt.Sprintf("\n       %s       \n", "Customer's Credit Cards"))

		for _, card := range allCards { fmt.Println(card) }

		fmt.Println(fmt.Sprintf("\n       %s       \n", "Product Catalog"))

		for _, item := range items { fmt.Println(item) }

		fmt.Println(fmt.Sprintf("\n       %s       \n", "Customer Addresses"))

		for _, adress := range addresses { fmt.Println(adress) }
	}

	createOrder(people[0], 1, items[3], items[6], items[1])
}

//  Extract the data and time only
//  Return a slide with date, time and am/pm

func timeAndDate (t time.Time) []string {

    formatted := fmt.Sprintf(t.Format("2006-01-02 03:04 pm"))

    return strings.Split(formatted," ")
}


func justDate (t time.Time) string {

    return timeAndDate(t)[0]		// Date is first slice element
}

func justTime (t time.Time) string {

    time := timeAndDate(t)		// Time is second slice element

    timeOnly := time[1] + " " + time[2]	// Include the am/pm

    return timeOnly
}
/*
func printOrderSummary(order Order, items []Item, timeWhen time.Time, dateWhen time.Time, how Card) {

//	Include leading spaces; blank line as a "Header" to the Order

	fmt.Println(fmt.Sprintf("     %s\n\n", order.Name))
	fmt.Println(fmt.Sprintf("     %s %s\n\n", justDate(timeWhen), justTime(dateWhen)))

	for _, item := range items {

		fmt.Println(item)
	}

	fmt.Println(fmt.Sprintf("\n     %s\n\n", how))
	
	return
}
*/
/*
type Order struct {
    Name	string
    Items	[]Item
    OrderDate	time.Time
    ShipDate	time.Time
    ShipMethod	string
}
*/

func createOrder(customer Person, selectedCard int, itemList ...Item) {

	order := Order{}

	orderDate := time.Now()
	shipDate := orderDate.Add(time.Hour * 72)	// Est. Ship Date, in 3 days


	order.Name = customer.First + customer.Last

	order.Items = itemList

	wallet := cardSet[customer.ID]

	cardDetails := wallet[selectedCard]

	order.OrderDate = orderDate
	order.ShipDate = shipDate

	cart := []string{}

	for _, item := range order.Items {

		purchase := fmt.Sprintf("%s: $%d ", item.Name, item.Price)

		cart = append(cart, purchase)
	}

	fmt.Println(fmt.Sprintf("\n       %s       \n", "Customer Order"))


	fmt.Println("Name:", order.Name)
	fmt.Println("Address:", customer.Home)
	fmt.Println("Items:", cart)
	fmt.Println("Order Date:", justDate(orderDate))
	fmt.Println("Est. Ship Date:", justDate(shipDate))
	fmt.Println("Credit Card:",cardDetails)
}

