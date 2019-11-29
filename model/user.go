package model

import (
    // "log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"strings"
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
	Email string  `gorm:"type:varchar(100);unique_index"`
	Password string
}


/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (user *User) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(user.Email, "@") {
		return utils.Message(false, "Email address is required"), false
	}

	if len(user.Password) < 6 {
		return utils.Message(false, "Password is required"), false
	}

	temp := &User{}

	//check for errors and duplicate emails
	err := GetDB().Table("users").Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return utils.Message(false, "Email address already in use by another user."), false
	}

	return utils.Message(true, "success"), true
}

func (user *User) Create() (map[string]interface{}) {
	if resp, ok := user.Validate(); !ok {
		return resp
	}

	fmt.Println(user.Password)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	fmt.Println(hashedPassword)
	fmt.Println(string(hashedPassword))

	GetDB().Create(user)

	if user.ID <= 0 {
		return utils.Message(false, "Failed to create user, connection error.")
	}

	//Create new JWT token for the newly registered user
	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN_PASSWORD")))
	user.Token = tokenString

	user.Password = "" //delete password

	response := utils.Message(true, "Account has been created")
	response["user"] = user

	return response
}

func Login(email, password string) (map[string]interface{}) {
	user := &User{}
	err := GetDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "Email address not found")
		}
		return utils.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return utils.Message(false, "Invalid login credentials. Please try again")
	}
	//Worked! Logged In
	user.Password = ""

	//Create JWT token
	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN_PASSWORD")))
	user.Token = tokenString //Store the token in the response

	resp := utils.Message(true, "Logged In")
	resp["user"] = user

	return resp
}