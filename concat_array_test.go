package datautil

import (
	"reflect"
	"testing"
)

func TestConcatArrayInts(t *testing.T) {
	a := []int{1, 2}
	b := []int{3, 4}
	expected := []int{1, 2, 3, 4}
	result := ConcatArray(a, b)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestConcatArrayStrings(t *testing.T) {
	a := []string{"a", "b"}
	b := []string{"c"}
	expected := []string{"a", "b", "c"}
	result := ConcatArray(a, b)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestConcatArrayEmpty(t *testing.T) {
	a := []int{}
	b := []int{1, 2}
	expected := []int{1, 2}
	result := ConcatArray(a, b)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestConcatArrayFloats(t *testing.T) {
	a := []float64{1.1, 2.2}
	b := []float64{3.3}
	expected := []float64{1.1, 2.2, 3.3}
	result := ConcatArray(a, b)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestConcatArrayFromArrays(t *testing.T) {
	arr1 := [2]int{1, 2}
	arr2 := [3]int{3, 4, 5}
	a := arr1[:]
	b := arr2[:]
	expected := []int{1, 2, 3, 4, 5}
	result := ConcatArray(a, b)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestConcatArrayBothEmpty(t *testing.T) {
	a := []int{}
	b := []int{}
	expected := []int{}
	result := ConcatArray(a, b)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
