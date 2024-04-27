
func sayHello(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello")
	fmt.Println("Hello")

}
