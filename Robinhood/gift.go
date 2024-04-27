package main

import (
	"fmt"
)

type pair struct {
	src	int
	dest	int
}

var target int

var allPaths [][]int{}

func main() {

	endPoints := []pair{{1,2},{1,3},{1,4},{2,5},{2,6},{2,7},{2,8},{3,9},{3,10},{4,10},
	{4,12},{4,13},{6,14},{8,15},{9,15},{10,9},{10,11},{10,12},{10,13},{10,14},
	{10,16},{13,14},{13,14},{6,11},{5,6},{7,12},
	{4,16},{15,16},{14,16},{6,16},{1,15},{7,16},
	{5,10},{5,8},{5,12},{5,6},{5,13},{5,7}}
	
	start := 1
	target = 16

	endPointSets := make(map[int][]int)

	load(endPoints, endPointSets)

	for starts, path := range endPointSets{

		fmt.Println(starts, path)
	}
	//fmt.Println("Start:", start, "Edges:", endPointSets[start], "Target:", target)

	path := []int{}
	path = append(path, start)

	fmt.Println(path)

	found := false


	for _, endPoint := range endPointSets[start] {
	
		fmt.Println("Start:", endPoint, "Endpoints:", endPointSets[endPoint], "Target:", target)

		path, found = walk(endPoint, endPointSets[endPoint], path, endPointSets)

		fmt.Println("Walked:", path, len(path), found)

		allPaths = append(allPaths, path)

		path = []int{}
	}

	for _, path := range allPaths {

		fmt.Println("Path:", path, len(path))
	}
}

var found bool

func walk(src int, endPointSet []int, path []int, endPointSets map[int][]int) ([]int, bool) { 

	fmt.Println("Walked so far:", path, "Next:", src, endPointSet, target)


	for i, dest := range endPointSet {  // each dest

		fmt.Println("In Range:", i, dest, endPointSet)

		if dest == target {

			path = append(path,target)

			fmt.Println("Found a", lne(path), "step path:", path)

			continue
		}

		src = dest
		
		// Add src to path

		path = append(path,src)

		//fmt.Println("Path before walk(", i, ")", path)

		path, found = walk(src, endPointSets[src], path, endPointSets)

		if found {

			allPaths = append(allPaths, path)

			path = []int{}
		}

	}

	fmt.Println("No more endpoints")
	return path, false
}


func load(pairs []pair, endPointSets map[int][]int) {


	for _, pair := range pairs {

		_, ok := endPointSets[pair.src] 

		// If empty set

		if !ok {

			// New set of Edges

			endPointSets[pair.src] = []int{}
		}
		// Add dest to set of Edges

		endPointSets[pair.src] = append(endPointSets[pair.src], pair.dest)

	}
	//fmt.Println(endPointSets)
}

