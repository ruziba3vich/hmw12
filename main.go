package main

//  1. Oldindan belgilangan qoidalarga muvofiq foydalanuvchi
//     kiritishini ValidateUser digan method yarating.
//
// Qoidalar:
// Name empty bo'lishi kerak emas
// Name uzunligi kamida 6 belgigi bo'lishi kerak
// Age 0 dan kotta va 120 dan kichik bo'lishi
// Email empty bo'lishi kerak emas
// Email formatiga mos bolishi kerak (masalan example@domain.com)
//
// 2. Error slice yaratilgan holda barcha paydo bo'lgan errorlarni qoshib yuvorin
// 3. Foydalanuvchi ma'lumotlarni terminaldan oqib oling
// 4. Oqib oliniyatgan jarayonda errorlarni ohirida chiqarib berin
//
// Masalan:
// Name: asd
// Age: 123
// Email: ""

// Errors:
// Name: length cannot be less than a 6 characters
// Age: couldn't be more than 120
// Email: couldn't be empty

import (
	"database/sql"
	"fmt"
	registerservice "hmw12/registerService"
	validation "hmw12/validationService"
	"log"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	sendRequest "hmw12/requestSender"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "host=localhost port=5432 user=prodonik dbname=user_registration_hmw12 password=Dost0n1k sslmode=disable")
	/*
		Databaza yaratish

		CREATE DATABASE user_registration_hmw12
			WITH
			OWNER = prodonik
			ENCODING = 'UTF8'
			LC_COLLATE = 'C.UTF-8'
			LC_CTYPE = 'C.UTF-8'
			LOCALE_PROVIDER = 'libc'
			TABLESPACE = pg_default
			CONNECTION LIMIT = -1
			IS_TEMPLATE = False;
		--------------------------------------
		Table yaratish

		CREATE TABLE users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(64) NOT NULL,
			email VARCHAR(64) NOT NULL,
			password VARCHAR(64) NOT NULL
		);

	*/

	if err != nil {
		log.Fatal(err)
	}

	// requestlar bajarilishi uchun Gin dan foydalandim
	r := gin.Default()

	// shu portga web-reques ya'ni so'rov jo'natamiz, ya'ni service shu end-point ga murojat qiladi
	r.POST("/register", func(c *gin.Context) {
		registerservice.RegisterUser(c, db)
	})

	// Localhost dagi port, dastur shu portda run bo'ladi
	r.Run(":7777")
}

func main() {
	// Deferni ham ishlatdim, bundan oldin init ni ichida ishlatayotgandim, lekin init tugashi bilan
	// database yopilib qolyatganakan, shuning uchun main ni oxirida yopiladigan qildim
	defer db.Close()
	status := 1
	usernames := make(map[string]bool)

	for status == 1 {
		errors := [] error {}
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
		fmt.Scan(& status)
	}

	fmt.Println("Dasturdan foydalanganingiz uchun raxmat !")
}
