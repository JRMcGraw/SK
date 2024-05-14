package main

import (
	"fmt"
	"errors"
	"os"
	"strconv"
	"strings"
)

func main() {

	doWeekDays()

	A := []int{1,2,5,8,10,20}
	//A := []int{10,2,5,1,8,20}
	answer, isTriple := triples(A)

	fmt.Println("Answer:", answer, isTriple)

	tripleSetShow()
}

func doWeekDays() {

	args := os.Args

	fmt.Println(args)

//	futureDay := getNewDay("Mon", 24)
//	fmt.Println(futureDay)

	argPair, offset := 0, 0
	day := ""

	for i, arg := range args {

		if i == 0 { 

			command := strings.Split(arg, "\\")
			
			for _, directory := range command {
				fmt.Println("Directory:", directory)
			}
			size := len(command)

			fmt.Println("Command:", command[size-1])

			continue 
		}

		//fmt.Println(i, arg)

		switch argPair {
		case 0: day = arg
			argPair++

		case 1: offset, _ = strconv.Atoi(arg)
			argPair = 0


			fmt.Println("Calling with:", day, offset)
			futureDay := getNewDay(day, offset)
			fmt.Println(futureDay)

		default:  fmt.Println("Input Error:", args)
		}
	}

	//futureDay := getNewDay("Sat", 23)
	//futureDay := getNewDay("Sat", 0)
	//futureDay := getNewDay("Sat", 7)
	return
}

var dayOfWeek = []string{"Error", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

func getNewDay(dayName string, futureDay int) string {

	dayStart, _ := dayIndex(dayName)
	//if err != nil { return err.Error() }

	fullWeeks := int(futureDay/7)
	fullDays := fullWeeks * 7
	extraDays := futureDay - fullDays

	nextDay := dayStart + extraDays

	if nextDay > 7 { nextDay = nextDay - 7 }
	return dayOfWeek[nextDay] 
}

func dayIndex (dayName string) (int, error) {

	for i, day := range dayOfWeek {

		if dayName == day { return i, nil }
	}
	return 0, errors.New("Invalid Day")
}

/*	Find Triples

	Rules

	N integers

	Triplet		Any 3 from N
			0 <= P < Q < R < N

	pqr()		P + Q > R
	qrp()		Q + R > P
	rpq()		R + P > Q

			0 <= N1 < N2 < N3 < N

	Perimeter	P + Q + R
*/

type triplet struct {
	p	int
	q	int
	r	int
}

type perimeter struct {
	t	triplet
	size	int
}

func (t *triplet) pqr() bool {

	if (t.p + t.q) > t.r { return true }
	return false
}

func (t *triplet) qrp() bool {

	if (t.r + t.p) > t.q { return true }
	return false
}

func (t *triplet) rpq() bool {

	if (t.r + t.p) > t.q { return true }
	return false
}

func (p *triplet) isTriplet() bool {

	if p.pqr() && p.pqr() && p.rpq() { return true }
	return false
}


func triples(A []int) (perimeter, bool) {

	N := len(A)

	possibleAnswers := []perimeter{}
	//triplets := []triplet{}

	//0,1,2,3,4,5,6,7,8,9
	//
	//N = 10
	//The 7 is at N-3

	// Walk the sorted list of integers
	// Stopping at the last set of three

	for i, _ := range A {

		if i == N - 2 { break }

		// Create the triplet, 
		// current integer and 
		// following two integers

		t := triplet{ A[i], A[i+1], A[i+2] }
		
		if t.isTriplet() == false { continue }

		// It passes the triplet test
		// May not need this first slice

		//triplets = append(triplets, t)

		possibleAnswer := perimeter{ t, (t.p + t.q + t.r)}

		possibleAnswers = append(possibleAnswers, possibleAnswer)
	}

	if len(possibleAnswers) == 0 {

		return perimeter{}, false
	}

	// Get largest perimeter

	largest := 0

	answer := perimeter{}

	for _, t := range possibleAnswers {

		if t.size > largest {

			answer.t = t.t
			answer.size = t.size
			largest = t.size
		}
	}

	fmt.Println("Answer:", answer)

	fmt.Println(A)

	return answer, true

	args := os.Args

	fmt.Println(args)

//	futureDay := getNewDay("Mon", 24)
//	fmt.Println(futureDay)

	argPair, offset := 0, 0
	day := ""

	for i, arg := range args {

		if i == 0 { 

			command := strings.Split(arg, "\\")
			
			for _, directory := range command {
				fmt.Println("Directory:", directory)
			}
			size := len(command)

			fmt.Println("Command:", command[size-1])

			continue 
		}

		//fmt.Println(i, arg)

		switch argPair {
		case 0: day = arg
			argPair++

		case 1: offset, _ = strconv.Atoi(arg)
			argPair = 0


			fmt.Println("Calling with:", day, offset)
			futureDay := getNewDay(day, offset)
			fmt.Println(futureDay)

		default:  fmt.Println("Input Error:", args)
		}
	}

	fmt.Println(A)
	return perimeter{}, false
}



type triple2 struct {
        p, q, r int
}

func tripleSetShow() {

        tripleSet2 := []triple2{}

        //set := []int{10, 2, 5, 1, 8, 20}
	set := []int{1,2,5,8,10,20}

        //sort.Ints(set)

        end := len(set) - 2

        size := len(set)

        p, q, r := 0, 0, 0

        oneTriple := triple2{}

	if set[0] <= 0 { 

		fmt.Println("Triple Set:", tripleSet2)
		return
	}

        for i, _ := range set[:end] {

                p = set[i]
                q = set[i+1]
                r = set[i+2]

		fmt.Println(p,q,r)

		if r > size { break }

                //if size < p + q + r { continue }

		pqr := ispqr(p, q, r )
		qrp := isqrp(p, q, r )
		rpq := isrpq(p, q, r )

		if pqr || qrp || rpq { continue }

                oneTriple = triple2{p,q,r}

		fmt.Println(oneTriple)

                // Valid triple

                tripleSet2 = append(tripleSet2, oneTriple)
        }

        fmt.Println("Triple Set:", tripleSet2)


}

func ispqr(p, q, r int) bool {

	isGood := ( p + q ) > r
	fmt.Println(p,q,r, isGood)

        return ( p + q ) > r
}
func isqrp(p, q, r int) bool {

	isGood := ( q + r ) > p
	fmt.Println(p,q,r, isGood)

        return ( q + r ) > p
}
func isrpq(p, q, r int) bool {

	isGood := ( r + p ) > q
	fmt.Println(p,q,r, isGood)

        return ( r + p ) > q
}

