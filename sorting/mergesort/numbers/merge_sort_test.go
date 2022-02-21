package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMergeSort_NumbersInAcending(t *testing.T){
	input := []int{ 4, 6, 2, 1, 10, 56, 35, 82, 7, 22 }
	expected := []int{1, 2, 4, 6, 7, 10, 22, 35, 56, 82 }
	
	actual := MergeSortAcending(input)

	assert.Equal(t, expected, actual)

}

func TestMergeSort_NumbersInDecending(t *testing.T){
	input := []int{ 4, 6 , 0}
	expected := []int{ 6, 4, 0}
	
	actual := MergeSortDecending(input)

	assert.Equal(t, expected, actual)

}



//      
// [1,3,6,8,10,  30,56,    60, 70 , 80] 35
// mid = 4 ; nearSm=4; nearLm=7
// nearSm= 5 
// 
// 
// 
// 
// 

//      
// [1,3,6,8,10,  30,56,    60, 70 , 80] 81
// mid = 4 ; nearSm=7; nearLm=7
// nearSm= 5 
// 
// 
// 
// 
// 
