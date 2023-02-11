package main

import (	
	"reflect"
	"testing"
)


func TestMean(t *testing.T) {

	data := []float64{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 
		21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 
		31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 
		41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 
		51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 
		61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 
		71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 
		81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 
		91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
	}

	expected := 50.5
	result := mean(data)
	if result != expected {
		t.Errorf("mean of %v was incorrect, got: %f, want: %f.", data, result, expected)
	}
}


func TestMedian(t *testing.T) {

	data := []float64{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 
		21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 
		31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 
		41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 
		51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 
		61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 
		71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 
		81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 
		91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
	}

	expected := 50.0
	result := median(data)
	if result != expected {
		t.Errorf("median of %v was incorrect, got: %f, want: %f.", data, result, expected)
	}
}



func TestQuartile(t *testing.T) {
	data := []float64{1.0, 2.0, 3.0, 4.0}
	expected := 2.0
	result := quartile(data, 0.5)

	if result != expected {
		t.Errorf("Expected quartile of %v to be %v, got %v", data, expected, result)
	}
}

func TestMode(t *testing.T) {
	data := []float64{1.0, 2.0, 2.0, 3.0, 4.0}
	expected := 2
	result := mode(data)

	if result != expected {
		t.Errorf("Expected mode of %v to be %v, got %v", data, expected, result)
	}
}

func TestDataRange(t *testing.T) {
	data := []float64{1.0, 2.0, 3.0, 4.0}
	expected := 3.0
	result := data_range(data)

	if result != expected {
		t.Errorf("Expected data range of %v to be %v, got %v", data, expected, result)
	}
}

func TestDeMean(t *testing.T) {
	data := []float64{1.0, 2.0, 3.0, 4.0}
	expected := []float64{-1.5, -0.5, 0.5, 1.5}
	result := de_mean(data, mean)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected de-mean of %v to be %v, got %v", data, expected, result)
	}
}








// func TestQuartile(t *testing.T){

// 	data := []float64{
// 		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 
// 		11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 
// 		21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 
// 		31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 
// 		41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 
// 		51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 
// 		61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 
// 		71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 
// 		81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 
// 		91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
// 	}

// 	expected := 10.0
// 	result := quartile(data, 1.0)
// 	if result != expected {
// 		t.Errorf("quartile of %v was incorrect, got: %f, want: %f.", data, result, expected)
// 	}

// }
