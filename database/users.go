package database

import (
	"fmt"
	"log"
	"time"

	"github.com/loggercode/ecom/auth"
	uuid "github.com/nu7hatch/gouuid"
)

type User struct {
	Id           int64  `json:"id"`
	Uuid         string `json:"uuid"`
	Fisrtname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Username     string `json:"username"`
	Address      string
	Location     string `json:"location"`
	PhoneNumber  string `json:"phonenumber"`
	Email        string `json:"email"`
	PasswordHash string
	Password     string    `json:"password"`
	JoinedOn     time.Time `json:"joinedon"`
}

// CreateUser writes the userdetails to the database
func (user *User) CreateUser() (err error) {
	u4, err := uuid.NewV4()
	if err != nil {
		fmt.Println("error:", err)
		return

	}
	password_hash, err := auth.HashPassword(user.Password)
	if err != nil {
		return
	}

	_, err = DB.Exec(`insert into users(uuid , username,location, phonenumber, email, password_hash ) values(?,?,?,?,?,?)`, u4.String(), user.Username, user.Location, user.PhoneNumber, user.Email, string(password_hash))
	return
}

type UserCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserCredential) VerifyUser() (user User, err error) {
	err = DB.QueryRow("select uuid, username, location, phonenumber, password_hash from users where username = ?", u.Username).Scan(&user.Uuid, &user.Username, &user.Location, &user.PhoneNumber, &user.PasswordHash)
	if err != nil {
		log.Println(err)
		return
	}

	if err = auth.CompareHash(user.PasswordHash, u.Password); err != nil {
		log.Println(err)
		return
	}

	return
}
