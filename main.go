package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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

	sInt1, ok := r.URL.Query()["int1"]
	if !ok || len(sInt1[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Url Param 'int1' is missing")
		fmt.Fprintf(w, "Url Param 'int1' is missing")
		return
	}
	sInt2, ok := r.URL.Query()["int2"]
	if !ok || len(sInt2[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Url Param 'int2' is missing")
		fmt.Fprintf(w, "Url Param 'int2' is missing")
		return
	}
	sLimit, ok := r.URL.Query()["limit"]
	if !ok || len(sLimit[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Url Param 'limit' is missing")
		fmt.Fprintf(w, "Url Param 'limit' is missing")
		return
	}

	str1, ok := r.URL.Query()["str1"]
	if !ok || len(str1[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Url Param 'str1' is missing")
		fmt.Fprintf(w, "Url Param 'str1' is missing")
		return
	}
	str2, ok := r.URL.Query()["str2"]
	if !ok || len(str2[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Url Param 'str2' is missing")
		fmt.Fprintf(w, "Url Param 'str2' is missing")
		return
	}

	w.WriteHeader(http.StatusOK)
	if int1, err := strconv.ParseUint(sInt1[0], 10, 32); err == nil {
		if int2, err := strconv.ParseUint(sInt2[0], 10, 32); err == nil {
			if limit, err := strconv.ParseUint(sLimit[0], 10, 32); err == nil {
				fmt.Fprintf(w, "%s", do_fizzbuzz(uint(int1), uint(int2), uint(limit), str1[0], str2[0]))
			}
		}
	}
}

func main() {
	s := do_fizzbuzz(3, 5, 100, "fizz", "buzz")
	fmt.Println(s)
	r := mux.NewRouter()
	r.HandleFunc("/", FizzBuzz)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
