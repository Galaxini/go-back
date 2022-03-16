package services

import (
	"database/sql"
	"errors"
	"fmt"
	"go-back/app/models"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var err error

func RegisterUser(name string, password string, email string) (string, error) {
	// open db
	db, err = sql.Open("mysql", "root:@/my_app_db")
	if err != nil {
		return "", err
	}
	queryString := "insert into users(name, password, email, token) values (?, ?, ?, ?)"
	// prepare query
	stmt, err := db.Prepare(queryString)
	if err != nil {
		return "", err
	}
	// create hashed pass and token
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	token := models.GenerateSecureToken()
	// execute query
	_, err = stmt.Exec(name, hashedPassword, email, token)
	if err != nil {
		return "", err
	}
	// close db
	defer stmt.Close()
	// return success status
	return "Success\r\n", nil
}

func LoginUser(email string, password string) (string, error) {
	// open db
	db, err = sql.Open("mysql", "root:@/my_app_db")
	if err != nil {
		return "", err
	}
	// prepare query
	stmt, err := db.Prepare("SELECT name, email, password, id from users WHERE email = ?")
	if err != nil {
		fmt.Println("Failed to run query", err)
		return "", err
	}
	// execute query
	results, err := stmt.Query(email)
	if err != nil {
		fmt.Println("Failed to run query", err)
		return "", err
	}
	// close db
	defer results.Close()
	// create user
	var user models.User
	// scan results
	for results.Next() {
		err := results.Scan(&user.Name, &user.Email, &user.AccountPassword, &user.ID)
		if err != nil {
			panic(err.Error())
		}
	}
	// check for password
	err = bcrypt.CompareHashAndPassword([]byte(user.AccountPassword), []byte(password))
	if err != nil {
		return "", errors.New("Invalid username or password.\r\n")
	}
	// update token
	if user.ID != 0 {
		token := models.GenerateSecureToken()
		up_stmt, err := db.Prepare("UPDATE users SET token = ? WHERE email = ?")
		if err != nil {
			fmt.Println("Failed to run query", err)
			return "", err
		}
		_, err = up_stmt.Exec(token, email)
		if err != nil {
			return "", err
		}
		defer up_stmt.Close()
	} else {
		return "User not found", nil
	}

	return "Successfully logined", nil
}

// public function login($email, $password)
// {
//   try {
//     $password = User::getHash($password);
//     $user = User::where('email', $email)->where('password', $password)->first();
//     $user->token = User::generateToken();
//     $user->save();
//     return $user;
//   } catch(\Exception $e) {
//     return false;
//   }
// }
