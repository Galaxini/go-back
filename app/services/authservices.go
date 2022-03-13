package services

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var err error

func RegisterUser(name string, password string, email string) (string, error) {

	db, err = sql.Open("mysql", "root:@/my_app_db")
	if err != nil {
		panic(err.Error())
	}
	queryString := "insert into users(name, password, email) values (?, ?, ?)"

	stmt, err := db.Prepare(queryString)

	if err != nil {
		return "", err
	}

	defer stmt.Close()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	_, err = stmt.Exec(name, hashedPassword, email)

	if err != nil {
		return "", err
	}

	return "Success\r\n", nil

}
