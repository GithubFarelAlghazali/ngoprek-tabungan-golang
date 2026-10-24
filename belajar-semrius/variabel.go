package main

import "fmt"

func main() {
	var username string = "James"
	isActive := true

	score := 76
	score = 98

	email, city, age := "jameson@mail.son", "Texas", 19

	role, _ := "student", "sudwbubwi128r1"

	gradeTier := "9A"
	gradePointer := new(string)
	gradePointer = &gradeTier

	prevScore, latestScore := 67, 88
	prevScore, latestScore = latestScore, prevScore

	// prints

	fmt.Println("username : ", username)
	fmt.Println("isActive? : ", isActive)
	fmt.Println("Exam Score : ", score)
	fmt.Println("Email : ", email)
	fmt.Println("City : ", city)
	fmt.Println("Age : ", age)
	fmt.Println("Role : ", role)
	fmt.Println("Grade : ", *gradePointer)
	fmt.Println("Previous Score : ", prevScore)
	fmt.Println("Latest Score : ", latestScore)

}