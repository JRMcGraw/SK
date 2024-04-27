package main

import (
	"fmt"
	"time"
	"testing"
)


type testPerson struct {
    First	string
    Last	string
    ID		int
    Phone	string
    Email	string
}

var testDataLoadPeople = []testPerson{

	{  "JR", "McGraw", 1, "(512) 800-0124", "jrmcgraw@alum.mit.edu" },
	{  "Kaden", "Myers", 2, "(561) 427-8629", "kaden.myers@gmail.com" },
}

func TestLoadPeople(t *testing.T) {

	loadPeople(true)

	testFailed := false

	for p, person := range people {

		if person.First != testDataLoadPeople[p].First {

			testFailed = true
			break
		}
		if person.Last != testDataLoadPeople[p].Last {

			testFailed = true
			break
		}
		if person.ID != testDataLoadPeople[p].ID {

			testFailed = true
			break
		}
		if person.Phone != testDataLoadPeople[p].Phone {

			testFailed = true
			break
		}
		if person.Email != testDataLoadPeople[p].Email {

			testFailed = true
			break
		}
	}
	if testFailed { 
	
		t.Error() 
	} else {
		testResults := fmt.Sprintf("ALL (%d) Load People Tests PASSED.", len(people))
		fmt.Println(testResults)
	}
}

var testDataLoadAddress = []Address{ 

	{ 1, "2553 Lakewalk Cove", "Apt 109", "Casselberry", "FL", "32707", "USA"},
	{ 2, "16 Bayshore Drive", "", "St. Cloud", "FL", "33812", "USA"},
}

func TestLoadAddresses(t *testing.T) {

	loadAddresses(true)

	testFailed := false

	for _, address := range testDataLoadAddress {

		if address != customerAddress[address.ID] {

			fmt.Println("Load Address Test Failed.")
			fmt.Println("Expeted:", testDataLoadAddress)
			fmt.Println("Loaded:", address)

			testFailed = true
		}

	}

	if testFailed { 
	
		t.Error() 
	} else {

		testResults := fmt.Sprintf("ALL (%d) Load Address Tests PASSED.", len(testDataLoadAddress))
		fmt.Println(testResults)
	}
}

var testDataLoadItems = []Item{ 

	{"Laptop", 4255, 6 },
	{"Mouse", 12, 12 },
	{"Monitor", 188, 4 },
	{"Keyboard", 18, 11 },
	{"Headphones", 28, 32 },
	{"Desk", 350, 1 },
	{"Chair", 88, 2 },
}

func TestLoadItems(t *testing.T) {

	loadItems(true)

	testFailed := false

	for i, testItem := range testDataLoadItems {

		if items[i] != testItem {

			fmt.Println("Load Items Test Failed.")
			fmt.Println("Expeted:", testItem)
			fmt.Println("Loaded:", items[i])

			testFailed = true
		}

	}

	if testFailed { 
	
		t.Error() 
		
	} else {

		testResults := fmt.Sprintf("ALL (%d) Load Items Tests PASSED.", len(testDataLoadItems))
		fmt.Println(testResults)
	}
}


var testDataLoadCards = map[int][]Card{ 

	3:[]Card{{ 3, "AMEX", "3247 885307 15669", "05/27", "5306" },
		 { 3, "Visa", "3814 2829 2648 0773", "12/25", "712" },
		 { 3, "MasterCard", "3633 6534 7335 8744", "02/29", "770" },
	},

	1:[]Card{{ 1, "Visa", "3817 8721 2612 7013", "12/26", "192" },
		 { 1, "Visa", "3988 2319 2742 1073", "11/27", "597" },
		 { 1, "AMIX", "8914 2829 2142 0701", "12/26", "312" },
	},

	2:[]Card{{ 2, "Visa Debit", "3914 2329 2642 0003", "12/26", "392" },
	},
}


func TestLoadCards(t *testing.T) {

	loadCards(true)

	testFailed := false

	for id, customerSet := range cardSet {

		testSet := testDataLoadCards[id]

		if compareCardSlices(customerSet, testSet) == false {

			fmt.Println("Load Cards Test Failed.")
			fmt.Println("Expeted Set:", testSet)
			fmt.Println("Loaded Set:", customerSet)

			testFailed = true

		} else {
			result := fmt.Sprintf("Load Cards Test (%d) PASSED: %v", id, customerSet)
			fmt.Println(result)
		}
	}

	if testFailed {

		t.Error()

	} else {

		result := fmt.Sprintf("ALL (%d) Load Cards Tests PASSED.", len(cardSet))
		fmt.Println(result)
	}
}

func compareCardSlices(cardSlice1, cardSlice2 []Card) bool {

	if len(cardSlice1) != len(cardSlice2) { return false }

	for i := range cardSlice1 {

		if cardSlice1[i] != cardSlice2[i] { return false }
	}
	return true
}

// Test both justDate() and justTime()

// Fixtures

type haveData struct {
	have	int		// Switch parameter to get to the correct time.Date
	want	wantValues
}

type wantValues struct {
	date	string
	time	string
}

var check = []haveData{ { 0, wantValues{"2023-08-06","02:00 am"} },
			{ 1, wantValues{"2023-01-11","12:00 pm"} },
			{ 2, wantValues{"2023-02-06","11:00 am"} },
			{ 3, wantValues{"2023-11-30","06:00 pm"} },
			{ 4, wantValues{"2023-12-24","02:00 pm"} },
		      }

func getHaveData(have int) time.Time {

	t := time.Now()

	switch have {
	case 0: t = time.Date(2023, 8, 6, 2, 0, 0, 0, time.Local)
	case 1: t = time.Date(2023, 1, 11, 12, 0, 0, 0, time.Local)
	case 2: t = time.Date(2023, 2, 6, 11, 0, 0, 0, time.Local)
	case 3: t = time.Date(2023, 11, 30, 18, 0, 0, 0, time.Local)
	case 4: t = time.Date(2023, 12, 24, 14, 0, 0, 0, time.Local)
	}

	return t
}

func TestJustBoth(t *testing.T) {

	fail := false
	result := []string{}

	testType := "Just Date/Time"

	for _, c := range check {

		t := getHaveData(c.have)

		want := c.want

		dateOnly := justDate(t)
		timeOnly := justTime(t)

		if want.date != dateOnly { fail = true }

		if want.time != timeOnly { fail = true }

		if fail {

			reason := fmt.Sprintf("%s failed: [%s, %s], [%s %s]", testType,
			want.date, dateOnly, want.time, timeOnly)
			result = append(result, reason)

		} else {
			passDetail := fmt.Sprintf("%s (%d) PASSED: [%v], [%s %s]", testType, c.have,
			t, want.date, want.time )

			fmt.Println(passDetail)
		}
	}

	if fail {
		fmt.Println(result)

		t.Error()
		return
	}
	testResults := fmt.Sprintf("ALL (%d) Just Date/Time Tests PASSED.", len(check))
	fmt.Println(testResults)
}
