package main

import (
	"fmt"
	sendRequest "hmw12/requestSender"
	validation "hmw12/validationService"
)

func main() {
	fmt.Println("\n\n\n.")
	status := 1
	usernames := make(map[string]bool)

	for status == 1 {
		errors := []error{}
		var username string
		fmt.Print("username kiriting : ")
		fmt.Scan(&username)
		if err := validation.ValidateUsername(username, usernames); err != nil {
			errors = append(errors, err)
		} else {
			usernames[username] = true
		}

		var email string
		fmt.Print("e-mail kiriting : ")
		fmt.Scan(&email)
		if err := validation.ValidateEmail(email); err != nil {
			errors = append(errors, err)
		}

		var password string
		fmt.Print("parol kiriting : ")
		fmt.Scan(&password)

		var age int
		fmt.Print("yosh kiriting : ")
		fmt.Scan(&age)
		if err := validation.ValidateAge(age); err != nil {
			errors = append(errors, err)
		}

		if len(errors) > 0 {
			for _, err := range errors {
				fmt.Println(err)
			}
			fmt.Println()
		} else {
			results, err := sendRequest.SendRequest(username, email, password, age)
			if err != nil {
				fmt.Println(err)
			} else {
				for _, r := range results {
					fmt.Println(r)
				}
				fmt.Println()
			}
		}

		fmt.Print("Yana kimnidir ro'yhatdan o'tkazasizmi ? : 1 -> ha / 2 -> yo'q : ")
		fmt.Scan(&status)
	}

	fmt.Println("Dasturdan foydalanganingiz uchun raxmat !")
}
