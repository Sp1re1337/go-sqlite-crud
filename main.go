package main

import (
	"log"
    "telegram-bot-sql-example/bot"
    "telegram-bot-sql-example/database"
)

func main() {
    // Ініціалізуємо базу даних
    err := database.InitDB()
    if err != nil {
        log.Fatalf("Помилка ініціалізації бази даних: %v", err)
    }

    // Запускаємо Telegram бота
    err = bot.StartBot()
    if err != nil {
        log.Fatalf("Помилка запуску бота: %v", err)
    }
}