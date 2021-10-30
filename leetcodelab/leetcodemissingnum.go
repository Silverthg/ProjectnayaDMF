package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", missing_number_handler)
	http.ListenAndServe(":8080", nil)
}

func missing_number(arr []int) {
	sum := 0
	idx := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			idx = i
		} else {
			sum += arr[i]
		}
	}

	total := (len(arr) + 1) * len(arr) / 2

	fmt.Printf("missing number is %d at index %d ", total-sum, idx)
}

func missing_number_handler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	missing_number(parr(r.FormValue("key")))

}

func parr(s string) []int {
	arr := make([]int, len(s)-1)
	sa, _ := strconv.Atoi(s)
	for i := 0; i < len(s)-1; i++ {
		arr[i] = sa % 10
		sa = sa / 10

	}
	return arr
}
