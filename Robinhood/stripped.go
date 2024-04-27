package main

import (
	"fmt"
//	"sort"
)

func main() {

	//fewestMeals := solution(1, 1, []int{1,1,1})

	fewestMeals := solution(1, 2, []int{1,4})
	fmt.Println(fewestMeals)
}

/* values can be food items
   
put items on plates
   add untli minimum
   enough too soon?
   add until enough
   new plate
no items?
   
 */


func solution (maxRange, minPlate int, allItems []int) int {

//  sort.Ints(allItems)

    meals := 0
    plateSize := 0

    lowItem := allItems[0]

    itemRange := 0

    if len(allItems) < minPlate { return -2 }

    for _, nextItem := range allItems {

	itemRange = nextItem - lowItem
	enough := itemRange > maxRange

	if plateSize < minPlate {

		if enough {

			return -1
		}

		plateSize++

	} else { 

		if enough {

			meals++
			plateSize = 1

			lowItem = nextItem	

		} else {

			plateSize++ 
		}
	}
    }

    meals++

    return meals
}
