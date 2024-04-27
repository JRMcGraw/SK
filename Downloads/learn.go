

package main

        import "fmt"
        import "math"

func main() {


//	learn1()
//	learn2()
//	learn3()
//	learn4()
//	learn5()
//	fmt.Println(learn6(1,2,"Hello"))
//	learn7(8)
	play1()

}

type geometry interface {

	area() float64
	perimeter() float64
}


type circle struct {

	radius  float64
}


type rect struct {

	length  float64
	width   float64
}

func (c circle) area() float64 {

	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {

	return math.Pi * c.radius * 2
}

func (s rect) area() float64 {

	return s.length * s.width
}

func (s rect) perimeter() float64 {

	return 2 * (s.length + s.width)
}

func measure(g geometry) {

	fmt.Println(g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.perimeter())
}

func learn1() {

	s1 := rect{ length:4, width:6 }
	s2 := circle{ radius:5 }

	fmt.Println("Area:", s1.area())
	fmt.Println("Perimeter:", s1.perimeter())

	fmt.Println("Area:", s2.area())
	fmt.Println("Circumfersnce:", s2.perimeter())

	measure(s1)
	measure(s2)


}


/*
	pop()
	append() -- push(item)
	print_top()

*/


func push (item int) {

	ls = append(ls,item)

	return 
}

func pop() int {


	top := len(ls)-1

	item := ls[top]

	ls = ls[0:top]
	
	return item

}

func print_top() {

	top := len(ls)-1

	fmt.Println("Top:", ls[top])

	return

}


var ls []int


func learn2() {

//	Pre-load stack

	ls = []int{}

	for i := 1; i <  11; i++ {

		ls = append(ls,i) }
	fmt.Println(ls)

	b1 := []int{11,12,13}

	size := len(b1)

	for i := 0; i < size; i++ {

		push(b1[i])	

	}

	fmt.Println(b1)
	fmt.Println(ls)


	for j := 1; j <5; j++ {

		fmt.Println(pop())
		fmt.Println(ls)
	
	}
	print_top()

}


func swap(a,b int) []int {

	b, a = a, b

	return []int{a,b}
}

func learn3() {

	for i := 1 ; i < 5; i++ {

		for j := 1; j < 5; j++ {

			fmt.Println(i,j,swap(i,j))
		}
	}
}

func learn4() {

	s := []int{}

	for i := 1 ; i < 12; i++ {

		s = append(s, i)
	}

	fmt.Println(s)

	fmt.Println(reverse(s))
}

func reverse(s []int) []int {

	for back, forth := len(s) - 1, 0; (back - forth) > 0; back, forth = back-1, forth+1 {

		s[back], s[forth] =  s[forth], s[back]

	}
	return s
}

func learn5() {

	test_slice := []int{5,2}

	fmt.Println(test_slice, zero_slice(test_slice))

	test_slice = []int{}

	fmt.Println(test_slice, zero_slice(test_slice))
}

func zero_slice(i []int) bool {

	if len(i) > 0 { return false }

	return true

}

/*
*/

func learn6(a, b int, s1 string) string {

	nope := fmt.Sprintf("%d, %d, %s", a, b, s1)
	return nope
}



func learn7(last int) {

	sum := 0

	result := make(chan []int)

	for i := 1; i <= last; i++ {

		fmt.Println("Launched:", i)
		go do_square(result, i)
	}

	for i := 1; i <= last; i++ {

		one_sq :=  <- result

		sum = sum + one_sq[1]

		fmt.Println(one_sq[0], sum)
	}
	fmt.Println(last, sum)

}

func do_square(one_num chan []int, num int) {


	sq := num * num

	done := []int{num, sq}

	fmt.Println("Squared:", done)

	one_num <- done
}


type	Dog	struct {
	name	string
	speed	int			// Where he uses the bathroom
	owner	map[string]int		// Who, for how long
}

type	Tree	struct {
	name	string
	age	int
}
	
type 	Person 	struct {
	name	string
	age	int
	home	address
	pets	map[string]string
	cities	[]string
}

type	thing	interface {
	going() int
}

type	address struct {
	street	string
	city	string
	state	string
	zip	string
}

func (d Dog) going() int {
	fmt.Println("I am a Dog:", d)
	return d.speed
}

func (p Person) going() int {
	fmt.Println("I am a Person:", p)
	return p.age
}

func (t Tree) going() int {
	fmt.Println("I am a Tree:", t)
	return t.age
}

func play1() {

	dexter	:= Dog{name:"Baxter", speed:15}
kaden	:= Person{name:"Kaden Myers",age:9,home:address{state:"Florida",city:"St. Cloud"},cities:[]string{"Orlando","West Palm","Katie"},pets:map[string]string{"Dexter":"dog", "Luna":"cat", "Mercury Fat Rat":"Rat"}}
	oak	:= Tree{name:"Live", age:18}

	fmt.Println(dexter.going())
	fmt.Println(kaden.going())
	fmt.Println(oak.going())
}



