package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(3, 5)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and check sum is %d", sum))
}

// 這個function不會被外部使用(private)，所以小寫開頭
func addValues(x, y int) int {
	sum := x + y
	return sum
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(":8000", nil)
}
