package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func do_fizzbuzz(int1 uint, int2 uint, limit uint, str1 string, str2 string) string {

	fizzbuzz := fmt.Sprintf(" %s%s,", str1, str2)
	fizz := fmt.Sprintf(" %s,", str1)
	buzz := fmt.Sprintf(" %s,", str2)
	fizzbuzzString := ""
	for i := uint(1); i < limit; i++ {
		if i%(int1*int2) == 0 {
			fizzbuzzString += fizzbuzz
		} else if i%int1 == 0 {
			fizzbuzzString += fizz
		} else if i%int2 == 0 {
			fizzbuzzString += buzz
		} else {
			fizzbuzzString += fmt.Sprintf(" %d,", i)
		}
	}
	return fizzbuzzString
}

func FizzBuzz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", do_fizzbuzz(3, 5, 100, "fizz", "buzz"))
}

func main() {
	s := do_fizzbuzz(3, 5, 100, "fizz", "buzz")
	fmt.Println(s)
	r := mux.NewRouter()
	r.HandleFunc("/", FizzBuzz)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
