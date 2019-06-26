package otus_lesson4

import (
	"strconv"
	"testing"
)

func TestFindMax(t *testing.T) {
	myslice := convertIntSlices([]int{9, 3, 66, 8})

	result, err := FindMax(myslice, func(i, j int) bool {
		return myslice[i].(int) > myslice[j].(int)
	})

	if err != nil {
		t.Fatal(err)
	}

	if result.(int) != 66 {
		t.Fatal("Result != 66. Actual value: " + strconv.Itoa(result.(int)))
	}

	myslice = convertStringSlices([]string{"9", "3", "667", "8"})

	result, err = FindMax(myslice, func(i, j int) bool {
		return len(myslice[i].(string)) > len(myslice[j].(string))
	})

	if err != nil {
		t.Fatal(err)
	}

	if result.(string) != "667" {
		t.Fatal("Result != 667. Actual: " + result.(string))
	}
}

func convertIntSlices(slice []int) []interface{} {
	result := make([]interface{}, len(slice))
	for idx := range slice {
		result[idx] = slice[idx]
	}
	return result
}

func convertStringSlices(slice []string) []interface{} {
	result := make([]interface{}, len(slice))
	for idx := range slice {
		result[idx] = slice[idx]
	}
	return result
}