package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func log(r *http.Request){
	ip = r.Header.Get("X-Forwarded-For")
	fmt.Println(ip)

func main() {
	// Говорим: "На запрос /search отвечай функцией getSearch"
	http.HandleFunc("GET /search", getSearch)
	
	fmt.Println("Сервер запущен на http://localhost:8080")
	fmt.Println("Пример: http://localhost:8080/search?q=google")
	
	// Запускаем сервер на порту 8080
	http.ListenAndServe(":8080", nil)
}

// getSearch - обрабатывает GET запросы на /search
func getSearch(w http.ResponseWriter, r *http.Request) {
	// 1. Получаем параметр q из URL (например, ?q=google)
	q := r.URL.Query().Get("q")
	
	// Если q пустой - просто ничего не делаем
	if q == "" {
		return
	}
	
	// 2. Загружаем данные из data.json
	dat := Data("data.json")
	
	// 3. Ищем в данных строку q
	res := Search(dat, q)
	
	// 4. Говорим браузеру, что отправляем JSON
	w.Header().Set("Content-Type", "application/json")

	log(r)
	
	// 5. Превращаем результат в JSON и отправляем
	json.NewEncoder(w).Encode(res)
}
