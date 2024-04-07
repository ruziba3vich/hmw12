package registerservice

import (
	"database/sql"
	model "hmw12/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func RegisterUser(c *gin.Context) {
	// yangi userni yaratyabmiz
	var user model.User
	// c *gin.Context ichida bizga kereli ma'lumotlar JSON formatda keladi, shuni Go dagi structga o'giramiz
	if err := c.BindJSON(&user); err != nil {
		// Agar qanaqadir ma'lumot noto'g'ri keb qosa error qaytaramiz, masalan nimadir kam kiritilishi mumkin
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.SetPassword(encrypt(user.GetPassword(), 19))

	// databazaga yaratgan userimizi qo'shish
	_, err := db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.GetEmail(), user.GetPassword())
	if err != nil {
		// agar databazaga qo'shishda muammo bo'lsa error qaytarish
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// databazada users ning `id` field i SERIAL yani auto increment qib ochilgan, shu sababli
	// user databazadan olgan id sini userga inject qib qo'yyabmiz

	var userID int
	err = db.QueryRow("SELECT id FROM users WHERE username = $1 AND password = $2", user.GetEmail(), user.GetPassword()).Scan(&userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.SetId(userID)

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
