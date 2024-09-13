package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	server := &http.Server{Addr: ":8080"}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Сервер запущен на порту "+server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка сервера: %v", err)
		}
	}()

	<-stop
	log.Println("Получен сигнал завершения работы сервера...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка при завершении работы сервера: %v", err)
	}

	log.Println("Сервер корректно завершил работу")
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
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://83.136.232.77:8091/users", nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания запроса: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса к серверу: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("неверный статус ответа: %s", resp.Status)
	}

	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, fmt.Errorf("ошибка декодирования JSON: %w", err)
	}

	return users, nil
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
