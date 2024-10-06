package main

import (
	"database/sql"
  "fmt"
  "tgsql-example/database"
  "log"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID   int
  Name string
  Age  int
}

func main() {
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	database.CreateTable(db)

	database.AddUser(db, "Вася", 28)
  database.AddUser(db, "Коля", 24)

	users, err := database.GetUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Список користувачів:")
	for _, user := range users {
		fmt.Printf("ID: %d, Ім'я: %s, Вік: %d\n", user.ID, user.Name, user.Age)
	}

	database.DeleteUser(db, 1)

	users, err = database.GetUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Оновлений список користувачів:")
	for _, user := range users {
		fmt.Printf("ID: %d, Ім'я: %s, Вік: %d\n", user.ID, user.Name, user.Age)
	}
}