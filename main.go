package main

import (
	"fmt"
	"math/big"
	"random_draw/cryptops"
)

var (
	participants = 11
	winners = 1
)

// CreateList initializes a []int ordered list from 0 to n elements
func CreateList(n int)([]int){
	list := make([]int, n)

	for i:=0; i < n; i++{
		list[i] = i
	}
	return list
}

// HexStringToBytes receives a string in hexadecimal format and returns a []byte slice version of the string
func HexStringToBytes(txt string, ad []byte)([]byte){
	a := big.NewInt(0)
	a.SetString(txt, 16)

	return append(a.Bytes(), ad...)
}

func main(){

	// Block at depth 619560 in the Bitcoin blockchain (2020-02-29 14:00)
	//seed := "00000000000000000004334492dda582eb9b12a5b6b87f15195581ceb177ab53"

	// Block at depth 620554 in the Bitcoin blockchain (2020-03-06 20:31)
	seed2 := "00000000000000000004a1a1fb5e42b514ed037e9e202ba9371f3857415c09c9"

	AD := []byte("Chaum")
	list := CreateList(participants)
	shuffledList := cryptops.ShuffleList(HexStringToBytes(seed2, AD), list)

	fmt.Println("Shuffled List: ", shuffledList)

	for j:=0; j < winners; j++{
		fmt.Println("The winner", j+1, "of the contest is the participant #", shuffledList[j])
	}
}
