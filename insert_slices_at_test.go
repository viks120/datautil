package datautil

import (
	"reflect"
	"testing"
)

func TestInsertSliceAt_Middle(t *testing.T) {
	original := []int{1, 2, 5, 6}
	toInsert := []int{3, 4}
	position := 3
	expected := []int{1, 2, 3, 4, 5, 6}

	result, err := InsertSliceAt(original, toInsert, position)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestInsertSliceAt_Start(t *testing.T) {
	original := []int{3, 4, 5}
	toInsert := []int{1, 2}
	position := 1
	expected := []int{1, 2, 3, 4, 5}

	result, err := InsertSliceAt(original, toInsert, position)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestInsertSliceAt_End(t *testing.T) {
	original := []int{1, 2, 3}
	toInsert := []int{4, 5}
	position := 4
	expected := []int{1, 2, 3, 4, 5}

	result, err := InsertSliceAt(original, toInsert, position)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestInsertSliceAt_InvalidPositionZero(t *testing.T) {
	original := []int{1, 2, 3}
	toInsert := []int{4, 5}
	position := 0

	_, err := InsertSliceAt(original, toInsert, position)
	if err == nil {
		t.Errorf("Expected error for position 0, got nil")
	}
}

func TestInsertSliceAt_InvalidPositionTooBig(t *testing.T) {
	original := []int{1, 2, 3}
	toInsert := []int{4, 5}
	position := 6

	_, err := InsertSliceAt(original, toInsert, position)
	if err == nil {
		t.Errorf("Expected error for position out of range, got nil")
	}
}

func TestInsertSliceAt_EmptyOriginal(t *testing.T) {
	original := []int{}
	toInsert := []int{1, 2}
	position := 1
	expected := []int{1, 2}

	result, err := InsertSliceAt(original, toInsert, position)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestInsertSliceAt_EmptyToInsert(t *testing.T) {
	original := []int{1, 2, 3}
	toInsert := []int{}
	position := 2
	expected := []int{1, 2, 3}

	result, err := InsertSliceAt(original, toInsert, position)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
