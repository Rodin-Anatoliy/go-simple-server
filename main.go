package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Id        uint64  `json:"id"`
	Email     string  `json:"email,omitempty"`
	Amount    uint64  `json:"amount"`
	Profile   Profile `json:"profile"`
	Username  string  `json:"username,omitempty"`
	CreatedAt string  `json:"createdAt"`
	CreatedBy string  `json:"createdBy"`
}

type Profile struct {
	Avatar    string `json:"avatar,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	FirstName string `json:"first_name,omitempty"`
}

func main() {
	http.HandleFunc("/", userHandler)

	log.Println("Сервер запущен на порту :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL.Path)

	data, err := getUsers()
	if err != nil {
		log.Printf("Ошибка получения пользователей: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	res := dataFilter(data)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("Ошибка кодирования: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func getUsers() ([]User, error) {
	data, err := http.Get("http://83.136.232.77:8091/users")
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса к серверу: %w", err)
	}
	
	var res []User
	if err := json.NewDecoder(data.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("ошибка декодирования JSON: %w", err)
	}

	return res, nil
}

func dataFilter(users []User) []User {
	var filtered []User
	for _, user := range users {
		if user.Amount > 50000 {
			user.Email = ""
			user.Username = ""
			user.Profile = Profile{}
		}
		filtered = append(filtered, user)
	}
	return filtered
}