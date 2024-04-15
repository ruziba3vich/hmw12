package registerservice

import (
	"bufio"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"

	model "github.com/ruziba3vich/hmwnmb12/models"

	"github.com/gin-gonic/gin"
)

func RegisterService(c *gin.Context, db *sql.DB) {
	file, _ := os.Create("user_data.txt")
	defer file.Close()
	writer := bufio.NewWriter(file)

	// yangi userni yaratyabmiz
	var user model.User
	// c *gin.Context ichida bizga kereli ma'lumotlar JSON formatda keladi, shuni Go dagi structga o'giramiz
	if err := c.BindJSON(&user); err != nil {
		// Agar qanaqadir ma'lumot noto'g'ri keb qosa error qaytaramiz, masalan nimadir kam kiritilishi mumkin
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.SetPassword(encrypt(user.Password, 19))

	// databazaga yaratgan userimizi qo'shish
	_, err := db.Exec("INSERT INTO users (username, email, password, age) VALUES ($1, $2, $3, $4)", user.Username, user.Email, user.Password, user.Age)
	if err != nil {
		// agar databazaga qo'shishda muammo bo'lsa error qaytarish
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// databazada users ning `id` field i SERIAL yani auto increment qib ochilgan, shu sababli
	// user databazadan olgan id sini userga inject qib qo'yyabmiz

	var userID int
	row := db.QueryRow("SELECT id FROM users WHERE username = $1 AND password = $2", user.Email, user.Password)
	fmt.Fprintf(writer, "Username: %s\n", user.Username)
	fmt.Fprintf(writer, "Email: %s\n", user.Email)
	fmt.Fprintf(writer, "Password: %s\n", user.Password)
	fmt.Fprintf(writer, "Age: %s\n", user.Age)
	fmt.Fprintln(writer, "Came here")
	// Attempt to scan the result into userID
	err = row.Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned by the query, which means the user credentials are invalid
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}
		// Other error occurred, handle it accordingly
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Flush the writer to ensure all data is written to the file
	writer.Flush()

	user.Id = userID

	c.JSON(http.StatusCreated, user)
}

func encrypt(text string, shift int) string {
	var result strings.Builder

	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			result.WriteByte((byte(char)-'a'+byte(shift))%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			result.WriteByte((byte(char)-'A'+byte(shift))%26 + 'A')
		} else {
			result.WriteByte(byte(char))
		}
	}

	return result.String()
}

// package registerservice

// import (
// 	"bufio"
// 	"database/sql"
// 	"fmt"
// 	model "hmw12/models"
// 	"net/http"
// 	"os"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// )

// func RegisterUser(c *gin.Context, db *sql.DB) {
// 	file, err := os.Create("user_data.txt")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	defer file.Close()
// 	writer := bufio.NewWriter(file)

// 	// Create a new user instance
// 	var user model.User
// 	// Bind the JSON data from the request to the user struct
// 	if err := c.BindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Encrypt the user's password
// 	user.SetPassword(encrypt(user.Password, 19))

// 	// Insert the user into the database
// 	_, err = db.Exec("INSERT INTO users (username, email, password, age) VALUES ($1, $2, $3, $4)", user.Username, user.Email, user.Password, user.Age)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Query the database to fetch the user's ID
// 	var userID int
// 	row := db.QueryRow("SELECT id FROM users WHERE username = $1 AND password = $2", user.Email, user.Password)

// 	// Attempt to scan the result into userID
// 	err = row.Scan(&userID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			// No rows were returned by the query, which means the user credentials are invalid
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
// 			return
// 		}
// 		// Other error occurred, handle it accordingly
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Write user data to the file
// 	fmt.Fprintf(writer, "Username: %s\n", user.Username)
// 	fmt.Fprintf(writer, "Email: %s\n", user.Email)
// 	fmt.Fprintf(writer, "Password: %s\n", user.Password)
// 	fmt.Fprintf(writer, "Age: %s\n", user.Age)

// 	// Flush the writer to ensure all data is written to the file
// 	err = writer.Flush()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Assign the userID to the user struct
// 	user.Id = userID

// 	// Respond with status 201 and the user data
// 	c.JSON(http.StatusCreated, user)
// }

// func encrypt(text string, shift int) string {
// 	var result strings.Builder

// 	for _, char := range text {
// 		if char >= 'a' && char <= 'z' {
// 			result.WriteByte((byte(char)-'a'+byte(shift))%26 + 'a')
// 		} else if char >= 'A' && char <= 'Z' {
// 			result.WriteByte((byte(char)-'A'+byte(shift))%26 + 'A')
// 		} else {
// 			result.WriteByte(byte(char))
// 		}
// 	}

// 	return result.String()
// }
