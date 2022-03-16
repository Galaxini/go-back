package services

import (
	"database/sql"
	"fmt"
	"go-back/app/models"

	_ "github.com/go-sql-driver/mysql"
)

func AddItem(userId int, title string, description string, price string) (string, error) {
	db, err = sql.Open("mysql", "root:@/my_app_db")
	if err != nil {
		return "", err
	}
	queryString := "insert into items(user_id, title, description, price) values (?, ?, ?, ?)"
	stmt, err := db.Prepare(queryString)
	if err != nil {
		return "", err
	}
	// execute query
	_, err = stmt.Exec(userId, title, description, price)
	if err != nil {
		return "", err
	}
	// close db
	defer stmt.Close()
	// return success status
	return "New item was created", nil
}

func GetItems(userId int) (models.Item, error) {
	db, err = sql.Open("mysql", "root:@/my_app_db")
	if err != nil {
		return models.Item{}, err
	}
	stmt, err := db.Prepare("SELECT id, title, description, price, user_id from items WHERE user_id = ?")
	if err != nil {
		fmt.Println("Failed to run query", err)
		return models.Item{}, err
	}
	// execute query
	results, err := stmt.Query(userId)
	if err != nil {
		fmt.Println("Failed to run query", err)
		return models.Item{}, err
	}
	// close db
	defer results.Close()
	var item models.Item
	for results.Next() {
		err := results.Scan(&item.ID, &item.Title, &item.Description, &item.Price, &item.UserId)
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Printf("%+v\n", item)
	return item, nil
}

func EditItems(title string, description string, price string, id string) (string, error) {
	db, err = sql.Open("mysql", "root:@/my_app_db")
	if err != nil {
		return "", err
	}
	queryString := "UPDATE items SET title = ?, description = ?, price = ? WHERE id = ?"

	stmt, err := db.Prepare(queryString)
	if err != nil {
		return "", err
	}
	// execute query
	_, err = stmt.Exec(title, description, price, id)
	if err != nil {
		return "", err
	}
	// close db
	defer stmt.Close()
	// return success status
	return "Item was succesfully updated", nil
}

func DeleteItems(id string) (string, error) {
	db, err = sql.Open("mysql", "root:@/my_app_db")
	if err != nil {
		return "", err
	}
	queryString := "DELETE FROM Items WHERE id = ?"

	stmt, err := db.Prepare(queryString)
	if err != nil {
		return "", err
	}
	// execute query
	_, err = stmt.Exec(id)
	if err != nil {
		return "", err
	}
	// close db
	defer stmt.Close()
	// return success status
	return "Item was succesfully deleted", nil
}

func GetUserID(token string) (int, error) {
	// prepare query
	db, err = sql.Open("mysql", "root:@/my_app_db")
	if err != nil {
		return 0, err
	}
	stmt, err := db.Prepare("SELECT name, email, password, id from users WHERE token = ?")
	if err != nil {
		fmt.Println("Failed to run query", err)
		return 0, err
	}
	// execute query
	results, err := stmt.Query(token)
	if err != nil {
		fmt.Println("Failed to run query", err)
		return 0, err
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
	if user.ID != 0 {
		return user.ID, nil
	} else {
		return 0, nil
	}
}
