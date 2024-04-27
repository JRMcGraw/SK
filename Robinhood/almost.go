package main

import (
	"fmt"
//	"sort"
)

func main() {

	fewestMeals := solution(1, 1, []int{1,1,1})

	//fewestMeals := solution(1, 2, []int{1,4})
	fmt.Println(fewestMeals)
}

/* 

-food items have values, use values as items
-put items on plate
-needs min items
-plate can only hold max-min (range) values
-if full, new plate, or keep adding

-can't fill a plate, fail
-not enough items
-too much value (full)

*/


func solution (maxRange, minPlate int, allItems []int) int {

//  sort.Ints(allItems)

    meals := 0
    plateSize := 0

    lowItem := allItems[0]

    itemRange := 0

    if len(allItems) < minPlate { return -1 }

    for _, nextItem := range allItems {

	itemRange = nextItem - lowItem
	isFull := itemRange > maxRange

	if plateSize < minPlate {

		if isFull { return -1 }

		plateSize++

	} else { 

		if isFull {

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
