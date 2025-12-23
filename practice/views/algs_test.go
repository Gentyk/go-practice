package views

import (
	"fmt"
	"reflect"
	"testing"
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
		n        int
		expected int
	}{
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
		n        int
		expected int
	}{
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

func TestBubleSort(t *testing.T) {
	testCases := []struct {
		startArray    []int
		expectedArray []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{4, 2, 3, 1}, []int{1, 2, 3, 4}},
	}

	for ind, tc := range testCases {
		t.Run("bubleSort", func(t *testing.T) {
			result := tc.startArray
			bumbleSort(result)
			if reflect.DeepEqual(result, tc.expectedArray) {
				fmt.Printf("bubleSort: success case %d \n", ind)
			} else {
				t.Errorf("get %d, expected %d", result, tc.expectedArray)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		startArray    []int
		expectedArray []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{4, 2, 3, 1}, []int{1, 2, 3, 4}},
		{[]int{4, 5, 2, 7, 3, 0, 1}, []int{0, 1, 2, 3, 4, 5, 7}},
	}

	for ind, tc := range testCases {
		t.Run("quickSort", func(t *testing.T) {
			result := tc.startArray
			quickSort(result, 0, len(result)-1)
			if reflect.DeepEqual(result, tc.expectedArray) {
				fmt.Printf("quickSort: success case %d \n", ind)
			} else {
				t.Errorf("get %d, expected %d", result, tc.expectedArray)
			}
		})
	}
}
