package main

import (
	"fmt"
)

func main() {

	nutritionValues := []int{1,4}

	fewestMeals := solution(1, 2, nutritionValues)

	fmt.Println(fewestMeals)
}

func solution (maxNutRange, minPlate int, allItems []int) int {

    totalMeals := 0
    plateSize := 0

    lowNutrient := allItems[0]

    nutRange := 0

    // Can we fill a plate?

    if len(allItems) < minPlate { return -5 }

//  Treat each Nutritional Value as an Meal Item
//  Start loading the first plate (meal)

    for _, nextItem := range allItems {	// Next Meal Item is also its Nutrional Value

	nutRange = nextItem - lowNutrient

	// Get a min. plate full first

	if plateSize < minPlate {

		// No.  Try to add nextItem

		if nutRange > maxNutRange {	// Too much nutrition for partial plate

			return -1
		}

		plateSize++ 			// Add an item

	} else { // Met. min. items

		if nutRange > maxNutRange {	// Nutrition range exceeded; new Meeal

			totalMeals++
			plateSize = 1

			lowNutrient = nextItem	// New Min. (first meal item)

		} else {	// Keep adding items

			plateSize++ 		// Add an item

		}

	}
    }

    totalMeals++

    return totalMeals
}
