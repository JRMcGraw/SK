package main
import "fmt"
import "sort"


func uniqueNames(a, b []string) []string {
    var result []string

	both := []string{}

	fmt.Println(a)
	fmt.Println(b)
	//fmt.Println(both)

	both = append(a, b...)
	sort.Strings(both)

	fmt.Println(both)

	last := ""
	for _, name := range both {


		if name != last {

			result = append(result,name)
			//fmt.Println(result)

			last = name
		}
	}

	return result
}


func main() {

// should print Ava, Emma, Olivia, Sophia

	fmt.Println(uniqueNames(
	[]string{"Ava", "Emma", "Olivia"}, 
	[]string{"Olivia", "Sophia", "Emma"}))  
}

