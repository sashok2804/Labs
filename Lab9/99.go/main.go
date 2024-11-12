package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var token string

type User struct {
	ID       uint   `json:"id,omitempty"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

func register() {
	var user User
	fmt.Print("Введите имя: ")
	fmt.Scan(&user.Name)
	fmt.Print("Введите возраст: ")
	fmt.Scan(&user.Age)
	fmt.Print("Введите email: ")
	fmt.Scan(&user.Email)
	fmt.Print("Введите пароль: ")
	fmt.Scan(&user.Password)

	data, _ := json.Marshal(user)
	resp, err := http.Post("http://localhost:8000/register", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Ошибка регистрации: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Ответ сервера:", string(body))
}

func login() {
	var user User
	fmt.Print("Введите email: ")
	fmt.Scan(&user.Email)
	fmt.Print("Введите пароль: ")
	fmt.Scan(&user.Password)

	data, _ := json.Marshal(user)
	resp, err := http.Post("http://localhost:8000/login", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Ошибка авторизации: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		var result map[string]string
		if err := json.Unmarshal(body, &result); err != nil {
			log.Fatalf("Ошибка обработки ответа сервера: %v", err)
		}
		token = result["token"]
		fmt.Println("Авторизация успешна. Токен сохранен в памяти.")
	} else {
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Ошибка авторизации:", string(body))
	}
}

func getUsers() {
	if !isAuthenticated() {
		fmt.Println("Необходимо авторизоваться!")
		return
	}

	resp := makeRequest("GET", "http://localhost:8000/users", nil)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Список пользователей:", string(body))
	} else {
		fmt.Println("Ошибка получения пользователей:", string(body))
	}
}

func createUser() {
	if !isAuthenticated() {
		fmt.Println("Необходимо авторизоваться!")
		return
	}

	var user User
	fmt.Print("Введите имя: ")
	fmt.Scan(&user.Name)
	fmt.Print("Введите возраст: ")
	fmt.Scan(&user.Age)
	fmt.Print("Введите email: ")
	fmt.Scan(&user.Email)

	data, _ := json.Marshal(user)
	resp := makeRequest("POST", "http://localhost:8000/users/create", data)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusCreated {
		fmt.Println("Пользователь добавлен:", string(body))
	} else {
		fmt.Println("Ошибка создания пользователя:", string(body))
	}
}

func updateUser() {
	if !isAuthenticated() {
		fmt.Println("Необходимо авторизоваться!")
		return
	}

	var id uint
	fmt.Print("Введите ID пользователя для обновления: ")
	fmt.Scan(&id)

	var user User
	fmt.Print("Введите новое имя: ")
	fmt.Scan(&user.Name)
	fmt.Print("Введите новый возраст: ")
	fmt.Scan(&user.Age)
	fmt.Print("Введите новый email: ")
	fmt.Scan(&user.Email)

	data, _ := json.Marshal(user)
	url := fmt.Sprintf("http://localhost:8000/users/update/%d", id)
	resp := makeRequest("PUT", url, data)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Пользователь обновлен:", string(body))
	} else {
		fmt.Println("Ошибка обновления пользователя:", string(body))
	}
}

func deleteUser() {
	if !isAuthenticated() {
		fmt.Println("Необходимо авторизоваться!")
		return
	}

	var id uint
	fmt.Print("Введите ID пользователя для удаления: ")
	fmt.Scan(&id)

	url := fmt.Sprintf("http://localhost:8000/users/delete/%d", id)
	resp := makeRequest("DELETE", url, nil)
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		fmt.Println("Пользователь удален.")
	} else {
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Ошибка удаления пользователя:", string(body))
	}
}

func isAuthenticated() bool {
	return token != ""
}

func makeRequest(method, url string, data []byte) *http.Response {
	client := &http.Client{}
	var req *http.Request
	var err error

	if data != nil {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(data))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		log.Fatalf("Ошибка создания запроса: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	if data != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка выполнения запроса: %v", err)
	}

	return resp
}

func logout() {
	if !isAuthenticated() {
		fmt.Println("Необходимо авторизоваться!")
		return
	}

	resp := makeRequest("POST", "http://localhost:8000/logout", nil)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Выход из системы успешен.")
		token = ""
	} else {
		fmt.Println("Ошибка выхода из системы:", string(body))
	}
}

func main() {
	for {
		fmt.Println("\nМеню:")
		fmt.Println("1. Регистрация")
		fmt.Println("2. Авторизация")
		fmt.Println("3. Получить всех пользователей")
		fmt.Println("4. Добавить пользователя")
		fmt.Println("5. Обновить пользователя") // Updated menu option
		fmt.Println("6. Удалить пользователя")
		fmt.Println("7. Выход из системы")
		fmt.Println("0. Выход")

		var choice int
		fmt.Print("Выберите действие: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			register()
		case 2:
			login()
		case 3:
			getUsers()
		case 4:
			createUser()
		case 5:
			updateUser() // Updated case
		case 6:
			deleteUser()
		case 7:
			logout()
		case 0:
			fmt.Println("Выход из программы...")
			os.Exit(0)
		default:
			fmt.Println("Некорректный выбор, попробуйте снова.")
		}
	}
}
