// +build unit

package main

import (
	"testing"
	"io/ioutil"
	"io"
	"net/http"
	"net/http/httptest"
	"fmt"
)

type AddResult struct {
	x        int
	y        int
	expected int
}

var addResults = []AddResult{
	{1, 1, 2},
	{2, 2, 4},
}

func TestAdd(t *testing.T) {
	for _, test := range addResults {
		result := Add(test.x, test.y)
		if result != test.expected {
			t.Fatal("Expected", test.x, "plus", test.y, "to equal", test.expected)
		}
	}
}

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("expected 2 + 2 to equal 4")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{99999, 100001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}

func TestReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("test-data/test.data")
	if err != nil {
		t.Fatal("Could not open file")
	}
	if string(data) != "hello world" {
		t.Fatal("String contents do not match expected")
	}
}

func TestHttpRequest(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{ \"status\": \"good\" }")
	}

	req := httptest.NewRequest("GET", "https://tutorialedge.net", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if 200 != resp.StatusCode {
		t.Fatal("Status Code Not OK")
	}
}

// 9 minutes in https://www.youtube.com/watch?v=S1O0XI0scOM