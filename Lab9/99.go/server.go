package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid" // Импортируем пакет для генерации UUID
	"github.com/gorilla/mux"
)

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

var users = []User{}
var authTokens = map[string]string{} // Simple token storage, mapping email to token

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", registerHandler).Methods("POST")
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/logout", logoutHandler).Methods("POST")
	r.HandleFunc("/users", getUsersHandler).Methods("GET")                  // Получить всех пользователей
	r.HandleFunc("/users/create", createUserHandler).Methods("POST")        // Создать пользователя
	r.HandleFunc("/users/update/{id}", updateUserHandler).Methods("PUT")    // Обновить пользователя по ID
	r.HandleFunc("/users/delete/{id}", deleteUserHandler).Methods("DELETE") // Удалить пользователя по ID

	fmt.Println("Сервер запущен на порту 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Handler for user registration
func registerHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Неверные данные", http.StatusBadRequest)
		return
	}
	newUser.ID = uint(len(users) + 1)
	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// Handler for user login
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials User
	json.NewDecoder(r.Body).Decode(&credentials)
	for _, user := range users {
		if user.Email == credentials.Email && user.Password == credentials.Password {
			token := uuid.NewString() // Генерация уникального токена для каждого пользователя
			authTokens[user.Email] = token
			fmt.Printf("Токен выдан: %s для пользователя: %s при входе\n", token, user.Email) // Logging token issuance
			json.NewEncoder(w).Encode(map[string]string{"token": token})
			return
		}
	}
	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}

// Handler for user logout
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	email := ""
	for e, t := range authTokens {
		if "Bearer "+t == token {
			email = e
			break
		}
	}
	if email != "" {
		delete(authTokens, email)
		fmt.Printf("Пользователь: %s вышел из системы. Токен: %s\n", email, token) // Logging token on logout
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Logged out successfully"})
	} else {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
	}
}

// Handler to get all users
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if !isAuthenticated(token) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Printf("Запрос на получение всех пользователей. Токен: %s\n", token) // Logging token for the action
	json.NewEncoder(w).Encode(users)
}

// Handler to create a new user
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if !isAuthenticated(token) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Неверные данные", http.StatusBadRequest)
		return
	}
	newUser.ID = uint(len(users) + 1)
	users = append(users, newUser)
	fmt.Printf("Создан пользователь: %s. Токен: %s\n", newUser.Name, token) // Logging token for user creation
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// Handler to update a user by ID
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if !isAuthenticated(token) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Неверные данные", http.StatusBadRequest)
		return
	}
	for i, user := range users {
		if int(user.ID) == id {
			users[i].Name = updatedUser.Name
			users[i].Age = updatedUser.Age
			users[i].Email = updatedUser.Email
			fmt.Printf("Обновлен пользователь: %s (ID: %d). Токен: %s\n", users[i].Name, users[i].ID, token) // Logging token for user update
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

// Handler to delete a user by ID
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if !isAuthenticated(token) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for i, user := range users {
		if int(user.ID) == id {
			users = append(users[:i], users[i+1:]...)
			fmt.Printf("Удален пользователь с ID: %d. Токен: %s\n", id, token) // Logging token for user deletion
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

// Check if the request is authenticated
func isAuthenticated(token string) bool {
	token = strings.TrimPrefix(token, "Bearer ")
	for _, t := range authTokens {
		if t == token {
			return true
		}
	}
	return false
}
