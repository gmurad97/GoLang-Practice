package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Инициализируем маршрутизатор
	mux := http.NewServeMux()

	// Устанавливаем обработчики маршрутов
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/contact", contactHandler)

	// Запускаем веб-сервер
	log.Println("Запуск веб-сервера на :8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Добро пожаловать на главную страницу!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Это страница о нас.")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Свяжитесь с нами по электронной почте: example@example.com")
}
