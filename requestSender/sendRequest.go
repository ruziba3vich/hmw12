package requestsender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendRequest(username string, email string, password string) (results [] string, err error) {
	// So'rov jo'natish uchun JSON formatdagi ma'lumot yaratish
	requestBody, err := json.Marshal(map[string]string{
		"username": username,
		"email":    email,
		"password": password,
	})
	// Agar request (JSON formatdagi ma'lumot) yaratishda error chiqsa . . .
	if err != nil {
		return results, fmt.Errorf("error marshalling JSON: %s", err)
	}

	// yangi HTTP request yaratish, POST metodidan foydalanib
	req, err := http.NewRequest("POST", "localhost:7777/register", bytes.NewBuffer(requestBody))
	if err != nil {
		return results, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json") // Set the Content-Type header

	// so'rov jo'natish uchun bizga HTTP client kerak, shuning uchun yangi HTTP client yaratamiz
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return results, fmt.Errorf("error sending request: %s", err)
	}
	defer resp.Body.Close()

	results = append(results, fmt.Sprintf("Response Status: %s", resp.Status))
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	results = append(results, fmt.Sprintf("Response Body: %s", buf.String()))
	return results, nil
}
