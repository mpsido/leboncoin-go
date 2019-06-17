package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestFizzbuzz(t *testing.T) {
	s := do_fizzbuzz(3, 5, 100, "fizz", "buzz")
	if s != " 1, 2, fizz, 4, buzz, fizz, 7, 8, fizz, buzz, 11, fizz, 13, 14, fizzbuzz, 16, 17, fizz, 19, buzz, fizz, 22, 23, fizz, buzz, 26, fizz, 28, 29, fizzbuzz, 31, 32, fizz, 34, buzz, fizz, 37, 38, fizz, buzz, 41, fizz, 43, 44, fizzbuzz, 46, 47, fizz, 49, buzz, fizz, 52, 53, fizz, buzz, 56, fizz, 58, 59, fizzbuzz, 61, 62, fizz, 64, buzz, fizz, 67, 68, fizz, buzz, 71, fizz, 73, 74, fizzbuzz, 76, 77, fizz, 79, buzz, fizz, 82, 83, fizz, buzz, 86, fizz, 88, 89, fizzbuzz, 91, 92, fizz, 94, buzz, fizz, 97, 98, fizz," {
		t.Fail()
	}

	s = do_fizzbuzz(3, 5, 18, "fizz", "buzz")
	if s != " 1, 2, fizz, 4, buzz, fizz, 7, 8, fizz, buzz, 11, fizz, 13, 14, fizzbuzz, 16, 17," {
		t.Fail()
	}
}

func TestEndpoint(t *testing.T) {
	statMap = make(map[string]int)
	router := mux.NewRouter()
	router.HandleFunc("/", FizzBuzz)
	router.HandleFunc("/stats", Stats)

	{
		rr := httptest.NewRecorder()
		path := fmt.Sprintf("/stats")
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatal("application returned an error")
		}
		if rr.Body.String() != "No stats available" {
			t.Fatal(rr.Body.String())
		}
	}
	{
		rr := httptest.NewRecorder()
		path := fmt.Sprintf("/?int1=%d&int2=%d&limit=%d&string1=%s&string2=%s", 3, 5, 18, "fizz", "buzz")
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}
		q := req.URL.Query()
		q.Add("int1", "3")
		q.Add("int2", "5")
		q.Add("limit", "18")
		q.Add("str1", "fizz")
		q.Add("str2", "buzz")
		req.URL.RawQuery = q.Encode()

		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatal("application returned an error")
		}
		if rr.Body.String() != " 1, 2, fizz, 4, buzz, fizz, 7, 8, fizz, buzz, 11, fizz, 13, 14, fizzbuzz, 16, 17," {
			t.Fatal(rr.Body.String())
		}
	}
	{
		rr := httptest.NewRecorder()
		path := fmt.Sprintf("/stats")
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatal("application returned an error")
		}
		if rr.Body.String() != fmt.Sprintf("Most used query is %s\nUsed %d times", fmt.Sprintf("?int1=%d&int2=%d&limit=%d&string1=%s&string2=%s", 3, 5, 18, "fizz", "buzz"), 1) {
			t.Fatal(rr.Body.String())
		}
	}
}

func TestEndpointErrors(t *testing.T) {
	statMap = make(map[string]int)
	router := mux.NewRouter()
	router.HandleFunc("/", FizzBuzz)
	router.HandleFunc("/stats", Stats)

	// test int1 missing
	{
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		q := req.URL.Query()
		q.Add("int2", "5")
		q.Add("limit", "18")
		q.Add("str1", "fizz")
		q.Add("str2", "buzz")
		req.URL.RawQuery = q.Encode()

		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Fatal("application should have returned an error")
		}
		if rr.Body.String() != "Url Param 'int1' is missing" {
			t.Fatal(rr.Body.String())
		}
	}
	// test int2 missing
	{
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		q := req.URL.Query()
		q.Add("int1", "3")
		q.Add("limit", "18")
		q.Add("str1", "fizz")
		q.Add("str2", "buzz")
		req.URL.RawQuery = q.Encode()

		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Fatal("application should have returned an error")
		}
		if rr.Body.String() != "Url Param 'int2' is missing" {
			t.Fatal(rr.Body.String())
		}
	}
	// if many query parameters are missing only warn about the first encountered missing parameter
	{
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		q := req.URL.Query()
		q.Add("int1", "3")
		req.URL.RawQuery = q.Encode()

		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Fatal("application should have returned an error")
		}
		if rr.Body.String() != "Url Param 'int2' is missing" {
			t.Fatal(rr.Body.String())
		}
	}
	// test str1 missing
	{
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		q := req.URL.Query()
		q.Add("int1", "3")
		q.Add("int2", "5")
		q.Add("limit", "18")
		q.Add("str2", "buzz")
		req.URL.RawQuery = q.Encode()

		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Fatal("application should have returned an error")
		}
		if rr.Body.String() != "Url Param 'str1' is missing" {
			t.Fatal(rr.Body.String())
		}
	}
	// as a result no stat had been registered
	{
		rr := httptest.NewRecorder()
		path := fmt.Sprintf("/stats")
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatal("application returned an error")
		}
		if rr.Body.String() != "No stats available" {
			t.Fatal(rr.Body.String())
		}
	}
}