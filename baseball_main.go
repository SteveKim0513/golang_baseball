package main

import (
	"fmt"
	"math/rand"
	"time"
)

func computerNum() [3]int {

	var randNum [3]int

	rand.Seed(time.Now().UnixNano())

	// Computer picks three random numbers

	for {
		for i := 0; i < 3; i++ {
			randNum[i] = rand.Intn(10)
		}

		// Check the numbers if overlapped
		if checkInput(randNum) > 0 {
			fmt.Println("Error, Computer will pick Numbers again")
		} else {
			break
		}
	}
	return randNum
}

func inputUserNum() [3]int {

	var inputNum [3]int

	// Input the answer
	for {
		fmt.Println("Input Example : 1 2 5")
		fmt.Printf("My Answer : ")
		fmt.Scanf("%d %d %d", &inputNum[0], &inputNum[1], &inputNum[2])

		// Check the numbers - Range, Overlap
		if checkInput(inputNum) > 0 {
			fmt.Println("Error, input the numbers again")
		} else {
			break
		}
	}

	return inputNum
}

// Detect Number Error - Range, Overlap
func checkInput(arr [3]int) int {

	var resultNum int

	var temp [3]int
	temp = arr

	for i := 0; i < 3; i++ {

		// Check Range
		if arr[i] > 9 || arr[i] < 0 {
			resultNum++
			fmt.Printf("%d번째 variable : Out of Range\n", i+1)
			return resultNum
		}

		// Check Overlap
		for j := 0; j < 3; j++ {
			if i == j {
				continue
			} else {
				if arr[i] == temp[j] {
					resultNum++
					fmt.Printf("%d, %d Overlapped\n", i+1, j+1)
					return resultNum
				}
			}
		}
	}
	return resultNum
}

// Compare user input and computer input
func compareNum(arr1, arr2 [3]int) [2]int {

	// [0]=Ball, [1]=Strike
	var resultNum [2]int

	for i := 0; i < 3; i++ {

		for j := 0; j < 3; j++ {

			if arr1[i] == arr2[j] {
				if i == j {
					resultNum[1]++
				} else {
					resultNum[0]++
				}
			}
		}
	}

	return resultNum
}

func main() {

	var userNum [3]int
	var comNum [3]int
	var result [2]int

	comNum = computerNum()

	gameCount := 0
	for {
		userNum = inputUserNum()

		result = compareNum(userNum, comNum)
		gameCount++

		fmt.Printf("%d B, %d S \n", result[0], result[1])

		if result[1] == 3 {
			break
		}
	}
	fmt.Println("Correct ~ !")
	fmt.Printf("You win on %d Round\n", gameCount)

}
