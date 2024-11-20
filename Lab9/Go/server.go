package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
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
var sessionTokens = map[string]string{} // Хранение токенов с привязкой к email

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/register", handleRegister).Methods("POST")
	router.HandleFunc("/login", handleLogin).Methods("POST")
	router.HandleFunc("/logout", handleLogout).Methods("POST")
	router.HandleFunc("/users", getAllUsers).Methods("GET")                // Получение списка пользователей
	router.HandleFunc("/users/create", createUser).Methods("POST")         // Создание пользователя
	router.HandleFunc("/users/update/{id}", updateUser).Methods("PUT")     // Обновление пользователя
	router.HandleFunc("/users/delete/{id}", deleteUser).Methods("DELETE")  // Удаление пользователя
	router.HandleFunc("/users/update-info", updateUserInfo).Methods("PUT") // Обновление информации о пользователе (имя/email)

	fmt.Println("Сервер запущен на порту 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Обработчик для регистрации пользователя
func handleRegister(w http.ResponseWriter, r *http.Request) {
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Неверные данные", http.StatusBadRequest)
		return
	}

	// Валидация имени и возраста
	if newUser.Name == "" {
		http.Error(w, "Имя не может быть пустым", http.StatusBadRequest)
		return
	}

	if newUser.Age < 0 {
		http.Error(w, "Возраст не может быть отрицательным", http.StatusBadRequest)
		return
	}

	newUser.ID = uint(len(users) + 1)
	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// Обработчик для входа пользователя
func handleLogin(w http.ResponseWriter, r *http.Request) {
	var credentials User
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Неверные данные", http.StatusBadRequest)
		return
	}

	for _, user := range users {
		if user.Email == credentials.Email && user.Password == credentials.Password {
			token := uuid.NewString() // Генерация уникального токена
			sessionTokens[user.Email] = token
			json.NewEncoder(w).Encode(map[string]string{"token": token})
			return
		}
	}

	http.Error(w, "Неверные данные для входа", http.StatusUnauthorized)
}

// Обработчик для выхода пользователя
func handleLogout(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	email := ""
	for e, t := range sessionTokens {
		if "Bearer "+t == token {
			email = e
			break
		}
	}

	if email != "" {
		delete(sessionTokens, email)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Выход успешен"})
	} else {
		http.Error(w, "Неверный токен", http.StatusUnauthorized)
	}
}

// Обработчик для получения всех пользователей
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if !isAuthenticated(token) {
		http.Error(w, "Невозможно авторизоваться", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(users)
}

// Обработчик для создания нового пользователя
func createUser(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if !isAuthenticated(token) {
		http.Error(w, "Невозможно авторизоваться", http.StatusUnauthorized)
		return
	}

	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Неверные данные", http.StatusBadRequest)
		return
	}

	// Валидация имени и возраста
	if newUser.Name == "" {
		http.Error(w, "Имя не может быть пустым", http.StatusBadRequest)
		return
	}

	if newUser.Age < 0 {
		http.Error(w, "Возраст не может быть отрицательным", http.StatusBadRequest)
		return
	}

	newUser.ID = uint(len(users) + 1)
	users = append(users, newUser)
	json.NewEncoder(w).Encode(newUser)
}

// Обработчик для обновления информации о пользователе
func updateUser(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if !isAuthenticated(token) {
		http.Error(w, "Невозможно авторизоваться", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Неверные данные", http.StatusBadRequest)
		return
	}

	// Валидация имени и возраста
	if updatedUser.Name == "" {
		http.Error(w, "Имя не может быть пустым", http.StatusBadRequest)
		return
	}

	if updatedUser.Age < 0 {
		http.Error(w, "Возраст не может быть отрицательным", http.StatusBadRequest)
		return
	}

	for i, user := range users {
		if int(user.ID) == id {
			users[i].Name = updatedUser.Name
			users[i].Age = updatedUser.Age
			users[i].Email = updatedUser.Email
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}

	http.Error(w, "Пользователь не найден", http.StatusNotFound)
}

// Обработчик для удаления пользователя
func deleteUser(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if !isAuthenticated(token) {
		http.Error(w, "Невозможно авторизоваться", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	for i, user := range users {
		if int(user.ID) == id {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Пользователь не найден", http.StatusNotFound)
}

// Новый обработчик для изменения имени или почты пользователя
func updateUserInfo(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if !isAuthenticated(token) {
		http.Error(w, "Невозможно авторизоваться", http.StatusUnauthorized)
		return
	}

	var newUserData User
	if err := json.NewDecoder(r.Body).Decode(&newUserData); err != nil {
		http.Error(w, "Неверные данные", http.StatusBadRequest)
		return
	}

	// Найти пользователя по токену
	email := ""
	for e, t := range sessionTokens {
		if "Bearer "+t == token {
			email = e
			break
		}
	}

	if email == "" {
		http.Error(w, "Неверный токен", http.StatusUnauthorized)
		return
	}

	// Обновить имя или email
	for i, user := range users {
		if user.Email == email {
			if newUserData.Name != "" {
				users[i].Name = newUserData.Name
			}
			if newUserData.Email != "" {
				users[i].Email = newUserData.Email
			}
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}

	http.Error(w, "Пользователь не найден", http.StatusNotFound)
}

// Проверка авторизации пользователя по токену
func isAuthenticated(token string) bool {
	token = strings.TrimPrefix(token, "Bearer ")
	for _, t := range sessionTokens {
		if t == token {
			return true
		}
	}
	return false
}
