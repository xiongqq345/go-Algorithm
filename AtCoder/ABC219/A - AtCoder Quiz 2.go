package main

import "fmt"

//In the Kingdom of AtCoder, there is a standardized test of competitive programming proficiency.
//
//An examinee will get a score out of 100 and obtain a rank according to the score, as follows:
//
//Novice, for a score not less than 0 but less than 40;
//Intermediate, for a score not less than 40 but less than 70;
//Advanced, for a score not less than 70 but less than 90;
//Expert, for a score not less than 90.
//Takahashi took this test and got X points.
//
//Find the minimum number of extra points needed to go up one rank higher. If, however, his rank was Expert, print expert, as there is no higher rank than that.

func main() {
	var x int
	fmt.Scan(&x)
	switch {
	case x < 40:
		fmt.Println(40 - x)
	case x < 70:
		fmt.Println(70 - x)
	case x < 90:
		fmt.Println(90 - x)
	default:
		fmt.Println("expert")
	}
}
