package main

import (
	"fmt"
	"sort"
)

func main() {

	nutritionValues := []int{1,2,3,4,5,6,7}
//	minPlate := 2
//	maxPlateNutrition := 5

	fmt.Println(nutritionValues)

//	fewestMeals := solution(maxPlateNutrition, minPlate, nutritionValues)
	fewestMeals := solution(5, 2, nutritionValues)

	fmt.Println("We Done.")
	fmt.Println(fewestMeals)
	fmt.Println("We Really Done.")
}

func solution (maxMealNutrition, minMealItems int, allItems []int) int {


    sort.Ints(allItems)

    fmt.Println(allItems)

    return 0

    totalMeals := 0
    mealItems := 0

    minMealNutrition := allItems[0]

    nutritionRange := 0

    // Can we fill a plate?

    if len(allItems) < minMealItems { return -1 }

    return 0

//  Treat each Nutritional Value as an Meal Item
//  Start loading the first plate (meal)

    for _, nextItem := range allItems {	// Next Meal Item is also its Nutrional Value

        fmt.Println(nextItem)
	nutritionRange = nextItem - minMealNutrition

	// Get a min. plate full first

	if mealItems < minMealItems {

		// No.  Try to add nextItem


		if nutritionRange > maxMealNutrition {	// Too much nutrition for partial plate

			return -1
		}

		mealItems++ 			// Add an item

	} else { // Met. min. items

		if nutritionRange > maxMealNutrition {	// Nutrition range exceeded; new Meeal

			totalMeals++
			mealItems = 1

			minMealNutrition = nextItem	// New Min. (first meal item)

		} else {	// Keep adding items

			mealItems++ 		// Add an item

			//latestNutrition = nextItem
		}

	}
	fmt.Println("Total Meals:", totalMeals)
	fmt.Println("Total Items:", mealItems)
    }

    // Is it okay if we coun't even fill one plate?

    totalMeals++

    return totalMeals

}
