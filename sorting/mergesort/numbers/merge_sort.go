package main

import "fmt"


func MergeSortAcending(numbers []int) []int {
	current_input_length := len(numbers)
	if current_input_length > 1 {
		mid := current_input_length / 2

		leftHalf := numbers[:mid]
		rightHalf := numbers[mid:]

		sorted_left := MergeSortAcending(leftHalf)
		sorted_right := MergeSortAcending(rightHalf)

		i := 0
		j := 0
		merged := make([]int, 0)

		for i < len(sorted_left) && j < len(sorted_right) {
			if sorted_left[i] < sorted_right[j] {
				merged = append(merged, sorted_left[i])
				i++
			} else {
				merged = append(merged, sorted_right[j])
				j++
			}
		}

		for i < len(sorted_left) {
			merged = append(merged, sorted_left[i])
			i++
		}

		for j < len(sorted_right) {
			merged = append(merged, sorted_right[j])
			j++
		}

		return merged
	}
	return numbers
}


func MergeSortDecending(numbers []int) []int {
	current_input_length := len(numbers)
	fmt.Println("input: ", numbers)
	if current_input_length > 1 {
		mid := current_input_length / 2

		leftHalf := numbers[:mid]
		rightHalf := numbers[mid:]

		sorted_left := MergeSortDecending(leftHalf)
		sorted_right := MergeSortDecending(rightHalf)

		i := 0
		j := 0
		merged := make([]int, 0)

		for i < len(sorted_left) && j < len(sorted_right) {
			if sorted_left[i] > sorted_right[j] {
				merged = append(merged, sorted_left[i])
				i++
			} else {
				merged = append(merged, sorted_right[j])
				j++
			}
		}

		for i < len(sorted_left) {
			merged = append(merged, sorted_left[i])
			i++
		}

		for j < len(sorted_right) {
			merged = append(merged, sorted_right[j])
			j++
		}

		fmt.Println("left: ", leftHalf)
		fmt.Println("right: ", rightHalf)
		fmt.Println("sorted left: ", sorted_left)
		fmt.Println("sorted right: ", sorted_right)
		fmt.Println("merged: ", merged)
		return merged
	}
	return numbers
}
