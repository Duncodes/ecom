package database

import (
	"fmt"
	"log"
	"time"

	"github.com/Duncodes/ecom/auth"
	"github.com/twinj/uuid"
)

// User ...
type User struct {
	ID           int64  `json:"id"`
	UUID         string `json:"uuid"`
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
	u4 := uuid.NewV4()
	if err != nil {
		fmt.Println("error:", err)
		return

	}
	passwordhash, err := auth.HashPassword(user.Password)
	if err != nil {
		return
	}
	log.Println(u4)
	user.UUID = u4.String()
	_, err = DB.Exec(`insert into users(uuid , username,location, phonenumber,
	email, password_hash ) values(?,?,?,?,?,?)`, u4,
		user.Username, user.Location, user.PhoneNumber, user.Email, string(passwordhash))
	return
}

func GetUserByUUID(uuid string) (User, error) {
	var user User
	err := DB.QueryRow(`select uuid, username, location, phonenumber,
					password_hash from users where uuid = ?`, uuid).Scan(&user.UUID,
		&user.Username, &user.Location, &user.PhoneNumber, &user.PasswordHash)

	return user, err
}

// UserCredential ...
type UserCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// VerifyUser ...
func (u *UserCredential) VerifyUser() (user User, err error) {
	err = DB.QueryRow(`select uuid, username, location, phonenumber,
	password_hash from users where username = ?`, u.Username).Scan(&user.UUID,
		&user.Username, &user.Location, &user.PhoneNumber, &user.PasswordHash)
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
