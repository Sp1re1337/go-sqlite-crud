package database

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/mattn/go-sqlite3"
)

// Структура для представлення користувача
type User struct {
    ID   int
    Name string
    Age  int
}

// Створення таблиці користувачів
func CreateTable(db *sql.DB) {
    createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        age INTEGER
    );`
    statement, err := db.Prepare(createTableSQL)
    if err != nil {
        log.Fatal(err)
    }
    statement.Exec()
    fmt.Println("Таблиця створена або вже існує.")
}

// Додавання нового користувача
func AddUser(db *sql.DB, name string, age int) {
    insertUserSQL := `INSERT INTO users (name, age) VALUES (?, ?)`
    statement, err := db.Prepare(insertUserSQL)
    if err != nil {
        log.Fatal(err)
    }
    _, err = statement.Exec(name, age)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Користувача додано:", name)
}

// Отримання всіх користувачів
func GetUsers(db *sql.DB) ([]User, error) {
    query := `SELECT id, name, age FROM users`
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        err = rows.Scan(&user.ID, &user.Name, &user.Age)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}

// Видалення користувача за ID
func DeleteUser(db *sql.DB, id int) {
    deleteUserSQL := `DELETE FROM users WHERE id = ?`
    statement, err := db.Prepare(deleteUserSQL)
    if err != nil {
        log.Fatal(err)
    }
    _, err = statement.Exec(id)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Користувача з ID", id, "видалено.")
}
