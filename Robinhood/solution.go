package main

import (
	"fmt"
//	"sort"
)

func main() {

	//nutritionValues := []int{1,3,5,6,8,10,18,23,25,26,28,40,58}
	//nutritionValues := []int{1,1,1}
	nutritionValues := []int{1,4}
//	minPlate := 2
//	maxPlateNutrition := 5

	fmt.Println(nutritionValues)

//	fewestMeals := solution(maxPlateNutrition, minPlate, nutritionValues)
	//fewestMeals := solution(8, 3, nutritionValues)
	//fewestMeals := solution(1, 1, nutritionValues)
	fewestMeals := solution(1, 2, nutritionValues)

	fmt.Println("We Done.")
	fmt.Println(fewestMeals)
	fmt.Println("We Really Done.")
}

func solution (maxMealNutrition, minMealItems int, allItems []int) int {

	allPlates := make(map[int][]int)

	currentPlate := []int{}

	fmt.Println(allPlates)
	fmt.Println(currentPlate)

//    sort.Ints(allItems)

    fmt.Println(allItems)

 //   return 0

    totalMeals := 0
    mealItems := 0

    minMealNutrition := allItems[0]

    nutritionRange := 0

    // Can we fill a plate?

    if len(allItems) < minMealItems { return -5 }

  //  return 0

//  Treat each Nutritional Value as an Meal Item
//  Start loading the first plate (meal)

currentPlate = []int{}
    for _, nextItem := range allItems {	// Next Meal Item is also its Nutrional Value

	    fmt.Println("Next Item:", nextItem)
	nutritionRange = nextItem - minMealNutrition
	    fmt.Println("Range:", nutritionRange)

	// Get a min. plate full first
fmt.Println("Meal Items:", mealItems)

	if mealItems < minMealItems {

		// No.  Try to add nextItem

		fmt.Println("Partial Plate:", mealItems, minMealItems)

		fmt.Println(currentPlate)

		if nutritionRange > maxMealNutrition {	// Too much nutrition for partial plate

		fmt.Println("Range > Max", nutritionRange,maxMealNutrition)
fmt.Println(allPlates)
			return -1
		}

		currentPlate = append(currentPlate, nextItem)

		mealItems++ 			// Add an item

	} else { // Met. min. items

		fmt.Println("Plate Size:", mealItems)

		if nutritionRange > maxMealNutrition {	// Nutrition range exceeded; new Meeal

			totalMeals++
			mealItems = 1

		fmt.Println("New Meal:", totalMeals)

	fmt.Println(currentPlate)
	allPlates[totalMeals] = currentPlate
	currentPlate = []int{}
	currentPlate = append(currentPlate, nextItem)


			minMealNutrition = nextItem	// New Min. (first meal item)

		} else {	// Keep adding items

			mealItems++ 		// Add an item

		currentPlate = append(currentPlate, nextItem)

		fmt.Println("Added:", mealItems)


			//latestNutrition = nextItem
		}

	}
	fmt.Println("Total Meals:", totalMeals)
	fmt.Println("Total Items:", mealItems)
    }

fmt.Println(allPlates)

    // Is it okay if we coun't even fill one plate?

    totalMeals++

    return totalMeals

}
