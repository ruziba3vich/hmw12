package validation

import (
	"errors"
	"fmt"
	"strings"
)

func ValidateEmail(email string) error {
	if len(email) == 0 {
		return errors.New("length of an email cannot be blank")
	}
	ss := strings.Split(email, ".")
	if len(ss) != 2 {
		return fmt.Errorf("the email %s does not support correct e-mail format", email)
	}
	dd := strings.Split(ss[0], "@")

	if len(dd) != 2 {
		return fmt.Errorf("the email %s does not support correct e-mail format", email)
	}

	return nil
}

func ValidateAge(age int) error {
	if age > 0 && age < 121 {
		return nil
	}
	return fmt.Errorf("the age %d is not valid", age)
}

func ValidateUsername(username string, usernames map[string]bool) error {
	if usernames[username] {
		return fmt.Errorf("the username %s already exists", username)
	}
	if len(username) < 7 {
		return fmt.Errorf("length of a username `%s` cannot be less than 7", username)
	}
	return nil
}
