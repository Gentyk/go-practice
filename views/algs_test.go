package views

import (
	"testing"
	"fmt"
)

func TestSwapString(t *testing.T) {
	testCases := []struct {
		startString    string
		expectedString string
	}{
		{"", ""},
		{"a", "a"},
		{"фаt12", "21tаф"},
		{"3фаt12", "21tаф3"},
	}

	for ind, test := range testCases {
		t.Run("test swap string", func(t *testing.T) {
			result := swapString(test.startString)
			if result != test.expectedString {
				t.Errorf("got %s; want %s", result, test.expectedString)
			} else {
				fmt.Printf("swap: success %d case\n", ind)
			}
		})
	}
}

func TestFib(t *testing.T) {
	testCases := []struct {
		n	int
		expected int
	} {
		{0, 0},
		{2, 2},
		{-1, 0},
		{4, 24},
	}

	for ind, testCase := range testCases {
		t.Run("fib", func(t *testing.T) {
			result := fib(testCase.n)
			if result != testCase.expected {
				t.Errorf("get %d, expected %d", result, testCase.expected)
			} else {
				fmt.Printf("fib: success case %d \n", ind)
			}
		})
	}
}

func TestFib2(t *testing.T) {
	testCases := []struct {
		n	int
		expected int
	} {
		{0, 0},
		{2, 2},
		{-1, 0},
		{4, 24},
	}

	for ind, testCase := range testCases {
		t.Run("fib2", func(t *testing.T) {
			result := fib2(testCase.n)
			if result != testCase.expected {
				t.Errorf("get %d, expected %d", result, testCase.expected)
			} else {
				fmt.Printf("fib2: success case %d \n", ind)
			}
		})
	}
}
